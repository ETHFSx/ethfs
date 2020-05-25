/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-21
 */
package commands

import (
	"fmt"

	"github.com/Yihen/ethfs/core/downloader"

	"github.com/Yihen/ethfs/core/uploader"

	"github.com/ETHFSx/go-ipfs/shell/ipfs"

	"github.com/Yihen/ethfs/cmd/commands/utils"

	"github.com/urfave/cli"
)

var DataCommand = cli.Command{
	Name:        "data",
	Usage:       "Handle assets",
	Description: "Asset management commands can check account balance, USDT transfers, and so on.",
	Subcommands: []cli.Command{
		{
			Action:      doUpload,
			Name:        "upload",
			Usage:       "./node data upload [arguments...]",
			ArgsUsage:   "[arguments...]",
			Description: "Upload data from local to ethfs network, you need to special file location.",
			Flags: []cli.Flag{
				utils.PathFlag,
				utils.CopyNumFlag,
			},
		},
		{
			Action:      doDownload,
			Name:        "download",
			Usage:       "./node data download [arguments...]",
			ArgsUsage:   "[arguments...]",
			Description: "Download data from ethfs network, this maybe need to wait for a moment before beginning loading as for contract to be verified",
			Flags: []cli.Flag{
				utils.HashFlag,
			},
		},
	},
}

func doUpload(ctx *cli.Context) error {
	path := ctx.String(utils.GetFlagName(utils.PathFlag))
	copyNum := ctx.Uint(utils.GetFlagName(utils.CopyNumFlag))
	fmt.Println("do upload in commands:", path, copyNum)
	go ipfs.MainStart("daemon")
	if err := uploader.DoUpload(path, uint32(copyNum)); err != nil {
		fmt.Println("upload err:", err, ",path:", path, ",copy number:", copyNum)
	}
	return nil
}

func doDownload(ctx *cli.Context) error {
	hash := ctx.String(utils.GetFlagName(utils.HashFlag))
	fmt.Println("do download commands:", hash)
	go ipfs.MainStart("daemon")
	if err := downloader.DoDownload(hash); err != nil {
		fmt.Println("download err:", err, " ,hash:", hash)
	}
	return nil
}
