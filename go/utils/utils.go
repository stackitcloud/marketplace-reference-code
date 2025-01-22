package utils

import (
	"fmt"
	"strings"

	"github.com/stackitcloud/stackit-sdk-go/services/stackitmarketplace"
)

func VendorSubscriptionToString(subscription *stackitmarketplace.VendorSubscription) string {
	builder := strings.Builder{}

	builder.WriteString("\n📦 Vendor Subscription Details:\n")
	builder.WriteString("══════════════════════════════\n")

	// Main subscription details
	builder.WriteString(fmt.Sprintf("🔑 ID:              %s\n", *subscription.SubscriptionId))
	builder.WriteString(fmt.Sprintf("📊 Lifecycle State: %s\n", *subscription.LifecycleState))

	// Project details
	if subscription.ProjectId != nil {
		builder.WriteString(fmt.Sprintf("📁 Project ID:      %s\n", *subscription.ProjectId))
	}

	// Subscription Product details
	if subscription.Product != nil {
		builder.WriteString("🛍️  Product:\n")

		// Basic product information
		builder.WriteString(fmt.Sprintf("\t📦 ID:               %s\n", *subscription.Product.ProductId))
		builder.WriteString(fmt.Sprintf("\t📝 Name:             %s\n", *subscription.Product.ProductName))

		// Vendor information
		builder.WriteString(fmt.Sprintf("\t🏢 Vendor:           %s\n", *subscription.Product.VendorName))
		builder.WriteString(fmt.Sprintf("\t🌐 Vendor Website:   %s\n", *subscription.Product.VendorWebsiteUrl))

		// Product status and delivery
		builder.WriteString(fmt.Sprintf("\t🚚 Delivery Method:  %s\n", *subscription.Product.DeliveryMethod))
		builder.WriteString(fmt.Sprintf("\t📊 Lifecycle State:  %s\n", *subscription.Product.LifecycleState))

		// Pricing information
		builder.WriteString(fmt.Sprintf("\t💰 Price Type:       %s\n", *subscription.Product.PriceType))
		builder.WriteString(fmt.Sprintf("\t💳 Pricing Plan:     %s\n", *subscription.Product.PricingPlan))
	}

	builder.WriteString("\n══════════════════════════════\n")
	return builder.String()
}
