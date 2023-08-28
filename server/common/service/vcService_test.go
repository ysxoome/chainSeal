package service

import (
	"crypto/ecdsa"
	"fmt"
	"seal/blockchain/tool"
	"seal/server/common/constant"
	"seal/server/common/models"
	"seal/server/common/utils"
	"seal/server/config"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/jinzhu/copier"
)

func TestGenPersonVC(t *testing.T) {
	keyStore, _ := utils.CreateCredentials(config.COMPANY_PRIVATE_KEY)

	personVC := models.PersonVC{
		KeyStore: keyStore,
		Name:     "张三",
		IdNumber: "340123202112120120",
	}
	credential := GenPersonVC(personVC)
	fmt.Println(credential)

	// verify
	pri := utils.DeSerializePri(config.ADMIN_PRIVATE_KEY)
	rawData := utils.GenCredentialThumbprintWithoutSig(credential, nil)
	sig := utils.StringToSig(credential.Proof[constant.CREDENTIAL_SIGNATURE])
	hash := crypto.Keccak256([]byte(rawData))
	publicKey := crypto.FromECDSAPub(pri.Public().(*ecdsa.PublicKey))
	valid := crypto.VerifySignature(publicKey, hash, sig[:(len(sig))-1])
	fmt.Println(valid)

	disclosureMap := make(map[string]string)
	copier.Copy(&disclosureMap, credential.Claim)

	pri = utils.DeSerializePri(config.ADMIN_PRIVATE_KEY)
	rawData = utils.GenCredentialThumbprintWithoutSig(credential, disclosureMap)
	sig = utils.StringToSig(credential.Proof[constant.CREDENTIAL_SIGNATURE])
	hash = crypto.Keccak256([]byte(rawData))
	publicKey = crypto.FromECDSAPub(pri.Public().(*ecdsa.PublicKey))
	valid = crypto.VerifySignature(publicKey, hash, sig[:(len(sig))-1])
	fmt.Println(valid)
}

func TestPersonCredentialHash(t *testing.T) {
	keyStore, _ := utils.CreateCredentials(config.COMPANY_PRIVATE_KEY)

	personVC := models.PersonVC{
		KeyStore: keyStore,
		Name:     "张三",
		IdNumber: "340123202112120120",
	}

	credential := GenPersonVC(personVC)
	fmt.Println(credential.Hash())
	cid, _ := utils.CredentialCID(shell.NewShell("http://192.168.103.114:8081"), credential)
	fmt.Println(cid)
	// fmt.Println(credential.Hash())
}

func TestPersonVCVerify(t *testing.T) {
	keyStore, _ := utils.CreateCredentials(config.COMPANY_PRIVATE_KEY)

	personVC := models.PersonVC{
		KeyStore: keyStore,
		Name:     "张三",
		IdNumber: "340123202112120120",
	}

	credential := GenPersonVC(personVC)
	// fmt.Println(credential)
	identity := credential.Hash()
	fmt.Println(identity)
	identityBytes, _ := hexutil.Decode(identity)
	identitySession := tool.InitIdentity()

	_, rec, err := identitySession.AddIdentity(identityBytes)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(rec.Status)

	if verify(&credential) {
		fmt.Println("verify success")
	}
}
