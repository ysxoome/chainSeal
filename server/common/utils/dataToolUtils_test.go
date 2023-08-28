package utils

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"seal/server/config"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestFromConfig(t *testing.T) {
	privateKey := DeSerializePri(config.COMPANY_PRIVATE_KEY)
	fmt.Println(privateKey)

	addr := crypto.PubkeyToAddress(*privateKey.Public().(*ecdsa.PublicKey))
	fmt.Println(addr)
}

func TestStringToSig(t *testing.T) {
	data := "XAOy+RSwmIH01vY677z09GJOT7yHoItV/FyrsYi/HtlWg4mnTGOkqGqCzeCSEXAHkArfZLDTzaH9jYcFuq8qVQA="
	fmt.Println(StringToSig(data))
}

func TestBigInt2Time(t *testing.T) {
	ms := big.NewInt(1692925523050).Int64()
	//  1693182726508
	sec := ms / 1000
	nsec := (ms % 1000) * 1000000
	fmt.Println(time.Unix(sec, nsec).String()[:19])

	fmt.Println(time.Now().UnixMilli())
}
