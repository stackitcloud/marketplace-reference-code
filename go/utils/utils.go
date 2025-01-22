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
		builder.WriteString(SubscriptionProductToString(subscription.Product))
	}

	builder.WriteString("\n══════════════════════════════\n")
	return builder.String()
}

func SubscriptionProductToString(product *stackitmarketplace.SubscriptionProduct) string {
	builder := strings.Builder{}

	builder.WriteString("🛍️  Product:\n")

	// Basic product information
	builder.WriteString(fmt.Sprintf("\t📦 ID:               %s\n", *product.ProductId))
	builder.WriteString(fmt.Sprintf("\t📝 Name:             %s\n", *product.ProductName))

	// Vendor information
	builder.WriteString(fmt.Sprintf("\t🏢 Vendor:           %s\n", *product.VendorName))
	builder.WriteString(fmt.Sprintf("\t🌐 Vendor Website:   %s\n", *product.VendorWebsiteUrl))

	// Product status and delivery
	builder.WriteString(fmt.Sprintf("\t🚚 Delivery Method:  %s\n", *product.DeliveryMethod))
	builder.WriteString(fmt.Sprintf("\t📊 Lifecycle State:  %s\n", *product.LifecycleState))

	// Pricing information
	builder.WriteString(fmt.Sprintf("\t💰 Price Type:       %s\n", *product.PriceType))
	builder.WriteString(fmt.Sprintf("\t💳 Pricing Plan:     %s\n", *product.PricingPlan))

	return builder.String()
}
