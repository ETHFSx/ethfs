package localrpc

import (
	"net/http"
	"strconv"

	"fmt"

	cfg "github.com/Yihen/ethfs/common/config"
	"github.com/Yihen/ethfs/common/log"
)

const (
	LOCAL_HOST string = "127.0.0.1"
	LOCAL_DIR  string = "/local"
)

func StartLocalServer() error {
	log.Debug()
	http.HandleFunc(LOCAL_DIR, rpc.Handle)

	// TODO: only listen to local host
	err := http.ListenAndServe(":"+strconv.Itoa(int(cfg.DefConfig.Rpc.HttpLocalPort)), nil)
	if err != nil {
		return fmt.Errorf("ListenAndServe error:%s", err)
	}
	return nil
}
