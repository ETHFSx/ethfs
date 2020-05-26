package uploader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"

	ethfsAbi "github.com/Yihen/ethfs/account/abi"
	"github.com/Yihen/ethfs/encodings"

	"golang.org/x/crypto/sha3"
)

type InputDeconstruct struct {
	HashedEthfsFilepath [32]byte
	ContainerSignature  [ethcrypto.SignatureLength]byte
	Params              encodings.ProposeUploadParams
	Shardsize           uint64
	EthfsSPs            [][20]byte
}

type UploadTx struct {
	BlockHash          string `json:blockHash`
	BlockNumber        string `json:blockNumber`
	From               string `json:from`
	Gas                string `json:gas`
	GasPrice           string `json:gasPrice`
	Hash               string `json:hash`
	InputDeconstructed InputDeconstruct
	Input              string `json:input`
	Nonce              string `json:nonce`
	R                  string `json:r`
	S                  string `json:s`
	To                 string `json:to`
	TransactionIndex   string `json:transactionIndex`
	V                  string `json:v`
	AmountPaid         string
	PublicKey          [64]byte
	UsernameCompressed [32]byte
}

type GetTxByHashResult struct {
	BlockHash        string `json:blockHash`
	BlockNumber      string `json:blockNumber`
	From             string `json:from`
	Gas              string `json:gas`
	GasPrice         string `json:gasPrice`
	Hash             string `json:hash`
	Input            string `json:input`
	Nonce            string `json:nonce`
	R                string `json:r`
	S                string `json:s`
	To               string `json:to`
	TransactionIndex string `json:transactionIndex`
	V                string `json:v`
	Value            string `json:value`
}

func checkUploadRpcCalls(txHash [32]byte) (res GetTxByHashResult, err error) {
	// gettxbyid rpc call
	var bTxHash []byte
	bTxHash = append(bTxHash, txHash[:]...)
	hexTxHash := hexutil.Encode(bTxHash)
	var reqString string = "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getTransactionByHash\",\"params\": [\"" + hexTxHash + "\"],\"id\":1}"
	var reqBytes = []byte(reqString)
	req, err_req := http.NewRequest("POST", ethfsAbi.Rpc().Url, bytes.NewBuffer(reqBytes))
	type Response struct {
		Result GetTxByHashResult `json:"result"`
	}
	var response Response
	if err_req != nil {
		return response.Result, fmt.Errorf("Error checkUploadRpcCalls: http request initialization error")
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{Timeout: time.Second * 10}
	resp, err_resp := client.Do(req)
	if resp == nil || err_resp != nil {
		return response.Result, fmt.Errorf("Error checkUploadRpcCalls: http request error")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err_json := json.Unmarshal(body, &response)
	if err_json != nil {
		return response.Result, fmt.Errorf("Error checkUploadRpcCalls: response parsing error")
	}
	return response.Result, nil
}

type ProposeUploadArgs struct {
	HashedEthfsFilepath [32]byte
	ContainerSignatureR [32]byte
	ContainerSignatureS [32]byte
	Params              [32]byte
	Shardsize           uint64
	EthfsSPs            [][20]byte
}

func Unpack(methodName string, input []byte) (inter interface{}, err error) {
	if methodName == "proposeUpload" {
		if len(input)%32 != 0 && len(input) < (4*32) {
			type Empty struct{}
			return new(Empty), fmt.Errorf("This method cannot be unpacked")
		}
		var hashedEthfsFilepath [32]byte
		var containerSignatureR, containerSignatureS [32]byte
		var params [32]byte
		copy(hashedEthfsFilepath[0:32], input[0:32])
		copy(containerSignatureR[0:32], input[32:64])
		copy(containerSignatureS[0:32], input[(2*32):(3*32)])
		copy(params[0:32], input[(3*32):(4*32)])
		var shardsize uint64
		for i := 24; i < 32; i++ { // getting the uint64 from the byte32
			var shift = 8 * (32 - i - 1)
			shardsize += uint64(input[i+(4*32)]) << shift
		}

		var ethfsSPs [][20]byte
		var idx int = 0
		for i := (7 * 32); i < len(input); i += 32 {
			var ethfsSP [20]byte
			copy(ethfsSP[0:20], input[i+12:i+12+20])
			ethfsSPs = append(ethfsSPs, ethfsSP)
			idx++
		}
		ret := ProposeUploadArgs{HashedEthfsFilepath: hashedEthfsFilepath,
			ContainerSignatureR: containerSignatureR,
			ContainerSignatureS: containerSignatureS,
			Params:              params,
			Shardsize:           shardsize,
			EthfsSPs:            ethfsSPs}
		return ret, nil
	}
	type Empty struct{}
	return new(Empty), fmt.Errorf("This method cannot be unpacked")
}

func ECRecoverFromTx(data GetTxByHashResult) (retKey [64]byte, err error) {
	dataR := strings.Replace(data.R, "0x", "", 1)
	dataS := strings.Replace(data.S, "0x", "", 1)
	dataV := strings.Replace(data.V, "0x", "", 1)
	if len(dataR)%2 == 1 {
		dataR = "0" + dataR
	}
	if len(dataS)%2 == 1 {
		dataS = "0" + dataS
	}
	if len(dataV)%2 == 1 {
		dataV = "0" + dataV
	}
	var sig []byte // = r + s + v
	var bR []byte
	for i := 0; i < len(dataR); i += 2 {
		r, _ := hexutil.Decode("0x" + dataR[i:i+2])
		bR = append(bR, []byte(r)...)
	}
	var bS []byte
	for i := 0; i < len(dataS); i += 2 {
		r, _ := hexutil.Decode("0x" + dataS[i:i+2])
		bS = append(bS, []byte(r)...)
	}
	var bV []byte
	for i := 0; i < len(dataV); i += 2 {
		r, _ := hexutil.Decode("0x" + dataV[i:i+2])
		bV = append(bV, []byte(r)...)
	}
	sig = append(sig, bR[:]...)
	sig = append(sig, bS[:]...)
	sig = append(sig, bV[:]...)

	nonce, err_res := strconv.ParseUint(strings.Replace(data.Nonce, "0x", "", 1), 16, 64)
	if err_res != nil {
		var empty [64]byte
		return empty, fmt.Errorf("ECRecoverFromTx error: parsing nonce error")
	}
	gasPrice, err_gasPrice := strconv.ParseUint(strings.Replace(data.GasPrice, "0x", "", 1), 16, 64)
	if err_gasPrice != nil {
		var empty [64]byte
		return empty, fmt.Errorf("ECRecoverFromTx error: parsing gasPrice error")
	}
	bigGasPrice := new(big.Int)
	bigGasPrice.SetUint64(gasPrice)
	gas, err_gas := strconv.ParseUint(strings.Replace(data.Gas, "0x", "", 1), 16, 64)
	if err_gas != nil {
		var empty [64]byte
		return empty, fmt.Errorf("ECRecoverFromTx error: parsing gas error")
	}
	bigGas := new(big.Int)
	bigGas.SetUint64(gas)
	var bTo []byte
	for i := 2; i < len(data.To); i += 2 {
		r, _ := hexutil.Decode("0x" + data.To[i:i+2])
		bTo = append(bTo, []byte(r)...)
	}
	var bbTo [20]byte
	copy(bbTo[0:20], bTo[0:20])
	dataValue, err_dataValue := strconv.ParseUint(strings.Replace(data.Value, "0x", "", 1), 16, 64)
	if err_dataValue != nil {
		var empty [64]byte
		return empty, fmt.Errorf("ECRecoverFromTx error: parsing value error")
	}
	bigValue := new(big.Int)
	bigValue.SetUint64(dataValue)

	hw := sha3.NewLegacyKeccak256()
	input := strings.Replace(data.Input, "0x", "", 1)
	var bInput []byte
	for i := 0; i < len(input); i += 2 {
		r, _ := hexutil.Decode("0x" + input[i:i+2])
		bInput = append(bInput, []byte(r)...)
	}
	var chainID byte
	if ethfsAbi.ChainIs() == "Gorli" {
		chainID = byte(5)
	} // else adapt to other chains
	rlp.Encode(hw, []interface{}{
		nonce,
		bigGasPrice,
		bigGas,
		bbTo,
		bigValue,
		bInput,
		byte(chainID), uint(0), uint(0)})
	var h common.Hash
	hw.Sum(h[:0])
	var bH []byte
	bH = append(bH, []byte(h[:])...)
	if bV[0] == byte(46) || bV[0] == byte(38) {
		sig[64] = byte(1)
	} else {
		sig[64] = byte(0)
	}
	pubKey, err_pub := ethcrypto.Ecrecover(bH, sig)
	if err_pub != nil {
		var empty [64]byte
		return empty, fmt.Errorf("ECRecoverFromTx error: ethcrypto.Ecrecover error")
	}
	var bPubKey [64]byte
	copy(bPubKey[0:64], pubKey[1:65])
	return bPubKey, nil
}
