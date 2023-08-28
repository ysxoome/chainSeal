package service

import (
	"fmt"
	"math/big"
	"seal/blockchain/tool"
	"seal/server/common/constant"
	"seal/server/common/engine"
	"seal/server/common/models"
	"seal/server/common/utils"
	"seal/server/common/xerr"
	"seal/server/config"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jinzhu/copier"
)

func GenPersonVC(per models.PersonVC) models.Credential {
	keyStore := per.KeyStore
	did, _ := utils.AddrToGID(keyStore.Address.String())

	credential := models.Credential{
		Context:        constant.DEFAULT_CREDENTIAL_CONTEXT,
		Id:             utils.RandomUUID(),
		Issuer:         config.ADMIN_DID,
		IssuanceDate:   time.Now().String()[:19],
		ExpirationDate: time.Now().Add(time.Hour * 24 * 365 * 10).String()[:19],
		Claim: map[string]string{
			constant.PERSON_DID:  did,
			constant.PERSON_TYPE: config.ID_CARD,
			constant.PERSON_NAME: per.Name,
		},
	}
	credential.Proof = (utils.GenCredentialProof(credential, *utils.DeSerializePri(config.ADMIN_PRIVATE_KEY), nil))

	return credential
}

/**
 *	创建企业Vc
 *	部署企业签章合约
 *
 */
func GenCompanyVC(company models.CompanyVcEntity) models.Credential {

	companyAddress := CalculateSeal(company.LegalDid, company.CreditCode)
	companyDid, _ := utils.AddrToGID(companyAddress)

	credential := models.Credential{
		Context:        constant.DEFAULT_CREDENTIAL_CONTEXT,
		Id:             utils.RandomUUID(),
		Issuer:         config.ADMIN_DID,
		IssuanceDate:   time.Now().String()[:19],
		ExpirationDate: time.Now().Add(time.Hour * 24 * 365 * 10).String()[:19],
		Claim: map[string]string{
			constant.COMPANY_DID:       companyDid,
			constant.COMPANY_TYPE:      fmt.Sprint(company.CompanyType),
			constant.COMPANY_NAME:      company.Name,
			constant.COMPANY_CODE:      company.CreditCode,
			constant.COMPANY_LEGAL:     company.LegalName,
			constant.COMPANY_LEGAL_DID: company.LegalDid,
		},
	}
	credential.Proof = (utils.GenCredentialProof(credential, *utils.DeSerializePri(config.ADMIN_PRIVATE_KEY), nil))

	return credential
}

func verify(credential *models.Credential) bool {
	// 验证vc格式、有效期
	err := utils.IsCredentialVaild(credential)
	if err.GetErrCode() != xerr.OK {
		fmt.Println("credential input format error!")
		return false
	}
	// 验证链上issuer是否有效
	exist := verifyIssuerExistence(credential.Issuer)
	if !exist {
		return false
	}

	// 验证签名
	vaild := verifySignature(credential)
	if !vaild {
		return false
	}

	// 验证是否吊销 有效true无效false
	inTime := verifyIdetityTime(credential.Hash())

	return inTime
}

func verifyIdetityTime(hash string) bool {
	identitySession := tool.InitIdentity()
	hashByte, _ := hexutil.Decode(hash)
	vaild, err := identitySession.QueryIdentity(hashByte, big.NewInt(time.Now().UnixNano()))
	if err != nil {
		fmt.Println("identity contract QueryIdentity failed")
		return false
	}
	return vaild
}

func verifyIssuerExistence(issuer string) bool {
	identitySession := tool.InitIdentity()
	owner, err := identitySession.Owner()
	if err != nil || strings.ToLower(owner.String()) != utils.GidToAddr(issuer) {
		fmt.Println("issuer has no permission")
		return false
	}
	return true
}

func verifySignature(credential *models.Credential) bool {
	disclosureMap := make(map[string]string)
	copier.Copy(&disclosureMap, credential.Claim)

	issuerDid := credential.Issuer
	rawData := utils.GenCredentialThumbprintWithoutSig(*credential, disclosureMap)
	sig := utils.StringToSig(credential.Proof["signature"])
	hash := crypto.Keccak256([]byte(rawData))
	publicKey, err := crypto.Ecrecover(hash, sig)
	if err != nil {
		fmt.Println("get publicKey failed")
		return false
	}
	p, _ := crypto.UnmarshalPubkey(publicKey)
	addr := crypto.PubkeyToAddress(*p)
	issue, _ := utils.AddrToGID(addr.String())
	if !strings.EqualFold(issue, issuerDid) {
		fmt.Println("gid wrong")
		return false
	}
	valid := crypto.VerifySignature(publicKey, hash, sig[:(len(sig))-1])
	if !valid {
		fmt.Println("verify failed")
		return false
	}
	return true
}

func verifyByTime(credential models.Credential, time *big.Int) bool {
	// 验证vc格式、有效期
	err := utils.IsCredentialVaild(&credential)
	if err.GetErrCode() != xerr.OK {
		fmt.Println("verifyByTime error!")
		return false
	}
	// 验证签名
	vaild := verifySignature(&credential)

	// 验证是否吊销
	fmt.Println(credential.Hash(), time)
	inTime := engine.QueryIdentity(credential.Hash(), time)
	return vaild && inTime
}
