/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-21
 */
package utils

import (
	"strings"

	"github.com/urfave/cli"
)

var (
	RPCPortFlag = cli.UintFlag{
		Name:  "rpcport",
		Usage: "Json rpc server listening port `<number>`",
		Value: 6000,
	}
	PathFlag = cli.StringFlag{
		Name:  "path",
		Usage: "uploading file path",
		Value: "",
	}
	CopyNumFlag = cli.UintFlag{
		Name:  "copynum",
		Usage: "copy number for uploading data distributed",
		Value: 3,
	}
	HashFlag = cli.StringFlag{
		Name:  "hash",
		Usage: "hash value for downloading data file",
		Value: "",
	}
)

//GetFlagName deal with short flag, and return the flag name whether flag name have short name
func GetFlagName(flag cli.Flag) string {
	name := flag.GetName()
	if name == "" {
		return ""
	}
	return strings.TrimSpace(strings.Split(name, ",")[0])
}
