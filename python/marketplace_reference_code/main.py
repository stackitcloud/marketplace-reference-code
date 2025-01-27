import logging
import os

from stackit.core.configuration import Configuration  # type: ignore
from stackit.stackitmarketplace.api.default_api import DefaultApi

from marketplace_reference_code.step_1_validate_token import validate_token
from marketplace_reference_code.step_2_resolve_customer import resolve_customer
from marketplace_reference_code.step_3_approve_subscription import approve_subscription
from marketplace_reference_code.utils.exceptions import (
    CustomerResolutionError,
    MarketplaceError,
    SubscriptionApprovalError,
    TokenValidationError,
)
from marketplace_reference_code.utils.formatting import format_subscription_details

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


class MarketplaceClient:
    def __init__(self) -> None:
        self.token = os.getenv("MP_REF_CODE_REDIRECT_TOKEN")
        self.vendor_project_id = os.getenv("MP_REF_CODE_VENDOR_PROJECT_ID")

        if not self.token or not self.vendor_project_id:
            raise ValueError("Required environment variables are not set")

        self.config = Configuration(
            custom_endpoint="https://stackit-marketplace-dev.api.stg.stackit.cloud"
        )
        self.client = DefaultApi(self.config)

    def process_subscription(self) -> None:
        """Main workflow for processing a marketplace subscription."""
        try:
            # Step 1: Validate token
            logger.info("🔐 Authenticating token...")
            validate_token(self.token)
            logger.info("✅ Token authenticated successfully")

            # Step 2: Resolve customer
            logger.info("🔐 Resolving customer...")
            subscription = resolve_customer(
                self.client, self.vendor_project_id, self.token
            )
            logger.info("✅ Customer resolved successfully")
            logger.info(format_subscription_details(subscription))

            # Step 3: Approve subscription
            logger.info("🔐 Approving subscription...")
            approve_subscription(
                self.client, self.vendor_project_id, subscription.subscription_id
            )
            logger.info("✅ Subscription approved successfully")

        except TokenValidationError as e:
            logger.error(f"❌ Token validation failed: {str(e)}")
            raise
        except CustomerResolutionError as e:
            logger.error(f"❌ Customer resolution failed: {str(e)}")
            raise
        except SubscriptionApprovalError as e:
            logger.error(f"❌ Subscription approval failed: {str(e)}")
            raise
        except Exception as e:
            logger.error(f"❌ Unexpected error: {str(e)}")
            raise MarketplaceError("Failed to process subscription") from e


def main() -> None:
    client = MarketplaceClient()
    client.process_subscription()


if __name__ == "__main__":
    main()
