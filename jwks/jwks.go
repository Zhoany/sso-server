package jwks

import (
	"gopkg.in/square/go-jose.v2"
	"log"

)

type JWKS struct {
	Keys []jose.JSONWebKey `json:"keys"`
}

var jwks *JWKS

func InitJWKS() {
	
		// 使用RSA密钥生成JSON Web Key
		rsaKey := GenerateRSAKey()
		if err :=SavePrivateKeyToFile(rsaKey, "private_key.pem"); err != nil {
			log.Fatalf("Failed to save private key to file: %v", err)
		}
		jwk := jose.JSONWebKey{
			Key:       rsaKey.Public(),
			Algorithm: "RS256",
			Use:       "sig",
			KeyID:     "sample-key-id",
		}
		
		// 创建JWKS并将JWK添加到Keys列表
		jwks = &JWKS{
			Keys: []jose.JSONWebKey{jwk},
		}
	WriteJWKSToFile(jwks,"./jwks.json")
	
}

