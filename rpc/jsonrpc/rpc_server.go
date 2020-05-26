package jsonrpc

import (
	"net/http"
	"strconv"

	"github.com/Yihen/ethfs/rpc/handler"

	"fmt"

	"github.com/Yihen/ethfs/common/config"
	"github.com/Yihen/ethfs/common/log"
	"github.com/Yihen/ethfs/rpc/base"
)

func StartRPCServer() error {
	log.Debug()
	http.HandleFunc("/", base.Handle)
	base.HandleFunc("download", handler.DownloadData)
	base.HandleFunc("upload", handler.UploadData)
	base.HandleFunc("withdraw", handler.WithdrawToken)
	base.HandleFunc("pledge", handler.PledgeToken)
	//rpc.HandleFunc("getsysstatusscore", rpc.GetSysStatusScore)
	err := http.ListenAndServe(config.Parameters.PublicIP+":"+strconv.Itoa(int(config.Parameters.HttpJsonPort)), nil)
	if err != nil {
		return fmt.Errorf("ListenAndServe error:%s", err)
	}
	return nil
}
