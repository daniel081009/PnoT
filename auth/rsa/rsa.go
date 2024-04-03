package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

func CreateRootRsa(password string) *rsa.PrivateKey {
	// Create the RSA key
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		fmt.Println(err)
	}
	return privateKey
}

func EncryptMessage(publicKey *rsa.PublicKey, message string) ([]byte, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(message), nil)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func DecryptMessage(privateKey *rsa.PrivateKey, ciphertext []byte) (string, error) {
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func SignMessage(privateKey *rsa.PrivateKey, message string) ([]byte, error) {
	hashed := sha256.Sum256([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func VerifySignature(publicKey *rsa.PublicKey, message string, signature []byte) error {
	hashed := sha256.Sum256([]byte(message))
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return err
	}
	return nil
}

func SaveToFile(privateKey *rsa.PrivateKey, filepath string) error {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privateKeyFile, err := os.Create("private.key")
	if err != nil {
		return err
	}
	err = pem.Encode(privateKeyFile, privateKeyBlock)
	if err != nil {
		return err
	}
	privateKeyFile.Close()
	return nil
}

func LoadFromFile(filepath string) (*rsa.PrivateKey, error) {
	privateKeyFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	privateKeyBytes, err := io.ReadAll(privateKeyFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	privateKeyFile.Close()
	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return privateKey, nil
}
