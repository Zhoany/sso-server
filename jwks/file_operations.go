package jwks

import (
	"crypto/rsa"
	
	"encoding/json"
	"log"
	
	
	
	"io/ioutil"
	
	"os"
	"crypto/x509"
	"encoding/pem"
	
	
)

func SavePrivateKeyToFile(privKey *rsa.PrivateKey, filename string) error {
	
		// Convert the RSA private key to PEM format
		privKeyBytes := x509.MarshalPKCS1PrivateKey(privKey)
		privKeyPEM := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privKeyBytes,
		}
	
		// Create and write the private key file
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
	
		return pem.Encode(file, privKeyPEM)
	
}

func WriteJWKSToFile(jwks *JWKS, filename string) {
	
		data, err := json.MarshalIndent(jwks, "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshal JWKS: %v", err)
		}
	
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}
		defer file.Close()
	
		_, err = file.Write(data)
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}
	
}

func ReadLocalJWKS() (interface{}, error) {
	
		data, err := ioutil.ReadFile("/home/codeserver/Code/go/oidcprovider/cert/jwks.json")
		if err != nil {
			return nil, err
		}
	
		var jwks interface{}
		if err := json.Unmarshal(data, &jwks); err != nil {
			return nil, err
		}
		return jwks, nil
	
}

func LoadRSAPrivateKeyFromFile(filename string) (*rsa.PrivateKey, error) {
	
		// 1. Read the private key from the file
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
	
		// 2. Decode the PEM block
		block, _ := pem.Decode(data)
		if block == nil {
			return nil, err
		}
	
		// 3. Parse the PEM block to get the RSA private key
		priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	
		return priv, nil
	}
	func WriteJWTToFile(jwt, filename string) error {
		return ioutil.WriteFile(filename, []byte(jwt), 0644)
	}
	func ReadJWTFromFile(filename string) (string, error) {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}