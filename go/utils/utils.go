package utils

import (
	"fmt"
	"strings"

	"github.com/stackitcloud/stackit-sdk-go/services/stackitmarketplace/v1api"
)

func VendorSubscriptionToString(subscription *v1api.MarketplaceSubscriptionV1Subscription) string {
	builder := strings.Builder{}

	builder.WriteString("\n📦 Vendor Subscription Details:\n")
	builder.WriteString("══════════════════════════════\n")

	// Main subscription details
	fmt.Fprintf(&builder, "🔑 ID:              %s\n", subscription.SubscriptionId)
	fmt.Fprintf(&builder, "📊 Lifecycle State: %s\n", subscription.LifecycleState)

	// Project details
	if subscription.ProjectId != "" {
		fmt.Fprintf(&builder, "📁 Project ID:      %s\n", subscription.ProjectId)
	}

	// Subscription Product details
	if subscription.Product.ProductId != "" {
		builder.WriteString("🛍️  Product:\n")

		// Basic product information
		fmt.Fprintf(&builder, "\t📦 ID:               %s\n", subscription.Product.ProductId)
		fmt.Fprintf(&builder, "\t📝 Name:             %s\n", subscription.Product.ProductName)

		// Vendor information
		fmt.Fprintf(&builder, "\t🏢 Vendor:           %s\n", subscription.Product.VendorName)
		fmt.Fprintf(&builder, "\t🌐 Vendor Website:   %s\n", subscription.Product.VendorWebsiteUrl)

		// Product status and delivery
		fmt.Fprintf(&builder, "\t🚚 Delivery Method:  %s\n", subscription.Product.DeliveryMethod)
		fmt.Fprintf(&builder, "\t📊 Lifecycle State:  %s\n", subscription.Product.LifecycleState)

		// Pricing information
		fmt.Fprintf(&builder, "\t💰 Price Type:       %s\n", subscription.Product.PriceType)
		fmt.Fprintf(&builder, "\t💳 Pricing Plan:     %s\n", subscription.Product.PricingPlan)
	}

	builder.WriteString("\n══════════════════════════════\n")
	return builder.String()
}
