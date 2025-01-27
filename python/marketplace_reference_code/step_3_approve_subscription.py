from stackit.stackitmarketplace.api.default_api import DefaultApi


def approve_subscription(
    client: DefaultApi, vendor_project_id: str | None, subscription_id: str | None
) -> None:
    """Approve a marketplace subscription."""
    if not vendor_project_id or not subscription_id:
        raise ValueError(
            "Vendor project ID and subscription ID cannot be None or empty"
        )
    client.approve_subscription(vendor_project_id, subscription_id)
