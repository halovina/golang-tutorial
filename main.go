package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func main() {
	// Generate RSA private key menggunakan 2048 bits
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(privateKey)
	// Encode  private key ke PEM format
	privateKeyPem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	privateKeyFile, err := os.Create("private_key_2.pem")
	if err != nil {
		log.Fatal(err)
	}

	pem.Encode(privateKeyFile, privateKeyPem)
	privateKeyFile.Close()

	// Extract  public key dari the private key

	publicKey := &privateKey.PublicKey

	// Encode  public key ke PEM format
	publicKeyPem := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	publicKeyFile, err := os.Create("public_key_2.pem")
	if err != nil {
		log.Fatal(err)
	}
	pem.Encode(publicKeyFile, publicKeyPem)
	publicKeyFile.Close()

}
