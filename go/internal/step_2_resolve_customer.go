package internal

import (
	"context"

	"github.com/stackitcloud/stackit-sdk-go/services/stackitmarketplace"
)

func ResolveCustomer(ctx context.Context, client *stackitmarketplace.APIClient, vendorProjectID string, token string) (*stackitmarketplace.VendorSubscription, error) {
	// build the resolve customer request
	req := client.ResolveCustomer(ctx, vendorProjectID).ResolveCustomerPayload(stackitmarketplace.ResolveCustomerPayload{
		Token: &token,
	})

	// execute the request
	subscription, err := req.Execute()
	if err != nil {
		return nil, err
	}

	return subscription, nil
}
