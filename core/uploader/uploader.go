package uploader

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/Yihen/ethfs/common/constants"
	"github.com/ethereum/go-ethereum/common"

	proof "github.com/Yihen/contracts/dataproof/api"

	"github.com/Yihen/ethfs/common/log"

	"github.com/ETHFSx/go-ipfs/shell"

	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types" //NewTransaction
	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	ethfsAbi "github.com/Yihen/ethfs/account/abi"
	"github.com/Yihen/ethfs/account/wallet"
	"github.com/Yihen/ethfs/encodings"
	"github.com/Yihen/ethfs/rpc/utils"
)

var g_chainID int64 = ethfsAbi.ChainID()

type UploadAccessControlLevel uint8

const (
	Priv_UploaderOnly UploadAccessControlLevel = iota
	Priv_RestrictedGroup
	Public
)

type UploadParams struct {
	Wallet wallet.EthereumKeyset

	ServiceDuration    uint32
	MinSLARequirements int
	UploadPmt          uint64 // bid in marketplace, defaults to flat payment
	EthfsFilepath      string
	Filesize           uint64
	Shardsize          uint64
	FileContainerType  uint8
	EncryptionType     uint8
	CompressionType    uint8
	ShardContainerType uint8
	ErasureCodeType    uint8
	AccessControlLevel UploadAccessControlLevel
	CustomField        uint8

	ContainerSignature [ethcrypto.SignatureLength]byte
	SPsToUploadTo      [][20]byte // sps to whom the shards are to go
}

// To upload shards to sps, the uploader must make ProposeUpload tx to the
// smart contract with "SPsToUploadTo" being the result of the local
// marketplace instance with the upload as input. When the uploader sends
// the shards to the sps with the resultant txid, the sps will check that the
// txid has has data that matches the upload.. i.e. correct upload size,
// containerSignature, etc.
func ProposeUpload(params *UploadParams) (txid string, err error) {
	// construct tx
	nonce, nonce_err := utils.GetNonceForAddress(params.Wallet.Address)
	if nonce_err != nil {
		return "", nonce_err
	}
	amount := new(big.Int)
	amount.SetUint64(params.UploadPmt)
	height, h_err := utils.GetBlockHeight()
	if h_err != nil {
		return "", h_err
	}
	gasLimit, gl_err := utils.GetGasLimit(height)
	if gl_err != nil {
		return "", gl_err
	}
	r := strings.NewReader(ethfsAbi.Abi)
	scAbi, err_scAbi := abi.JSON(r) // reader io.Reader
	if err_scAbi != nil {
		return "", err_scAbi
	}
	bEthfsFilepath := []byte(params.EthfsFilepath)
	hashedEthfsFilepath := ethcrypto.Keccak256(bEthfsFilepath[:])
	proposeUploadParams := encodings.ProposeUploadParams{
		ServiceDuration:     params.ServiceDuration,
		MinSLARequirements:  params.MinSLARequirements,
		UploadPmt:           params.UploadPmt,
		Filesize:            params.Filesize,
		FileContainerType:   params.FileContainerType,
		EncryptionType:      params.EncryptionType,
		CompressionType:     params.CompressionType,
		ShardContainerType:  params.ShardContainerType,
		ErasureCodeType:     params.ErasureCodeType,
		ContainerSignatureV: params.ContainerSignature[64], // stashing V
		AccessControlLevel:  uint8(params.AccessControlLevel),
		CustomField:         params.CustomField}
	encodedParams, ep_err := encodings.EncodeProposeUploadParams(
		proposeUploadParams)
	if ep_err != nil {
		return "", ep_err
	}
	var bHashedEthfsFilepath [32]byte
	copy(bHashedEthfsFilepath[:], hashedEthfsFilepath[0:32])
	args := encodings.EthfsSCArgs{HashedEthfsFilepath: bHashedEthfsFilepath,
		ContainerSignature: params.ContainerSignature,
		Params:             encodedParams,
		Shardsize:          params.Shardsize,
		SPsToUploadTo:      params.SPsToUploadTo}
	var methodName string = "proposeUpload"
	dataFormatted, df_err := encodings.FormatData(scAbi, methodName, args)
	if df_err != nil {
		return "", df_err
	}
	contractAddressString := strings.Replace(ethfsAbi.ContractAddress(), "0x", "", 1)
	var contractAddress []byte
	for i := 0; i < len(contractAddressString); i += 2 {
		r, _ := hexutil.Decode("0x" + contractAddressString[i:i+2])
		contractAddress = append(contractAddress, []byte(r)...)
	}
	var bContractAddress [20]byte
	copy(bContractAddress[0:20], contractAddress[0:20])
	gasPrice, gp_err := utils.EstimateGas(params.Wallet.Address,
		bContractAddress,
		amount,
		dataFormatted)
	if gp_err != nil {
		return "", gp_err
	}
	accountHasEnoughEthers, balance, totalCost, err := utils.CheckTxCostAgainstBalance(params.UploadPmt, gasLimit, params.Wallet.Address)
	if err != nil {
		return "", err
	}
	if !accountHasEnoughEthers {
		return "", fmt.Errorf("error RegisterSP: totalCost of tx is ", totalCost, " but account balance is ", balance)
	}
	tx := types.NewTransaction(nonce,
		bContractAddress,
		amount,
		gasLimit,
		gasPrice,
		dataFormatted)
	// sign tx
	signedTx, err_signedTx := params.Wallet.SignTx(tx, height)
	if err_signedTx != nil {
		return "", err_signedTx
	}
	// send tx
	txid, tx_err := utils.SendRawTx(signedTx)
	if tx_err != nil {
		return "", tx_err
	}
	return txid, nil
}

func DoUpload(hash string, copyNum uint32, amount uint32) error {
	if hash == "" || copyNum < 1 {
		return errors.New("param value is error")
	}
	pdp, err := proof.NewProof(common.HexToAddress(constants.CONTRACT_ADDR), nil)
	if err != nil {
		log.Error("initialize new proof err:", err.Error())
		return err
	}
	sh := shell.NewLocalShell()
	err = sh.Push(hash, copyNum, false)
	if err != nil {
		log.Error("push err:", err.Error())
	} else {
		_, err = pdp.PledgeForFile(nil, amount, hash)
		if err != nil {
			log.Error("pledge for file error:", err.Error())
		}
		log.Error("success to push:", hash)
	}
	return err
}
