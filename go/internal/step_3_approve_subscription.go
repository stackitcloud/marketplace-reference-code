package internal

import (
	"context"

	"github.com/stackitcloud/stackit-sdk-go/services/stackitmarketplace"
)

func ApproveSubscription(ctx context.Context, client *stackitmarketplace.APIClient, vendorProjectID string, subscriptionID string) error {
	return client.ApproveSubscriptionExecute(ctx, vendorProjectID, subscriptionID)
}
