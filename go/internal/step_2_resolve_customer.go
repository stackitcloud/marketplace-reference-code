package internal

import (
	"context"

	"github.com/stackitcloud/stackit-sdk-go/services/stackitmarketplace/v1api"
)

func ResolveCustomer(ctx context.Context, client *v1api.APIClient, vendorProjectID string, token string) (*v1api.MarketplaceSubscriptionV1Subscription, error) {
	// build the resolve customer request
	req := client.DefaultAPI.ResolveCustomer(ctx, vendorProjectID).ResolveCustomerPayload(v1api.ResolveCustomerPayload{
		Token: token,
	})
	// execute the request
	return req.Execute()
}
