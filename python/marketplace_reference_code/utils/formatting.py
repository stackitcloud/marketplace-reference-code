from stackit.stackitmarketplace.models.vendor_subscription import VendorSubscription


def format_subscription_details(subscription: VendorSubscription) -> str:
    """Format subscription details into a readable string."""
    return f"""
📦 Subscription Details:
══════════════════════════════
🔑 ID:              {subscription.subscription_id}
📊 Lifecycle State: {subscription.lifecycle_state}
📁 Project ID:      {subscription.project_id}
🛍️  Product:
\t📦 ID:               {subscription.product.product_id}
\t📝 Name:             {subscription.product.product_name}
\t🏢 Vendor:           {subscription.product.vendor_name}
\t🌐 Vendor Website:   {subscription.product.vendor_website_url}
\t🚚 Delivery Method:  {subscription.product.delivery_method}
\t📊 Lifecycle State:  {subscription.product.lifecycle_state}
\t💰 Price Type:       {subscription.product.price_type}
\t💳 Pricing Plan:     {subscription.product.pricing_plan}
══════════════════════════════
"""
