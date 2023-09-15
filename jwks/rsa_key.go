package jwks

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

func GenerateRSAKey() *rsa.PrivateKey {
	
		// 使用RSA密钥生成JSON Web Key
		key, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			log.Fatalf("Error generating RSA key: %v", err)
			return nil // 或者返回一个默认密钥或其他处理方式
		}
		return key
	
}