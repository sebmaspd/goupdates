package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

// This example signs and verifies an update file using RSA and SHA-256.
func main() {
	// Simulated firmware data
	updateData := []byte("Firmware v2.0 binary content")

	// Generate RSA keys (for demo purposes)
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey

	// --- Signing (by vendor) ---
	hash := sha256.Sum256(updateData)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, 0, hash[:])
	if err != nil {
		panic("Signing failed: " + err.Error())
	}
	fmt.Println("Signature created.")

	// --- Verification (by device) ---
	// Tampering test: uncomment below to simulate malicious tampering
	// updateData[0] = 'X'

	hashCheck := sha256.Sum256(updateData)
	err = rsa.VerifyPKCS1v15(publicKey, 0, hashCheck[:], signature)
	if err != nil {
		fmt.Println("Firmware verification failed! Update aborted.")
	} else {
		fmt.Println("Firmware verified. Update accepted.")
	}
}
