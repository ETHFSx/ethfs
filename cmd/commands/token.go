/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-26
 */

package commands

import (
	"github.com/Yihen/ethfs/core/token"

	"github.com/Yihen/ethfs/common/log"

	"github.com/Yihen/ethfs/cmd/commands/utils"
	"github.com/urfave/cli"
)

var TokenCommand = cli.Command{
	Name:        "token",
	Usage:       "Handle assets",
	Description: "Asset management commands can check account balance, USDT transfers, and so on.",
	Subcommands: []cli.Command{
		{
			Action:      doWithdraw,
			Name:        "withdraw",
			Usage:       "./ethfs token withdraw [arguments...]",
			ArgsUsage:   "[arguments...]",
			Description: "When node leave the storage network, he can choose to withdraw the ETH token pledge at first.",
			Flags: []cli.Flag{
				utils.AmountFlag,
				utils.PasswordFlag,
			},
		},
		{
			Action:      doPledge,
			Name:        "pledge",
			Usage:       "./ethfs token pledge [arguments...]",
			ArgsUsage:   "[arguments...]",
			Description: "When a node want to join the storage network to suply service, he need to pledge some ETH token; as well, if one want to store data into network, need to pledge firstly.if one want to store data into network, need to pledge firstly.if one want to store data into network, need to pledge firstly.if one want to store data into network, need to pledge firstly.if one want to store data into network, need to pledge firstly.if one want to store data into network, need to pledge firstly.if one want to store data into network, need to pledge firstly.if one want to store data into network, need to pledge firstly.if one want to store data into network, need to pledge firstly.",
			Flags: []cli.Flag{
				utils.AddressFlag,
				utils.AmountFlag,
				utils.PasswordFlag,
			},
		},
	},
}

func doWithdraw(ctx *cli.Context) error {
	amount := ctx.Uint(utils.GetFlagName(utils.AmountFlag))
	password := ctx.String(utils.GetFlagName(utils.PasswordFlag))
	log.Info("do withdraw in commands:", amount)
	return token.DoWithdraw(password)
}

func doPledge(ctx *cli.Context) error {
	address := ctx.String(utils.GetFlagName(utils.AddressFlag))
	amount := ctx.Uint(utils.GetFlagName(utils.AmountFlag))
	password := ctx.String(utils.GetFlagName(utils.PasswordFlag))
	return token.DoPledge(amount, password, address)
}
