package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"seal/blockchain/tool"
	"seal/server/common/models"
	"strings"
	"time"

	"crypto/ecdsa"

	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/crypto"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/pkg/errors"
)

func DeSerializePri(priPath string) *ecdsa.PrivateKey {
	privateKeyByte, _, err := conf.LoadECPrivateKeyFromPEM(priPath)
	if err != nil {
		fmt.Println("无法读取 PEM 文件:", err)
		return &ecdsa.PrivateKey{}
	}
	privateKey, _ := crypto.ToECDSA(privateKeyByte)

	return privateKey
}

func StringToSig(data string) (sig []byte) {
	sig, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		errors.New("string to sig error")
	}

	return
}

func priKeyToPubAddr(priPath string) (addr string) {
	pri := DeSerializePri(priPath)
	addr = crypto.PubkeyToAddress(*pri.Public().(*ecdsa.PublicKey)).String()
	return
}

func verifyTime(issueDate, expire string) bool {
	layout := "2006-01-02 15:04:05.999999"
	t1, _ := time.Parse(layout, issueDate)
	t2, _ := time.Parse(layout, expire)
	return t1.Before(t2)
}

func VerifyPrefixedSignatureFromWeId(rawData []byte, signature, did string) bool {

	sigData, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		fmt.Println("signature failed or signature format wrong")
		return false
	}

	return verifyPrefixedSignature(rawData, sigData, did)
}

func verifyPrefixedSignature(rawByte []byte, sigData []byte, did string) bool {
	if rawByte == nil {
		fmt.Println("sigData is null")
		return false
	}

	pubKeyByte, err := crypto.Ecrecover(rawByte, sigData)
	if err != nil {
		fmt.Println("extract publicKey from signature failed")
		return false
	}

	publicKey, err := crypto.UnmarshalPubkey(pubKeyByte)
	if err != nil {
		fmt.Println("UnmarshalPubkey failed")
		return false
	}

	addr := crypto.PubkeyToAddress(*publicKey)
	extractedDid, _ := AddrToGID(addr.String())

	return strings.EqualFold(extractedDid, did)
}

func ContractCID(shell *shell.Shell, c models.Contract) (CID string, err error) {
	fileName := fmt.Sprintf("./%s.txt", c.Hash())
	c.WriteContractToFile(fileName)
	CID, err = tool.UploadFile(shell, fileName)
	if err != nil {
		return "", errors.Wrap(err, "file upload failed")
	}
	os.Remove(fileName)

	return CID, nil
}

func ContractFromCid(shell *shell.Shell, cid string) (*models.Contract, error) {
	fileName := fmt.Sprintf("./%s.txt", cid)
	tool.DownloadFile(shell, cid, fileName)
	var c models.Contract
	data, _ := os.ReadFile(fileName)
	err := json.Unmarshal(data, &c)
	if err != nil {
		return &models.Contract{}, errors.Wrap(err, "file format wrong")
	}
	os.Remove(fileName)
	return &c, nil
}

func CredentialCID(shell *shell.Shell, c models.Credential) (CID string, err error) {
	fileName := fmt.Sprintf("./%s.txt", c.Hash())
	c.WriteCredentialToFile(fileName)
	CID, err = tool.UploadFile(shell, fileName)
	if err != nil {
		return "", errors.Wrap(err, "file upload failed")
	}
	os.Remove(fileName)

	return CID, nil
}

func CredentialFromCID(shell *shell.Shell, cid string) (*models.Credential, error) {
	fileName := fmt.Sprintf("./%s.txt", cid)
	tool.DownloadFile(shell, cid, fileName)
	var c models.Credential
	data, _ := os.ReadFile(fileName)
	err := json.Unmarshal(data, &c)
	if err != nil {
		return &models.Credential{}, errors.Wrap(err, "file format wrong")
	}
	os.Remove(fileName)
	return &c, nil
}

/* ========================================================== */
