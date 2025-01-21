package internal

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	marketplacePublicKeyUrl = "https://ct-dev-stackit-marketplace-pubkeys.object.storage.eu01.onstackit.cloud/v6/resolve-customer.pub"
)

// GetMarketplacePublicKey fetches the public key from the marketplace
// and returns it as a *rsa.PublicKey
func GetMarketplacePublicKey() *rsa.PublicKey {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	log.Println("🔑 Starting public key fetch...")
	resp, err := client.Get(marketplacePublicKeyUrl)
	if err != nil {
		log.Fatalf("❌ Failed to fetch public key: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("❌ Failed to fetch public key: status code %d", resp.StatusCode)
	}

	publicKey, err := io.ReadAll(io.LimitReader(resp.Body, 1024*1024)) // 1MB limit
	if err != nil {
		log.Fatalf("❌ Failed to read public key: %v", err)
	}

	log.Println("🔄 Decoding public key...")

	// Base64 decode the public key
	// TODO: This should not be necessary, remove once this is updated in the key storage
	decodedPublicKey, err := base64.StdEncoding.DecodeString(string(publicKey))
	if err != nil {
		log.Fatalf("❌ Failed to decode public key: %v", err)
	}
	// Parse the public key
	parsedKey, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		log.Fatalf("❌ Failed to parse public key: %v", err)
	}

	return parsedKey
}

// VerifyToken verifies the token signature against the public key
func VerifyToken(tokenString string, marketplacePublicKey *rsa.PublicKey) error {
	log.Println("🔐 Verifying token signature...")
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return marketplacePublicKey, nil
	})
	return err
}

// ValidateToken performs all the necessary steps to validate a STACKIT Marketplace token
func ValidateToken(tokenString string) error {
	marketplacePublicKey := GetMarketplacePublicKey()
	log.Println("✅ Public key successfully fetched")

	if err := VerifyToken(tokenString, marketplacePublicKey); err != nil {
		return fmt.Errorf("❌ Token verification failed: %w", err)
	}
	log.Println("✅ Token signature verified successfully")
	return nil
}
