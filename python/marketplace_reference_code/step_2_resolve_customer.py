from stackit.stackitmarketplace.api.default_api import DefaultApi
from stackit.stackitmarketplace.models.resolve_customer_payload import (
    ResolveCustomerPayload,
)
from stackit.stackitmarketplace.models.vendor_subscription import VendorSubscription

from marketplace_reference_code.utils.exceptions import CustomerResolutionError


def resolve_customer(
    client: DefaultApi,
    vendor_project_id: str | None,
    token: str | None,
) -> VendorSubscription:
    """Main customer resolution function."""
    if not vendor_project_id or not token:
        raise ValueError("Vendor project ID and token are required")

    try:
        payload = ResolveCustomerPayload(token=token)
        subscription = client.resolve_customer(vendor_project_id, payload)
        return subscription

    except Exception as e:
        raise CustomerResolutionError(f"Failed to resolve customer: {str(e)}") from e
