package rsautil

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func SignWithSHA256(privateKey *rsa.PrivateKey, message string) (signature string, err error) {
	hash := crypto.Hash.New(crypto.SHA256)
	_, err = hash.Write([]byte(message))
	if err != nil {
		return "", fmt.Errorf("hash: %v", err)
	}

	hashed := hash.Sum(nil)
	signed, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", fmt.Errorf("sign: %v", err)
	}
	return base64.StdEncoding.EncodeToString(signed), nil
}

func VerifySignatureWithSHA256(publicKey *rsa.PublicKey, signature string, message string) error {
	b, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("base64 decode: %v, signature: %s", err, signature)
	}

	hashed := sha256.Sum256([]byte(message))
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], b)
	if err != nil {
		return fmt.Errorf("verify signature: %v, signature: %s, message: %s", err, signature, message)
	}
	return nil
}
