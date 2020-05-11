/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-11
 */
package main

import (
	"github.com/ipfs/go-ipfs/shell/ipfs"
	"os"
	"runtime"

	"github.com/nilhost/ipfs/nilfs/common/config"
	"github.com/urfave/cli"
)

func startEthfs()  {
	ipfs.MainStart()
}

func setupAPP() *cli.App {
	app := cli.NewApp()
	app.Usage = "Ethfs CLI"
	app.Action = startEthfs
	app.Version = config.Version
	app.Copyright = "Copyright in 2020 The ETHFS Authors"
	app.Commands = []cli.Command{
	}
	app.Flags = []cli.Flag{
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
