/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-11
 */
package handler

import (
	"fmt"

	"github.com/Yihen/ethfs/core/downloader"
	"github.com/Yihen/ethfs/core/token"
	"github.com/Yihen/ethfs/core/uploader"
)

func UploadData(params []interface{}) map[string]interface{} {
	if len(params) < 4 {
		return map[string]interface{}{
			"error":  20001,
			"desc":   "params is not enough",
			"result": "",
		}
	}
	path, ok := params[0].(string)
	if !ok {
		return map[string]interface{}{
			"error":  20002,
			"desc":   "params(path) type is ERROR",
			"result": "",
		}
	}
	fmt.Println("in upload, path:", path)
	copyNum, ok := params[1].(uint)
	if !ok {
		return map[string]interface{}{
			"error":  20002,
			"desc":   "params(copyNum) type is ERROR",
			"result": "",
		}
	}
	amount, ok := params[2].(uint)
	if !ok {
		return map[string]interface{}{
			"error":  20002,
			"desc":   "params(amount) type is ERROR",
			"result": "",
		}
	}
	fmt.Println("in upload, amount:", amount)
	pwd, ok := params[3].(string)
	if !ok {
		return map[string]interface{}{
			"error":  20002,
			"desc":   "params(pwd) type is ERROR",
			"result": "",
		}
	}
	fmt.Println("in upload, pwd:", pwd)
	if err := uploader.DoUpload(path, uint32(copyNum), uint32(amount), pwd); err != nil {
		return map[string]interface{}{
			"error":  20003,
			"desc":   "Download failed, path:" + path,
			"result": "",
		}
	}
	return map[string]interface{}{
		"error":  20000,
		"desc":   "Upload success",
		"result": "",
	}
}

func Login(params []interface{}) map[string]interface{} {
	if len(params) < 2 {
		return map[string]interface{}{
			"error":  20001,
			"desc":   "params is not enough",
			"result": "",
		}
	}
	hash, ok := params[0].(string)
	fmt.Println("hash:", hash)
	if !ok {
		return map[string]interface{}{
			"error":  20002,
			"desc":   "params hash is ERROR",
			"result": "",
		}
	}
	pwd, ok := params[1].(string)
	if !ok {
		return map[string]interface{}{
			"error":  20002,
			"desc":   "params pwd is ERROR",
			"result": "",
		}
	}

	fmt.Println("password:", pwd)

	return map[string]interface{}{
		"error":  20000,
		"desc":   "Open wallet success",
		"result": "",
	}
}
func DownloadData(params []interface{}) map[string]interface{} {
	if len(params) < 2 {
		return map[string]interface{}{
			"error":  20001,
			"desc":   "params is not enough",
			"result": "",
		}
	}
	hash, ok := params[0].(string)
	if !ok {
		return map[string]interface{}{
			"error":  20002,
			"desc":   "params hash is ERROR",
			"result": "",
		}
	}
	fmt.Println("in download, hash:", hash)
	pwd, ok := params[1].(string)
	if !ok {
		return map[string]interface{}{
			"error":  20002,
			"desc":   "params pwd is ERROR",
			"result": "",
		}
	}
	fmt.Println("in downlaod, pwd:", pwd)
	if err := downloader.DoDownload(hash, pwd); err != nil {
		return map[string]interface{}{
			"error":  20003,
			"desc":   "Download failed, hash:" + hash,
			"result": "",
		}
	}
	return map[string]interface{}{
		"error":  20000,
		"desc":   "Download success",
		"result": "",
	}
}

func PledgeToken(params []interface{}) map[string]interface{} {
	if len(params) < 3 {
		return map[string]interface{}{
			"error":  20001,
			"desc":   "params is not enough",
			"result": "",
		}
	}
	amount, ok := params[0].(uint)
	if !ok {
		return map[string]interface{}{
			"error":  20002,
			"desc":   "params hash is ERROR",
			"result": "",
		}
	}

	fmt.Println("in pledge, amount:", amount)
	pwd, ok := params[1].(string)
	if !ok {
		return map[string]interface{}{
			"error":  20002,
			"desc":   "params pwd is ERROR",
			"result": "",
		}
	}
	fmt.Println("in pledge, pwd:", pwd)

	address, ok := params[2].(string)
	if !ok {
		return map[string]interface{}{
			"error":  20002,
			"desc":   "params pwd is ERROR",
			"result": "",
		}
	}
	fmt.Println("in pledge, address:", address)

	if err := token.DoPledge(amount, pwd, address); err != nil {
		return map[string]interface{}{
			"error":  20003,
			"desc":   "pledge failed, address:" + address,
			"result": "",
		}
	}
	return map[string]interface{}{
		"error":  20000,
		"desc":   "pledge success",
		"result": "",
	}
}
func WithdrawToken(params []interface{}) map[string]interface{} {
	if len(params) < 1 {
		return map[string]interface{}{
			"error":  20001,
			"desc":   "params is not enough",
			"result": "",
		}
	}
	pwd, ok := params[0].(string)
	if !ok {
		return map[string]interface{}{
			"error":  20002,
			"desc":   "params pwd is ERROR",
			"result": "",
		}
	}
	fmt.Println("in quit, pwd:", pwd)
	if err := token.DoWithdraw(pwd); err != nil {
		return map[string]interface{}{
			"error":  20003,
			"desc":   "do withdraw failed",
			"result": "",
		}
	}
	return map[string]interface{}{
		"error":  20000,
		"desc":   "withdraw success",
		"result": "",
	}
}

func NodeStart(params []interface{}) map[string]interface{} {
	return map[string]interface{}{
		"error":  20000,
		"desc":   "node start mine success",
		"result": "",
	}
}

func NodeStop(params []interface{}) map[string]interface{} {
	return map[string]interface{}{
		"error":  20000,
		"desc":   "node stop mine success",
		"result": "",
	}
}

func GetUserInfo(params []interface{}) map[string]interface{} {
	return map[string]interface{}{
		"error":  20000,
		"desc":   "use info valid",
		"result": "",
	}
}

func GetMinerList(params []interface{}) map[string]interface{} {
	return map[string]interface{}{
		"error":  20000,
		"desc":   "miner list valid",
		"result": "",
	}
}
