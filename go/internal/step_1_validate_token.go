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
	//TODO: update this with the final public DNS address
	marketplacePublicKeyUrl = "https://ct-dev-stackit-marketplace-pubkeys.object.storage.eu01.onstackit.cloud/v6/resolve-customer.pub"
)

// GetMarketplacePublicKey fetches the public key from the marketplace
// and returns it as a *rsa.PublicKey
func GetMarketplacePublicKey() (*rsa.PublicKey, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	log.Println("🔑 Starting public key fetch...")
	resp, err := client.Get(marketplacePublicKeyUrl)
	if err != nil {
		return nil, fmt.Errorf("fetching public key: %w", err)
	}

	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			log.Printf("❌ Failed to close the public key fetch response body: %v", cerr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetching public key: status code %d", resp.StatusCode)
	}

	publicKey, err := io.ReadAll(io.LimitReader(resp.Body, 1024*1024)) // 1MB limit
	if err != nil {
		return nil, fmt.Errorf("reading public key: %w", err)
	}

	log.Println("🔄 Decoding public key...")

	// Base64 decode the public key
	// TODO: This should not be necessary, remove once this is updated in the key storage
	decodedPublicKey, err := base64.StdEncoding.DecodeString(string(publicKey))
	if err != nil {
		return nil, fmt.Errorf("decoding public key: %w", err)
	}
	// Parse the public key
	parsedKey, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, fmt.Errorf("parsing public key: %w", err)
	}

	return parsedKey, nil
}

// VerifyToken verifies the token signature against the public key
func VerifyToken(token string, marketplacePublicKey *rsa.PublicKey) error {
	log.Println("🔐 Verifying token signature...")
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return marketplacePublicKey, nil
	})
	return err
}

// ValidateToken performs all the necessary steps to validate a STACKIT Marketplace token
func ValidateToken(token string) error {
	marketplacePublicKey, err := GetMarketplacePublicKey()
	if err != nil {
		return fmt.Errorf("getting the marketplace public key: %w", err)
	}
	log.Println("✅ Public key successfully fetched")

	if err := VerifyToken(token, marketplacePublicKey); err != nil {
		return fmt.Errorf("verifying the marketplace token: %w", err)
	}
	log.Println("✅ Token signature verified successfully")
	return nil
}
