package rsautil

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func NewPrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, fmt.Errorf("decode private key failed")
	}
	if block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("pem type is not private key")
	}

	pk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse private key: %v", err)
	}
	return pk, nil
}

func NewPrivateKeyWithFile(file string) (*rsa.PrivateKey, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("read private key file: %v", err)
	}
	return NewPrivateKey(string(b))
}

func NewPublicKey(publicKey string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, fmt.Errorf("decode public key failed")
	}
	if block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("pem type is not public key")
	}

	pk, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse public key: %v", err)
	}
	return pk, nil
}

func NewPublicKeyWithFile(file string) (*rsa.PublicKey, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("read public key file: %v", err)
	}
	return NewPublicKey(string(b))
}
