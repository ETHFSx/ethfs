package utils

import (
	ethfsAbi "github.com/ETHFSx/ethfs/account/abi"
)

var g_ethRpc = ethfsAbi.Rpc()

type Response struct {
	Result string `json:"result"`
}

func GetStorageAt(hexStoragePosition string) (Response, error) {
	var blockParameter string = "latest"
	var reqString string = "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getStorageAt\",\"params\": [\"" + ethfsAbi.ContractAddress() + "\", \"" + hexStoragePosition + "\", \"" + blockParameter + "\"],\"id\":1}"
	var reqBytes = []byte(reqString)
	return HttpPostWResponse(reqBytes)
}
