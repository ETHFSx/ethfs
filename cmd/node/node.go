/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-11
 */
package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Yihen/ethfs/cmd/commands/utils"

	"github.com/ETHFSx/go-ipfs/shell/ipfs"
	"github.com/Yihen/ethfs/cmd/commands"

	"github.com/Yihen/ethfs/common/config"
	"github.com/Yihen/ethfs/rpc/jsonrpc"
	"github.com/urfave/cli"
)

func startRPCServer() {
	if err := jsonrpc.StartRPCServer(); err != nil {
		fmt.Println("start rpc server error:", err.Error())
	}
}

func startEthfs(ctx *cli.Context) {
	config.InitConfig()
	go ipfs.MainStart("daemon")
	go startRPCServer()
	select {}
}

func setupAPP() *cli.App {
	app := cli.NewApp()
	app.Usage = "Ethfs CLI"
	app.Action = startEthfs
	app.Version = config.Version
	app.Copyright = "Copyright in 2020 The ETHFS Authors"
	app.Commands = []cli.Command{
		commands.DataCommand,
		commands.TokenCommand,
		commands.StartCommand,
		commands.StopCommand,
	}
	app.Flags = []cli.Flag{
		utils.RPCPortFlag,
		utils.CopyNumFlag,
		utils.HashFlag,
		utils.PathFlag,
		utils.AddressFlag,
		utils.AmountFlag,
		utils.PasswordFlag,
	}
	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}

func main() {
	if err := setupAPP().Run(os.Args); err != nil {
		os.Exit(1)
	}
}
