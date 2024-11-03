package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// load private key from file
	privateKeyFile, err := ioutil.ReadFile("private_key_2.pem")
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBlock, _ := pem.Decode(privateKeyFile)
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(privateKey)

	// sign a messge using private key
	message := []byte("hello world")
	hash := sha256.Sum256(message)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(signature)

	// load public key from file
	publicKeyFile, err := ioutil.ReadFile("public_key_2.pem")
	if err != nil {
		log.Fatal(err)
	}
	publicKeyBlock, _ := pem.Decode(publicKeyFile)
	publicKeyInterface, err := x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	// Verify the signature using the public key
	hash = sha256.Sum256(message)
	err = rsa.VerifyPKCS1v15(publicKeyInterface, crypto.SHA256, hash[:], signature)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Signature is valid")
	}

}
