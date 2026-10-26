package internal

import (
	"context"

	"github.com/stackitcloud/stackit-sdk-go/services/stackitmarketplace/v1api"
)

func ApproveSubscription(ctx context.Context, client *v1api.APIClient, vendorProjectID string, subscriptionID string) error {
	// build the approve-subscription request
	req := client.DefaultAPI.ApproveSubscription(ctx, vendorProjectID, subscriptionID).ApproveSubscriptionPayload(
		v1api.ApproveSubscriptionPayload{
			InstanceTarget:       new("https://my-app.eu"), // optional
			AdditionalProperties: nil,
		},
	)
	// execute the request
	return req.Execute()
}
