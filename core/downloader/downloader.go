/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-11
 */
package downloader

import (
	"fmt"

	"github.com/ETHFSx/go-ipfs/shell"
)

func DoDownload(hash string) error {
	sh := shell.NewLocalShell()
	err := sh.Get(hash, "./")
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("success to download:", hash)
	}
	return nil
}
