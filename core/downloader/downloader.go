/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-11
 */
package downloader

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ETHFSx/go-ipfs/shell"

	proof "github.com/Yihen/contracts/dataproof/api"
	"github.com/Yihen/ethfs/common/constants"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Yihen/ethfs/common/log"
)

const DefaultKeepaliveInterval = 15 * time.Second

func DoDownload(hash, pwd string) error {
	if hash == "" {
		return errors.New("in downloader, param:hash value is empty")
	}

	conn, err := ethclient.Dial(constants.DEFAULT_ETH_WORKSPACE + "geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(constants.ACCOUNT_KEY), pwd)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	pdp, err := proof.NewProof(common.HexToAddress(constants.CONTRACT_ADDR), conn)
	if err != nil {
		log.Error("in downloaer, initialize new proof err:", err.Error())
		return err
	}
	_, err = pdp.Challenge(auth, hash)
	if err != nil {
		log.Error("in download file, challenge error:", err.Error())
		return err
	}

	t := time.NewTicker(DefaultKeepaliveInterval)
	for {
		select {
		case <-t.C:
			tx, err := pdp.GetChallengeList(auth)
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
			err = sh.Get(hash, "./"+hash+".download")
			if err != nil {
				log.Error("in downloader, ipfs cat file err:", err.Error())
				return err
			} else {
				log.Error("success to download:", hash)
			}

			fp, err := os.Open("./" + hash + ".download")
			if err != nil {
				return err
			}
			fileInfo, err := fp.Stat()
			if err != nil {
				return err
			}
			buffer := make([]byte, fileInfo.Size())
			fmt.Println("file.size:", fileInfo.Size())
			_, err = fp.Read(buffer)
			if err != nil {
				return err
			}
			Sha1Inst := sha1.New()
			Sha1Inst.Write(buffer)
			fileHash := Sha1Inst.Sum([]byte(""))
			_, err = pdp.PunishOrReward(auth, common.HexToAddress(peerInfo.Addrs[0]), hash, common.Bytes2Hex(fileHash) == hash)
			if err != nil {
				return err
			}
			break
		}
	}

	return nil
}
