package service

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"seal/blockchain/tool"
	"seal/server/common/engine"
	"seal/server/common/utils"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	shell "github.com/ipfs/go-ipfs-api"
)

func TestIPFS(t *testing.T) {
	// contract := GenSign(nil)
	// fmt.Println(contract)
	// contract.WriteContractToFile("1.txt")
	// fmt.Println(tool.UploadFile(shell.NewShell("http://192.168.103.114:8081"), "1.txt"))
	// QmYZs2PR1xavjxrjPboAjGSTt5wo4GLRsm78uXhvNReuuV
	tool.DownloadFile(shell.NewShell("http://192.168.103.114:8081"), "QmTfDSqLzaPnfGnwfviat8dP2gmBSsiPhGL4h6SvZ3A8nD", "./1.json")
}

func TestGenSign(t *testing.T) {
	contract := GenSign(nil)
	c := GenSign(&contract)
	// fmt.Println(c)
	// cc := GenSign(&c)
	// ccc := GenSign(&cc)

	// contract.WriteContractToFile("1.txt")
	ToChain(c)
	// cc.WriteContractToFile("1.txt")
	// d, _ := utils.ContractCID(shell.NewShell("http://192.168.103.114:8081"),ccc)
	// fmt.Println(d)
	// vaild := toChain(contract)
	// fmt.Println(vaild)
}

func TestRead(t *testing.T) {
	c, err := utils.ContractFromCid(shell.NewShell("http://192.168.103.114:8081"), "QmYZs2PR1xavjxrjPboAjGSTt5wo4GLRsm78uXhvNReuuV")
	if err != nil {
		t.Fatal(err)
	}
	c.WriteContractToFile("2.txt")
}

func TestContract(t *testing.T) {
	contract := GenSign(nil)
	vaild, _, _, _, _ := ToChain(contract)
	fmt.Println(vaild)
	// c := GenSign([]*models.Contract{&contract})
	// fmt.Println(c)
	// cc := GenSign([]*models.Contract{&c,&contract})
	// fmt.Println(cc.GetNestContract())
}

func TestQueryDetail(t *testing.T) {
	res := engine.QueryDetail("0xce779b350231c5eb8b1bba07167f8760f1312c5ef2c2226dd4fd488a2eb7c7f9", 0)
	fmt.Println(res)
	b := []byte{198, 53, 59, 112, 165, 169, 2, 82, 209, 255, 241, 6, 233, 142, 50, 180, 112, 58, 253, 46, 29, 200, 160, 176, 177, 36, 53, 19, 39, 211, 95, 241}
	// 得到的第二个字节数组是signHash的字节数组
	fmt.Println(hexutil.Encode(b))
	// b, _ = hexutil.Decode("0xde4f87d3456d9ea9b01abe8b15d60c6c9e9a9127")
	b, _ = hexutil.Decode("0x9fdaa0ec5e80223fc5382a18e4d7a8b12cc67a2fc779eef1b60efeb589449fae")
	fmt.Println(b)
	// 0xde4f87d3456d9ea9b01abe8b15d60c6c9e9a9127
	// [198 53 59 112 165 169 2 82 209 255 241 6 233 142 50 180 112 58 253 46 29 200 160 176 177 36 53 19 39 211 95 241]

}

func TestQuery(t *testing.T) {
	// query查询下标
	// queryDetail查询具体信息 用到下标
	chash := "0xce779b350231c5eb8b1bba07167f8760f1312c5ef2c2226dd4fd488a2eb7c7f9"
	sHash := "0xf4d646f1e90fe12d8d009f6290cfe8d01fe70fbb2b7a236f5f42a652794eb8d1"
	sealAddr := "0x07A17B46fCe6588A6ED1abE9cbf837b598B1eCB8"
	sealType := 1
	signAddr := "0x83309d045a19c44Dc3722D15A6AbD472f95866aC"

	index := engine.Query(chash, sHash, sealAddr, uint8(sealType), signAddr)
	fmt.Println(index)
	res := engine.QueryDetail(chash, index.Uint64()-1)
	fmt.Println(res)
	v, _ := res.(struct {
		Cid      string
		SignHash []byte
		Sealaddr common.Address
		SealType uint8
		SignAddr common.Address
		SignTime *big.Int
	})

	fmt.Println(base64.StdEncoding.EncodeToString(v.SignHash))

	fmt.Println((v.SignHash))
	fmt.Println(hexutil.Encode(v.SignHash))
	fmt.Println(v.SignTime)
}

func TestVerifyOnChain(t *testing.T) {
	contract := GenSign(nil)
	vaild, _, _, _, _ := ToChain(contract)
	if vaild {
		fmt.Println(verifyFileOnChain(contract))
	} else {
		fmt.Println("上链失败")
	}
}
