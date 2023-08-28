package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"

	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestGenPriKeyAndVerify(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	// fmt.Println(privateKey)

	msg := "hello, world"
	fmt.Printf("0x%x\n", crypto.Keccak256([]byte(msg)))
	sig, err := crypto.Sign(crypto.Keccak256([]byte(msg)), privateKey)
	if err != nil {
		panic(err)
	}
	// fmt.Println((sig))

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	// fmt.Println("publicKey: ", publicKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	valid := crypto.VerifySignature(publicKeyBytes, crypto.Keccak256([]byte(msg)), sig[:len(sig)-1]) // remove recovery id
	fmt.Println(valid)
}

func TestAddrToPublic(t *testing.T) {
	priKey, _ := crypto.GenerateKey()
	// 要验证的消息的哈希值
	messageHash := crypto.Keccak256Hash([]byte("Hello, World!")).Bytes()
	publicKey := crypto.FromECDSAPub(priKey.Public().(*ecdsa.PublicKey))
	fmt.Println(publicKey)
	// 签名的字节数组
	sig, _ := crypto.Sign(messageHash, priKey)

	// 验证签名并恢复公钥
	publicKey, err := crypto.Ecrecover(messageHash, sig)
	if err != nil {
		t.Fatal(err)
	}

	// 打印恢复的公钥
	fmt.Printf("Recovered public key:\n %v\n", publicKey)

	// verify
	vaild := crypto.VerifySignature(publicKey, messageHash, sig[:len(sig)-1])
	fmt.Println(vaild)
}

func TestPubKeyToAddr(t *testing.T) {
	pri, _ := crypto.GenerateKey()
	fmt.Println(pri)
	addr := crypto.PubkeyToAddress(*pri.Public().(*ecdsa.PublicKey))
	fmt.Println(addr)
}

func TestPubKeyWithSign(t *testing.T) {
	privateKey, _ := crypto.GenerateKey()

	msgHash := crypto.Keccak256([]byte("hello world"))

	// 对消息进行签名
	sig, _ := crypto.Sign(msgHash, privateKey)

	fmt.Println("签名原数据: ", sig)
	fmt.Println("签名数据: ", hexutil.Encode(sig))

	/* 验证签名并提取公钥 获取公钥方式的不同得到的公钥不同 需要统一方式 */
	pubKey, err := crypto.SigToPub(msgHash, sig)
	if err != nil {
		fmt.Println("无效的签名")
		return
	}

	address := crypto.PubkeyToAddress(*pubKey).Hex()
	fmt.Println("签名验证成功，地址为:", address)
}

func TestRSV(t *testing.T) {
	priKey, _ := crypto.GenerateKey()
	messageHash := crypto.Keccak256Hash([]byte("Hello, World!")).Bytes()

	// 签名的字节数组
	sig, _ := crypto.Sign(messageHash, priKey)
	// 获取 r 和 s 的值
	r, s := new(big.Int), new(big.Int)
	r.SetBytes(sig[:32])
	s.SetBytes(sig[32:64])

	// 获取 v 的值
	v := sig[64] + 27

	fmt.Println("r:", r.Bytes())
	fmt.Println("s:", s.Bytes())
	fmt.Println("v:", v)
}

func TestPem(t *testing.T) {
	// 读取 PEM 文件
	privateKeyByte, curveName, err := conf.LoadECPrivateKeyFromPEM("/Users/ysx/code/go/src/seal2go/blockchain/_ci/0x83309d045a19c44dc3722d15a6abd472f95866ac.pem")
	if err != nil {
		fmt.Println("无法读取 PEM 文件:", err)
		return
	}
	// fmt.Println("privateKeyByte:", privateKeyByte)
	fmt.Println("curveName:", curveName)

	message := []byte("Hello, world!")

	hash := crypto.Keccak256(message)

	privateKey, _ := crypto.ToECDSA(privateKeyByte)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])

	v := new(big.Int).SetUint64(27)
	if s.Bit(0) == 1 {
		v.Add(v, big.NewInt(1))
	}
	fmt.Println(r, s, v)

	if err != nil {
		fmt.Println("签名失败:", err)
		return
	}
	signatureESDA := append(r.Bytes(), s.Bytes()...)

	fmt.Printf("签名结果：%x\n", signatureESDA)

	// signatureByte, _ := crypto.Sign(hash[:], privateKey)
	// fmt.Printf("签名结果：%x\n", signatureByte)

	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	t.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	// }

	// publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// fmt.Println("publick key: ", hexutil.Encode(publicKeyBytes)[4:]) // publicKey in hex without "0x"

	// address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	// fmt.Println("address: ", address) // account address 0xbEE881DD425a1C5659061cc51E3Eb7fDA267cE11

}
