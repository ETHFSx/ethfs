/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-26
 */

package commands

import (
	"strings"

	"github.com/Yihen/ethfs/common/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

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
	conn, err := ethclient.Dial("~/.ethereum/geth.ipc")
	if err != nil {
		log.Fatalf("in withdraw, failed to connect to the Ethereum client: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(constants.ACCOUNT_KEY), password)
	if err != nil {
		log.Fatalf("in withdraw, failed to create authorized transactor: %v", err)
	}
	pdp, err := proof.NewProof(common.HexToAddress(constants.CONTRACT_ADDR), conn)
	if err != nil {
		log.Error("initialize new proof err:", err.Error())
		return err
	}
	tx, err := pdp.NodeLeave(auth)
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
	password := ctx.String(utils.GetFlagName(utils.PasswordFlag))
	conn, err := ethclient.Dial("~/.ethereum/geth.ipc")
	if err != nil {
		log.Fatalf("in pledge, failed to connect to the Ethereum client: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(constants.ACCOUNT_KEY), password)
	if err != nil {
		log.Fatalf("in pledge, failed to create authorized transactor: %v", err)
	}
	pdp, err := proof.NewProof(common.HexToAddress(constants.CONTRACT_ADDR), conn)
	if err != nil {
		log.Error("initialize new proof err:", err.Error())
		return err
	}
	tx, err := pdp.PledgeForNode(auth, uint32(amount), common.HexToAddress(address))
	if err != nil {
		log.Error("pledge for node err:", err.Error())
		return err
	}
	log.Info("pledge for node success, tx:", tx.Hash().String())

	return nil
}
