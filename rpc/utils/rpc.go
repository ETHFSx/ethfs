package utils

import (
	"fmt"

	ethfsAbi "github.com/Yihen/ethfs/account/abi"
)

func SetRpcUrl(r []string) error {
	rpcInitializer := ethfsAbi.SetRpcUrl(r)
	return SanitizeRpcUrls(rpcInitializer)
}

func SanitizeRpcUrls(rpcPtr *ethfsAbi.UrlPtr) error {
	var updatedUrls []string
	for i := 0; i < len(rpcPtr.Urls); i++ {
		rpcPtr.Url = rpcPtr.Urls[i]
		_, err := GetBlockHeight()
		if err == nil {
			updatedUrls = append(updatedUrls, rpcPtr.Urls[i])
		}
	}
	if len(updatedUrls) < 1 {
		return fmt.Errorf("Error SanitizeRpcUrls: all rpc urls invalid.")
	}
	rpcPtr.Urls = updatedUrls
	rpcPtr.Url = updatedUrls[0]
	return nil
}
