package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

// hashed pwd
func HashPassword(message string) string {
	messageToBye := []byte(message)
	sha512Hasher := sha512.New()
	sha512Hasher.Write(messageToBye)

	return hex.EncodeToString(sha512Hasher.Sum(nil))
}

// cek if two hashed match
func DoPasswordsMatch(hashPwd, currPwd string) bool {
	currPwdHash := HashPassword(currPwd)
	return hashPwd == currPwdHash
}

func main() {
	messageStr := "test-enckripsi-symetric"
	signatureStr := HashPassword(messageStr)

	fmt.Println("Signature: %s", signatureStr)

	fmt.Println(DoPasswordsMatch(signatureStr, messageStr))
}
