package jsonrpc

import (
	"net/http"
	"strconv"

	"fmt"

	cfg "github.com/ETHFSx/ethfs/common/config"
	"github.com/ETHFSx/ethfs/common/log"
	"github.com/ETHFSx/ethfs/rpc/base"
)

func StartRPCServer() error {
	log.Debug()
	http.HandleFunc("/", base.Handle)

	//rpc.HandleFunc("getsysstatusscore", rpc.GetSysStatusScore)

	err := http.ListenAndServe(":"+strconv.Itoa(int(cfg.DefConfig.Rpc.HttpJsonPort)), nil)
	if err != nil {
		return fmt.Errorf("ListenAndServe error:%s", err)
	}
	return nil
}
