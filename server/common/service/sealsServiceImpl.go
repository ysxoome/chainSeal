package service

import (
	"log"
	ContractEngine "seal/server/common/engine"
	"seal/server/common/models"
	"seal/server/common/utils"

	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

/**
 *@description: 创建签章地址
 *@param legalDid, creditCode
 *@return: sealAddress
 */
func CalculateSeal(legalDid string, creditCode string) string {
	return ContractEngine.CalculateSeal(legalDid, creditCode)
}

/**
 * @description: 查询签章是否授权,+本地验证签名，+上链
 * @param signatureAuth
 * @return: bool
 */
func ApproveSeal(signatureAuth models.SignatureAuth) (bool, *types.Receipt) {

	responseDataBool := ContractEngine.QuerySealApprovl(
		signatureAuth.SealAddress,
		signatureAuth.TypeAuth,
		signatureAuth.TargetDid)

	if responseDataBool {
		log.Println("已经授权")
		return true, nil
	} else {
		log.Println("未授权即将授权")
		encodePacked := ContractEngine.EncodePacked(
			signatureAuth.SealAddress,
			signatureAuth.TypeAuth,
			signatureAuth.TargetDid,
			signatureAuth.AuthTime)

		hashByte, _ := hexutil.Decode(crypto.Keccak256Hash(encodePacked).String())

		sigVerify := utils.VerifyPrefixedSignatureFromWeId(
			hashByte,
			signatureAuth.AuthSign,
			signatureAuth.AuthDid)
		if !sigVerify {
			log.Println("签名验证失败,非法签名")
			return false, nil
		}

		//验证签名地址是否是owner()
		isSealOwner := ContractEngine.IsSealOwner(signatureAuth.SealAddress, signatureAuth.AuthDid)
		if !isSealOwner {
			log.Println("签名地址没有权限授权")
			return false, nil
		}

		//上链
		isSuccess, rec := ContractEngine.ApprovalDelegate(
			signatureAuth.SealAddress,
			signatureAuth.TypeAuth,
			signatureAuth.TargetDid,
			signatureAuth.AuthTime,
			signatureAuth.AuthSign)
		if !isSuccess {
			log.Println("授权合约失败")
			return false, nil
		} else {
			log.Println("ApprovalDelegate success")
			return true, rec
		}
	}
}

/**
 *
 */
func RevokeSeal(signatureAuth models.SignatureAuth) (bool, *types.Receipt) {
	responseDataBool := ContractEngine.QuerySealApprovl(
		signatureAuth.SealAddress,
		signatureAuth.TypeAuth,
		signatureAuth.TargetDid)
	if !responseDataBool {
		log.Println("已经取消授权")
		return true, nil
	} else {
		//本地验证签名
		encodePacked := ContractEngine.EncodePacked(
			signatureAuth.SealAddress,
			signatureAuth.TypeAuth,
			signatureAuth.TargetDid,
			signatureAuth.AuthTime)

		hashByte, _ := hexutil.Decode(crypto.Keccak256Hash(encodePacked).String())

		sigVerify := utils.VerifyPrefixedSignatureFromWeId(
			hashByte,
			signatureAuth.AuthSign,
			signatureAuth.AuthDid)
		if !sigVerify {
			log.Println("签名验证失败,非法签名")
			return false, nil
		}

		//验证签名地址是否是owner()
		isSealOwner := ContractEngine.IsSealOwner(signatureAuth.SealAddress, signatureAuth.AuthDid)
		if !isSealOwner {
			log.Println("签名地址没有权限授权")
			return false, nil
		}

		//上链
		isSuccess, rec := ContractEngine.RevokeDelegate(
			signatureAuth.SealAddress,
			signatureAuth.TypeAuth,
			signatureAuth.TargetDid,
			signatureAuth.AuthTime,
			signatureAuth.AuthSign)
		if !isSuccess {
			log.Println("取消授权失败")
			return false, nil
		} else {
			return true, rec
		}
	}
}
