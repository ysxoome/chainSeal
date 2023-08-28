package service

import (
	"encoding/base64"
	"log"
	ContractEngine "seal/server/common/engine"
	"seal/server/common/models"
	"seal/server/common/utils"
	"seal/server/config"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

/* ========================================================== */

func TestAllSealAddress(t *testing.T) {
	setLog()

	ContractEngine.GetAllSealAddress()
}
func TestGenCompanyVC(t *testing.T) {
	setLog()

	company := models.CompanyVcEntity{
		LegalDid:    config.COMPANY_DID,
		Name:        "测试信息技术研究院有限公司",
		LegalName:   "张三",
		CreditCode:  "9132000073225110X",
		CompanyType: 1,
	}

	companyVcCredential := GenCompanyVC(company)
	// log.Println(companyVcCredential)
	companyVcCredential.WriteCredentialToFile("CVC.txt")

	sealResponseData := ContractEngine.CreateSeal(company.LegalDid, company.CreditCode)
	if !sealResponseData {
		log.Println("创建企业签章===失败===(已经创建也会报此错误)")
	}
	responseDate := ContractEngine.AddIdentity(companyVcCredential.Hash())
	if !responseDate {
		log.Println("创建企业VC上链===失败===")
	}

	isTrue := verify(&companyVcCredential)
	if !isTrue {
		log.Println("创建企业VC核验===失败===")
	}
	log.Println("企业签章地址:", CalculateSeal(company.LegalDid, company.CreditCode))
}

func TestApprovl(t *testing.T) {
	setLog()
	credentials, err := utils.CreateCredentials(config.COMPANY_PRIVATE_KEY)
	if err != nil {
		log.Println("createCredentials error:", err)
	}

	gid, _ := utils.AddrToGID(credentials.Address.String())
	signatureAuth := models.SignatureAuth{
		SealAddress: "0x07A17B46fCe6588A6ED1abE9cbf837b598B1eCB8",
		TypeAuth:    1,
		TargetDid:   "did:gid:83309d045a19c44dc3722d15a6abd472f95866ac",
		AuthDid:     gid,
		AuthTime:    time.Now().UnixMilli() / 1000,
	}
	encodePacked := ContractEngine.EncodePacked(signatureAuth.SealAddress,
		signatureAuth.TypeAuth,
		signatureAuth.TargetDid,
		signatureAuth.AuthTime)
	hash := crypto.Keccak256Hash(encodePacked)
	signatureData, err := crypto.Sign(hash[:], credentials.PrivateKey)
	if err != nil {
		log.Println("签名失败:", err)
	}

	signatureAuth.AuthSign = base64.StdEncoding.EncodeToString(signatureData)
	approvalSuccess, _ := ApproveSeal(signatureAuth)
	if !approvalSuccess {
		log.Println("签章授权失败")
	}

	// RevokeSeal(signatureAuth)
	// fmt.Println(signatureAuth)
}
