/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-26
 */
package commands

import (
	"time"

	"github.com/ETHFSx/go-ipfs/shell/ipfs"
	proof "github.com/Yihen/contracts/dataproof/api"
	"github.com/Yihen/ethfs/common/constants"
	"github.com/Yihen/ethfs/common/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"
)

var StartCommand = cli.Command{
	Name:        "start",
	Usage:       "./ethfs start",
	Description: "start an ethfs node to response challenge of a file",
	Action:      start,
}

const DefaultKeepaliveInterval = 15 * time.Second

func doChallengeTask(pdp *proof.Proof) error {
	tx, err := pdp.GetChallengeList(nil)
	if err != nil {
		log.Error("get challenge list err:", err.Error())
		return err
	}
	tx.Data()
	return nil
}

func start(ctx *cli.Context) error {
	if 0 != ipfs.MainStart("daemon") {
		log.Error("start ipfs node ERROR")
	}
	pdp, err := proof.NewProof(common.HexToAddress(constants.CONTRACT_ADDR), nil)
	if err != nil {
		log.Error("initialize new proof err:", err.Error())
		return err
	}

	t := time.NewTicker(DefaultKeepaliveInterval)
	go func() {
		for {
			select {
			case <-t.C:
				if err != doChallengeTask(pdp) {
					break
				}
			}
		}
	}()
	return nil
}
