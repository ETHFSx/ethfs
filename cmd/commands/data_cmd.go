/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-21
 */
package commands

import (
	"github.com/Yihen/ethfs/common/log"

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
			Usage:       "./ethfs data upload [arguments...]",
			ArgsUsage:   "[arguments...]",
			Description: "Upload data from local to ethfs network, you need to special file location.",
			Flags: []cli.Flag{
				utils.PathFlag,
				utils.CopyNumFlag,
				utils.AmountFlag,
				utils.PasswordFlag,
			},
		},
		{
			Action:      doDownload,
			Name:        "download",
			Usage:       "./ethfs data download [arguments...]",
			ArgsUsage:   "[arguments...]",
			Description: "Download data from ethfs network, this maybe need to wait for a moment before beginning loading as for contract to be verified",
			Flags: []cli.Flag{
				utils.HashFlag,
				utils.PasswordFlag,
			},
		},
	},
}

func doUpload(ctx *cli.Context) error {
	path := ctx.String(utils.GetFlagName(utils.PathFlag))
	copyNum := ctx.Uint(utils.GetFlagName(utils.CopyNumFlag))
	amount := ctx.Uint(utils.GetFlagName(utils.AmountFlag))
	password := ctx.String(utils.GetFlagName(utils.PasswordFlag))
	log.Info("do upload in commands:", path, copyNum, amount)
	go ipfs.MainStart("daemon")
	if err := uploader.DoUpload(path, uint32(copyNum), uint32(amount), password); err != nil {
		log.Error("upload err:", err, ",path:", path, ",copy number:", copyNum)
	}
	return nil
}

func doDownload(ctx *cli.Context) error {
	hash := ctx.String(utils.GetFlagName(utils.HashFlag))
	password := ctx.String(utils.GetFlagName(utils.PasswordFlag))
	log.Info("do download commands:", hash)
	go ipfs.MainStart("daemon")
	if err := downloader.DoDownload(hash, password); err != nil {
		log.Error("download err:", err, " ,hash:", hash)
	}
	return nil
}
