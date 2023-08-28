package sh

import "testing"

/**
* @descriton go执行创建账户脚本
 */
func TestGetAccount(t *testing.T) {
	//在当前目录下执行，生成account/pem私钥和公钥文件
	RunbashGetAccount()

	//指定PEM私钥文件计算账户地址
	GetAccountAddress("./accounts")
}