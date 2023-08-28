package sh

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {

	filepath := "get_account.sh"
	// 检查文件是否有可执行权限
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	mode := fileInfo.Mode()
	if mode&0100 != 0 {
		fmt.Println("文件有可执行权限")
	} else {
		//给文件增加可执行权限
		err := os.Chmod(filepath, 0700)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("文件增加可执行权限成功")
	}

}

/**
* @descriton go执行创建账户脚本
 */
func RunbashGetAccount() {

	// 返回一个 cmd 对象,相对路径
	cmd := exec.Command("sh", "-c", "./get_account.sh")

	// 收返回值[]byte, error
	b, err := cmd.Output()
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	} else {
		fmt.Println("success")
		fmt.Println(string(b))
	}
}

// 指定PEM私钥文件计算账户地址
func GetAccountAddress(filepath string) {
	cmd := exec.Command("bash", "get_account.sh", "-k ", filepath)
	b, err := cmd.Output()
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	} else {
		fmt.Println("success")
		fmt.Println(string(b))
	}
}
