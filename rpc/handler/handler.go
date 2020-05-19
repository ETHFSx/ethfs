/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-11
 */
package handler

import "fmt"

func UploadData(params []interface{}) map[string]interface{} {
	fmt.Println("Hello Upload")
	return map[string]interface{}{
		"error":  20000,
		"desc":   "Upload success",
		"result": "",
	}
}


func DownloadData(params []interface{}) map[string]interface{} {
	fmt.Println("Hello Download")
	return map[string]interface{}{
		"error":  20000,
		"desc":   "Download success",
		"result": "",
	}
}
