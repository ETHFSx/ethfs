/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-26
 */

package commands

import (
	"github.com/Yihen/ethfs/common/log"

	"github.com/Yihen/ethfs/common/constants"
	"github.com/ethereum/go-ethereum/common"

	proof "github.com/Yihen/contracts/dataproof/api"

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
			},
		},
	},
}

func doWithdraw(ctx *cli.Context) error {
	amount := ctx.Uint(utils.GetFlagName(utils.AmountFlag))
	log.Info("do withdraw in commands:", amount)
	pdp, err := proof.NewProof(common.HexToAddress(constants.CONTRACT_ADDR), nil)
	if err != nil {
		log.Error("initialize new proof err:", err.Error())
		return err
	}
	tx, err := pdp.NodeLeave(nil)
	if err != nil {
		log.Error("withdraw for node err:", err.Error())
		return err
	}
	log.Info("withdraw for node success, tx:", tx.Hash().String())
	return nil
}

func doPledge(ctx *cli.Context) error {
	address := ctx.String(utils.GetFlagName(utils.AddressFlag))
	amount := ctx.Uint(utils.GetFlagName(utils.AmountFlag))

	pdp, err := proof.NewProof(common.HexToAddress(constants.CONTRACT_ADDR), nil)
	if err != nil {
		log.Error("initialize new proof err:", err.Error())
		return err
	}
	tx, err := pdp.PledgeForNode(nil, uint32(amount), common.HexToAddress(address))
	if err != nil {
		log.Error("pledge for node err:", err.Error())
		return err
	}
	log.Info("pledge for node success, tx:", tx.Hash().String())

	return nil
}
