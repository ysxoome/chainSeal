package tool

import (
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

// 上传文件到 IPFS 并返回 CID
func UploadFile(sh *shell.Shell, filePath string) (string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 上传文件到 IPFS
	cid, err := sh.Add(file)
	if err != nil {
		return "", err
	}

	return cid, nil
}

// 从 IPFS 下载文件到本地
func DownloadFile(sh *shell.Shell, cid, filePath string) error {

	// 从 IPFS 下载文件并写入输出文件
	err := sh.Get(cid, filePath)
	if err != nil {
		return err
	}

	return nil
}
