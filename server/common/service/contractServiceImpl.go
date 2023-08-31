package service

import (
	"encoding/base64"
	"fmt"
	"seal/blockchain/tool"
	"seal/server/common/constant"
	"seal/server/common/engine"
	"seal/server/common/models"
	"seal/server/common/utils"
	"seal/server/config"
	"time"

	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeromicro/go-zero/core/logx"
)

func GenSign(nestContract *models.Contract) models.Contract {
	/* 创建文件签名 */

	company := models.CompanyVcEntity{
		LegalDid:    config.COMPANY_DID,
		Name:        "测试信息技术研究院有限公司",
		LegalName:   "张三",
		CreditCode:  "9132000073225110X",
		CompanyType: 1,
	}

	credential := GenCompanyVC(company)
	credentialBytes, _ := hexutil.Decode(credential.Hash())
	_, _, err := tool.InitIdentity().AddIdentity(credentialBytes)
	if err != nil {
		fmt.Println("company credential add identity error")
		return models.Contract{}
	}

	s := map[string]string{
		"signType": "0",
        "position": "0",
        "sign": "https://www.zbiti.com",
        "name": "测试科技研究院",
        "phone": "12345678912",
        "signIndex": "0",
        "x": "204.575",
        "y": "198.73",
	}

	entity := models.Contract{
		Context:       utils.GenDefaultCredentialContext(),
		NestSignature: nestContract,
		Id:            utils.RandomUUID(),
		SignerVC:      credential,
		SignDate:      time.Now().String()[:19],
		SealType:      constant.COMPANY,
		ContractHash:  "0xce779b350231c5eb8b1bba07167f8760f1312c5ef2c2226dd4fd488a2eb7c7f9",
		SealsClaim:    s,
	}
	entity.Proof = utils.GenContractProof(entity, config.ADMIN_PRIVATE_KEY)

	return entity
}

func verifyFileOnChain(contract models.Contract) bool {
	/* 已上链文件校验 */

	// 验证文件格式
	fileFormat := utils.IsContractVaild(&contract)
	if !fileFormat {
		return false
	}
	// 验证签名
	sigVerify := utils.VerifyPrefixedSignatureFromWeId(utils.GetContractHash(contract),
		contract.Proof[constant.CREDENTIAL_SIGNATURE],
		contract.Proof[constant.PROOF_CREATOR])
	if !sigVerify {
		return false
	}
	// 验证签名是否上链 查询上链时间
	var sealAddr string
	if contract.SealType == 0 {
		sealAddr = "0x0000000000000000000000000000000000000000"
	} else {
		sealAddr = utils.GidToAddr(contract.SignerVC.Claim[constant.COMPANY_DID])
	}
	contractHashByte, _ := hexutil.Decode(contract.ContractHash)
	exist := engine.IsContractExisted(contractHashByte, utils.GetContractHash(contract), sealAddr,
		contract.SealType, utils.GidToAddr(contract.Proof[constant.PROOF_CREATOR]))
	if !exist {
		return false
	}

	// 查询签署时间
	queryTime := engine.QuerySignTimel(contractHashByte, utils.GetContractHash(contract), sealAddr,
		contract.SealType, utils.GidToAddr(contract.Proof[constant.PROOF_CREATOR]))

	// 验证链上issuerVC 在签署文件前是否有效
	verify := verifyByTime(contract.SignerVC, queryTime)

	return verify
}

func verifyBeforeChain(contract models.Contract) bool {
	/* 上链前校验 */

	// 验证文件格式
	fileFormat := utils.IsContractVaild(&contract)
	if !fileFormat {
		logx.Error("fileformat")
		return false
	}

	// 验证链上issuerVC是否有效
	credential := contract.SignerVC
	verify := verify(&credential)
	if !verify {
		logx.Error("verify")
		return false
	}

	// 验证文件签名
	sigVerify := utils.VerifyPrefixedSignatureFromWeId(utils.GetContractHash(contract),
		contract.Proof[constant.CREDENTIAL_SIGNATURE],
		contract.Proof[constant.PROOF_CREATOR])
	if !sigVerify {
		logx.Error("sigVerify")
		return false
	}

	// 如果使用企业签章
	if contract.SealType != constant.PERSON {
		sealAddr := utils.GidToAddr(contract.SignerVC.Claim[constant.COMPANY_DID])
		addr := (contract.Proof[constant.PROOF_CREATOR])
		vaild := engine.QuerySealStatus(sealAddr)
		if !vaild {
			fmt.Println("签章合约无效")
			return false
		}

		// 验证签名人是否有此签章授权
		vaild = engine.QuerySealApprovl((sealAddr), contract.SealType, (addr))
		if !vaild {
			fmt.Println("签名人无此印章授权")
			return false
		}
	}

	return true
}

func ToChain(contract models.Contract) (vaild bool, rec *types.Receipt, signHash, sealAddr, signAddr string) {
	/* 文件上链 */

	// 校验
	vaild = verifyBeforeChain(contract)
	if !vaild {
		fmt.Println("verify before chain failed")
		return false, nil, "", "", ""
	}

	// 判断是否已经上链
	if contract.SealType == 0 {
		sealAddr = "0x0000000000000000000000000000000000000000"
	} else {
		sealAddr = utils.GidToAddr(contract.SignerVC.Claim[constant.COMPANY_DID])
	}
	contractHashByte, _ := hexutil.Decode(contract.ContractHash)
	exist := engine.IsContractExisted(contractHashByte, utils.GetContractHash(contract), sealAddr,
		contract.SealType, utils.GidToAddr(contract.Proof[constant.PROOF_CREATOR]))
	if exist {
		return false, nil, "", "", ""
	}

	// 判断sealAddr
	staus := engine.QuerySealStatus(sealAddr)
	if !staus {
		fmt.Println("sealAddr is revoked")
		return false, nil, "", "", ""
	}

	// 判断签名用户
	rawData := utils.GetContractHash(contract)
	sig, _ := base64.StdEncoding.DecodeString(contract.Proof[constant.CREDENTIAL_SIGNATURE])
	pubKey, err := crypto.SigToPub(rawData, sig)
	if err != nil {
		return false, nil, "", "", ""
	}

	addr := (crypto.PubkeyToAddress(*pubKey).Hex())
	gid, _ := utils.AddrToGID(addr)
	if !engine.IsOwner_Contract(gid) {
		fmt.Println("sign error")
		return false, nil, "", "", ""
	}

	signAddr = utils.GidToAddr(contract.Proof[constant.PROOF_CREATOR])
	signHashByte := utils.GetContractHash(contract)
	signHash = hexutil.Encode(signHashByte)
	// 文件签名上链验证
	vaild, rec = engine.AddSign("QmVXdf2wdEnMn5K1azT7px4uzTxJrTKosoSZpsMTLvB1EG", contractHashByte, signHashByte, sealAddr, contract.SealType, signAddr)
	return
}
