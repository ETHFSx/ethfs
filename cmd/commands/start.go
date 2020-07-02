/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-26
 */
package commands

import (
	"strings"
	"time"

	"github.com/Yihen/ethfs/common/constants"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ETHFSx/go-ipfs/shell"

	"github.com/Yihen/ethfs/cmd/commands/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ETHFSx/go-ipfs/shell/ipfs"
	proof "github.com/Yihen/contracts/dataproof/api"
	"github.com/Yihen/ethfs/common/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"
)

var stopCh chan struct{}

var StartCommand = cli.Command{
	Name:        "start",
	Usage:       "./ethfs start",
	Description: "start an ethfs node to response challenge of a file",
	Action:      start,
	Flags: []cli.Flag{
		utils.PasswordFlag,
	}}

var StopCommand = cli.Command{
	Name:        "start",
	Usage:       "./ethfs stop",
	Description: "stop an ethfs node to ",
	Action:      stop,
}

const DefaultKeepaliveInterval = 15 * time.Second

func doChallengeTask(pdp *proof.Proof, opts *bind.TransactOpts) error {
	tx, err := pdp.GetChallengeList(opts)
	if err != nil {
		log.Error("get challenge list err:", err.Error())
		return err
	}
	hash := common.Bytes2Hex(tx.Data())
	sh := shell.NewLocalShell()
	_, err = sh.Cat(hash)
	if err != nil {
		log.Errorf("cat file %s err:%s", hash, err.Error())
		return err
	}
	return nil
}

func start(ctx *cli.Context) error {
	password := ctx.String(utils.GetFlagName(utils.PasswordFlag))
	if 0 != ipfs.MainStart("daemon") {
		log.Error("start ipfs node ERROR")
	}
	conn, err := ethclient.Dial(constants.DEFAULT_ETH_WORKSPACE + "geth.ipc")
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

	t := time.NewTicker(DefaultKeepaliveInterval)
	go func() {
		for {
			select {
			case <-t.C:
				if err != doChallengeTask(pdp, auth) {
					break
				}
			case <-stopCh:
				log.Info("mine ethfs stop.")
				return
			}
		}
	}()
	log.Info("start ethfs node mine success")
	return nil
}

func stop(ctx *cli.Context) {
	close(stopCh)
}
