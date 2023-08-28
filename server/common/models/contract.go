package models

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/ethereum/go-ethereum/crypto"
)

type Contract struct {
	Context       string            `json:"context"`
	Id            string            `json:"id"`
	ContractHash  string            `json:"contractHash"`
	NestSignature *Contract         `json:"nestSignature"`
	SealType      uint8             `json:"sealType"`
	SignerVC      Credential        `json:"signerVC"`
	SignDate      string            `json:"signDate"`
	SealsClaim    map[string]string `json:"sealsClaim"`
	Proof         map[string]string `json:"proof"`
}

func (c *Contract) Hash() string {
	b, _ := json.Marshal(c)
	return crypto.Keccak256Hash(b).String()
}

func (c *Contract) IsContractEmpty() bool {
	if c == nil {
		return true
	}

	v := reflect.ValueOf(c).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.String, reflect.Map:
			if field.Len() > 0 {
				return false
			}
		case reflect.Ptr, reflect.Slice:
			if !field.IsNil() {
				return false
			}
		default:
			if !field.IsZero() {
				return false
			}
		}
	}

	return true
}

func (c Contract) WriteContractToFile(filename string) error {
	output, _ := json.MarshalIndent(c, "", "    ")
	err := os.WriteFile(filename, output, 0644)
	if err != nil {
		return fmt.Errorf("无法写入文件：%v", err)
	}

	return nil
}
