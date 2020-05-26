/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-11
 */
package handler

import (
	"github.com/Yihen/ethfs/core/downloader"
	"github.com/Yihen/ethfs/core/uploader"
)

func UploadData(params []interface{}) map[string]interface{} {
	if len(params) < 2 {
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
	copyNum, ok := params[1].(uint32)
	if !ok {
		return map[string]interface{}{
			"error":  20002,
			"desc":   "params(copyNum) type is ERROR",
			"result": "",
		}
	}
	if err := uploader.DoUpload(path, copyNum); err != nil {
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

func DownloadData(params []interface{}) map[string]interface{} {
	if len(params) < 1 {
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
			"desc":   "params type is ERROR",
			"result": "",
		}
	}
	if err := downloader.DoDownload(hash); err != nil {
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
	if len(params) < 1 {
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
			"desc":   "params type is ERROR",
			"result": "",
		}
	}
	if err := downloader.DoDownload(hash); err != nil {
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
func WithdrawToken(params []interface{}) map[string]interface{} {
	if len(params) < 1 {
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
			"desc":   "params type is ERROR",
			"result": "",
		}
	}
	if err := downloader.DoDownload(hash); err != nil {
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
