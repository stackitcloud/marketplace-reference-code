package internal

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	marketplacePublicKeySetUrl = "https://keys.marketplace.stackit.cloud/v1/resolve-customer/keys.json"
)

// GetMarketplacePublicKey fetches the public key from the marketplace
// and returns it as a *rsa.PublicKey
func GetMarketplacePublicKey() (map[string]*rsa.PublicKey, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	log.Println("🔑 Starting public key fetch...")
	resp, err := client.Get(marketplacePublicKeySetUrl)
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

	publicKeyMapString, err := io.ReadAll(io.LimitReader(resp.Body, 1024*1024)) // 1MB limit
	if err != nil {
		return nil, fmt.Errorf("reading public key map: %w", err)
	}

	log.Println("🔄 Decoding public keys...")

	var unparsedKeyMap map[string]string
	err = json.Unmarshal(publicKeyMapString, &unparsedKeyMap)
	if err != nil {
		return nil, fmt.Errorf("parsing public key map: %w", err)
	}

	var parsedKeyMap = make(map[string]*rsa.PublicKey, len(unparsedKeyMap))
	for k, v := range unparsedKeyMap {
		parsedKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(v))
		if err != nil {
			return nil, fmt.Errorf("parsing public key: %w", err)
		}
		parsedKeyMap[k] = parsedKey
	}

	return parsedKeyMap, nil
}

// VerifyToken verifies the token signature against the public key
func VerifyToken(token string, marketplacePublicKeySet map[string]*rsa.PublicKey) error {
	log.Println("🔐 Verifying token signature...")
	var kid string
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, fmt.Errorf("failed to parse claims")
		}
		tokenIss, ok := claims["iss"].(string)
		if !ok || !strings.EqualFold(tokenIss, marketplacePublicKeySetUrl) {
			return nil, fmt.Errorf("unexpected issuer '%s', expected '%s' (claims[iss] is string: %t)", claims["iss"], marketplacePublicKeySetUrl, ok)
		}
		kid, ok = token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("failed to extract kid")
		}
		pubKey, ok := marketplacePublicKeySet[kid]
		if !ok {
			return nil, fmt.Errorf("no public key for kid: %s", kid)
		}
		return pubKey, nil
	})
	if err != nil || !parsedToken.Valid {
		return fmt.Errorf("verifying token: %w", err)
	}
	return nil
}

// ValidateToken performs all the necessary steps to validate a STACKIT Marketplace token
func ValidateToken(token string) error {
	marketplacePublicKeySet, err := GetMarketplacePublicKey()
	if err != nil {
		return fmt.Errorf("getting the marketplace public key: %w", err)
	}
	log.Println("✅ Public key successfully fetched")

	if err := VerifyToken(token, marketplacePublicKeySet); err != nil {
		return fmt.Errorf("verifying the marketplace token: %w", err)
	}
	log.Println("✅ Token signature verified successfully")
	return nil
}
