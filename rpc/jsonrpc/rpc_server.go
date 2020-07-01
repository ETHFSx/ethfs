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
	base.HandleFunc("api/v1/user/login", handler.Login)
	base.HandleFunc("download", handler.DownloadData)
	base.HandleFunc("upload", handler.UploadData)
	base.HandleFunc("quit", handler.WithdrawToken)
	base.HandleFunc("pledge", handler.PledgeToken)
	base.HandleFunc("start", handler.NodeStart)
	base.HandleFunc("stop", handler.NodeStop)
	base.HandleFunc("get_user_info", handler.GetUserInfo)
	base.HandleFunc("get_miner_list", handler.GetMinerList)
	//rpc.HandleFunc("getsysstatusscore", rpc.GetSysStatusScore)
	err := http.ListenAndServe(config.Parameters.PublicIP+":"+strconv.Itoa(int(config.Parameters.HttpJsonPort)), nil)
	if err != nil {
		return fmt.Errorf("ListenAndServe error:%s", err)
	}
	return nil
}
