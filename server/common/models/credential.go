package models

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/ethereum/go-ethereum/crypto"
)

type Credential struct {
	Context        string            `json:"context"`
	Id             string            `json:"id"`
	Issuer         string            `json:"issuer"`
	IssuanceDate   string            `json:"issuanceDate"`
	ExpirationDate string            `json:"expirationDate"`
	Claim          map[string]string `json:"claim"`
	Proof          map[string]string `json:"proof"`
}

func (c *Credential) Hash() string {
	b, _ := json.Marshal(c)
	return crypto.Keccak256Hash(b).String()
}

func (c *Credential)IsCredentialEmpty() bool {
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
		default:
			// 检查其他类型的字段是否为零值
			if !field.IsZero() {
				return false
			}
		}
	}

	return true
}

func (c Credential) WriteCredentialToFile(filename string) error {
	output, _ := json.MarshalIndent(c, "", "    ")
	err := os.WriteFile(filename, output, 0644)
	if err != nil {
		return fmt.Errorf("无法写入文件：%v", err)
	}

	return nil
}
