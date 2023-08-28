package utils

import (
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"seal/server/common/constant"
	"seal/server/common/models"
	"seal/server/common/xerr"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/sha3"
)

func GenCredentialProof(credential models.Credential, privateKey ecdsa.PrivateKey, disclosureMap map[string]string) (proof map[string]string) {
	proof = map[string]string{
		constant.PROOF_CREATED:        credential.IssuanceDate,
		constant.PROOF_CREATOR:        credential.Issuer,
		constant.PROOF_TYPE:           constant.ECDSA,
		constant.CREDENTIAL_SIGNATURE: genCredentialSignature(credential, privateKey, disclosureMap),
	}
	return
}

func genCredentialSignature(credential models.Credential, privateKey ecdsa.PrivateKey, disclosureMap map[string]string) (sigData string) {
	rawData := GenCredentialThumbprintWithoutSig(credential, disclosureMap)
	sigData = signMessage(rawData, privateKey)

	return
}

func signMessage(msg string, privateKey ecdsa.PrivateKey) (sigData string) {
	hash := crypto.Keccak256([]byte(msg))
	sig, err := crypto.Sign(hash, &privateKey)
	if err != nil {
		panic(err)
	}

	sigData += base64.StdEncoding.EncodeToString(sig)

	return
}

func GenCredentialThumbprintWithoutSig(credential models.Credential, disclosures map[string]string) string {
	rawCredential := credential
	rawCredential.Proof = nil

	return genCredentialThumbprint(rawCredential, disclosures)
}

func genCredentialThumbprint(credential models.Credential, disclosures map[string]string) string {
	credMap := make(map[string]string)
	copier.Copy(&credMap, &credential)
	claimHash := genClaimHash(credential, disclosures)
	credMap[constant.CLAIM] = claimHash
	jsonStr, _ := json.Marshal(credMap)

	return string(jsonStr)
}

func genClaimHash(credential models.Credential, disclosures map[string]string) (hash string) {
	claim := credential.Claim
	var disclosureMap map[string]string
	if disclosures == nil {
		copier.Copy(&disclosureMap, &claim)
		for k := range claim {
			disclosureMap[k] = ""
		}
	} else {
		copier.Copy(&disclosureMap, &disclosures)
	}

	claimHashMap := make(map[string]string)
	for k := range disclosureMap {
		hash := sha3.New256()
		hash.Write([]byte(claimHashMap[k]))
		claimHashMap[k] = string(hash.Sum(nil))
	}

	// sort
	list := make(map[string]string)
	keys := make([]string, 0, len(claimHashMap))
	for k := range claimHashMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		list[k] = claimHashMap[k]
	}
	for _, v := range list {
		hash += v
	}

	return
}

func GenDefaultCredentialContext() (context string) {
	context = constant.DEFAULT_CREDENTIAL_CONTEXT
	return
}

func GenSealsClaim(claim string) map[string]string {
	var jsonObj map[string]interface{}
	err := json.Unmarshal([]byte(claim), &jsonObj)
	if err != nil {
		fmt.Println("unmarshal company claim error")
		return nil
	}

	sealsClaim := make(map[string]string)
	for key, value := range jsonObj {
		sealsClaim[key] = fmt.Sprint(value)
	}

	return sealsClaim
}

func GenContractProof(contract models.Contract, privateKey string) (proof map[string]string) {
	creator, _ := AddrToGID(priKeyToPubAddr(privateKey))
	proof = map[string]string{
		constant.PROOF_CREATED:        contract.SignDate,
		constant.PROOF_CREATOR:        creator,
		constant.PROOF_TYPE:           constant.ECDSA,
		constant.CREDENTIAL_SIGNATURE: genContractSignature(contract, privateKey),
	}
	return
}

func genContractSignature(contract models.Contract, privateKey string) string {
	rawData := GetContractHash(contract)
	priKey := DeSerializePri(privateKey)
	sig, _ := crypto.Sign(rawData, priKey)

	return base64.StdEncoding.EncodeToString(sig)
}

func GetContractHash(contract models.Contract) []byte {
	rawData := getContractThumbprintWithoutSig(contract)
	b := crypto.Keccak256([]byte(rawData))
	return b
}

func getContractThumbprintWithoutSig(contract models.Contract) string {
	rawContract := new(models.Contract)
	copier.Copy(&rawContract, contract)
	rawContract.Proof = nil

	return genContractThumbprint(rawContract)
}

func genContractThumbprint(contract *models.Contract) string {
	credMap := new(models.Contract)
	copier.Copy(credMap, contract)
	data, _ := json.Marshal(credMap)

	return string(data)
}

func IsCredentialVaild(args *models.Credential) (err xerr.CodeError) {
	err = *xerr.NewErrCode(xerr.CREDENTIAL_ERROR)
	if args == nil || !isGidVaild(args.Issuer) {
		return err
	}
	if args.IssuanceDate == "" || args.ExpirationDate == "" {
		return err
	}
	if args.IssuanceDate != "" && !verifyTime(args.IssuanceDate, args.ExpirationDate) || args.Claim == nil {
		return err
	}

	return *xerr.NewErrCode(xerr.OK)
}

func IsContractVaild(contract *models.Contract) bool {
	if contract.IsContractEmpty() {
		fmt.Println("contract is null")
		return false
	}
	signerVC := contract.SignerVC

	if signerVC.IsCredentialEmpty() {
		fmt.Println("contract signerVC is null")
		return false
	}

	issuanceDate := contract.SignDate
	if issuanceDate == "" {
		fmt.Println("issuanceDate is null")
		return false
	}

	if contract.SealsClaim == nil {
		fmt.Println("contract seals claim is null")
		return false
	}

	if contract.Id == "" {
		fmt.Println("contract id is null")
		return false
	}

	if contract.Context == "" {
		fmt.Println("contract context is null")
		return false
	}

	proof := contract.Proof
	proofType := proof[constant.PROOF_TYPE]
	if proofType != constant.ECDSA && proofType != constant.ECDSA_R {
		fmt.Println("proof tpye is not in range")
		return false
	}

	created := proof[constant.PROOF_CREATED]
	if created == "" {
		fmt.Println("proof created is wrong")
		return false
	}

	creator := proof[constant.PROOF_CREATOR]
	if creator == "" || !isGidVaild(creator) {
		fmt.Println("creator format is wrong")
		return false
	}

	/* check signature */
	signature := proof[constant.CREDENTIAL_SIGNATURE]
	_, err := base64.StdEncoding.DecodeString(signature)
	if signature == "" || err != nil {
		fmt.Println("signature format is wrong")
		return false
	}

	return true
}

/* ===================================================================================== */

/**
 *@description: 使用以太坊钱包库创建以太坊账户凭证
 *@param:privateKey 私钥
 *@return: ecKeyPair,address
 */
func CreateCredentials(pri string) (*keystore.Key, error) {
	privateKey := DeSerializePri(pri)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, nil
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	credentials := keystore.Key{
		Id:         uuid.New(), //多的参数，暂时用不到 Version 4
		Address:    address,
		PrivateKey: &ecdsa.PrivateKey{D: privateKey.D, PublicKey: *publicKeyECDSA},
	}
	// log.Println(credentials)
	return &credentials, nil

}

/**
 *@description: 生成一个椭圆曲线加密算法（ECDSA）的密钥对，其中包括一个私钥和相应的公钥。
 *@param:
 *@return: *ecdsa.PrivateKey,error
 */
func CreateKeyPairToAddress() string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	// 从私钥导出公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to convert public key to ECDSA")
	}

	// 计算地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	gid, err := AddrToGID(address.String())
	if err != nil {
		panic(err)
	}
	return gid
}
