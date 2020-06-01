/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-11
 */
package downloader

import (
	"context"
	"errors"
	"time"

	"github.com/ETHFSx/go-ipfs/shell"

	proof "github.com/Yihen/contracts/dataproof/api"
	"github.com/Yihen/ethfs/common/constants"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Yihen/ethfs/common/log"
)

const DefaultKeepaliveInterval = 15 * time.Second

func DoDownload(hash string) error {
	if hash == "" {
		return errors.New("in downloader, param:hash value is empty")
	}
	pdp, err := proof.NewProof(common.HexToAddress(constants.CONTRACT_ADDR), nil)
	if err != nil {
		log.Error("in downloaer, initialize new proof err:", err.Error())
		return err
	}
	_, err = pdp.Challenge(nil, hash)
	if err != nil {
		log.Error("in download file, challenge error:", err.Error())
		return err
	}

	t := time.NewTicker(DefaultKeepaliveInterval)
	for {
		select {
		case <-t.C:
			tx, err := pdp.GetChallengeList(nil)
			// if challenge has been responsed, break else  do again in next time interval loop
			if err != nil {
				continue
			}

			sh := shell.NewLocalShell()
			peerInfo, err := sh.FindPeer(common.Bytes2Hex(tx.Data()))
			if err != nil {
				log.Error("in downloader, find peer is err:", err.Error())
				return err
			}
			err = sh.SwarmConnect(context.Background(), peerInfo.Addrs[0])
			if err != nil {
				log.Error("in downloader, swarm connect is err:", err.Error())
				return err
			}
			err = sh.Get(hash, "./")
			if err != nil {
				log.Error("in downloader, ipfs cat file err:", err.Error())
				return err
			} else {
				log.Error("success to download:", hash)
			}
			break
		}
	}

	return nil
}
