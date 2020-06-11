/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-06-11
 */
package token

import (
	"strings"

	"github.com/Yihen/ethfs/common/log"

	proof "github.com/Yihen/contracts/dataproof/api"
	"github.com/Yihen/ethfs/common/constants"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DoWithdraw(password string) error {
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

func DoPledge(amount uint, password, address string) error {
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
