package main

import (
	"context"
	"log"
	"os"

	"github.com/stackitcloud/marketplace-reference-code/internal"
	"github.com/stackitcloud/marketplace-reference-code/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/stackitmarketplace"
)

const (
	VendorProjectIdEnvVar  = "MP_REF_CODE_VENDOR_PROJECT_ID"
	MarketplaceTokenEnvVar = "MP_REF_CODE_REDIRECT_TOKEN"
)

func main() {
	ctx := context.Background()

	log.Println("🔐 Setting up SDK client...")
	client, err := stackitmarketplace.NewAPIClient()
	if err != nil {
		log.Fatalf("❌ Failed to setup SDK client: %v", err)
	}
	log.Println("✅ SDK client setup successfully")

	// read the token and the vendor project ID from the environment variables
	tokenString, ok := os.LookupEnv(MarketplaceTokenEnvVar)
	if !ok {
		log.Fatalf("❌ The required environment variable %s is not set", MarketplaceTokenEnvVar)
	}
	vendorProjectID, ok := os.LookupEnv(VendorProjectIdEnvVar)
	if !ok {
		log.Fatalf("❌ The required environment variable %s is not set", VendorProjectIdEnvVar)
	}

	// x-stackit-marketplace-token authentication
	log.Println("🔐 Authenticating token...")
	err = internal.ValidateToken(tokenString)
	if err != nil {
		log.Fatalf("❌ Token authentication failed: %v", err)
	}
	log.Println("✅ Token authenticated successfully")

	// resolve customer
	log.Println("🔐 Resolving customer...")
	subscription, err := internal.ResolveCustomer(ctx, client, vendorProjectID, tokenString)
	if err != nil {
		log.Fatalf("❌ Failed to resolve customer: %v", err)
	}
	log.Printf("✅ Customer resolved successfully \n%+v", utils.VendorSubscriptionToString(subscription))

	// approve subscription
	log.Println("🔐 Approving subscription...")
	err = internal.ApproveSubscription(ctx, client, vendorProjectID, *subscription.SubscriptionId)
	if err != nil {
		log.Fatalf("❌ Failed to approve subscription: %v", err)
	}
	log.Println("✅ Subscription approved successfully")
}
