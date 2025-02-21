import base64
import logging
from typing import Optional

import jwt
import requests

from marketplace_reference_code.utils.exceptions import (
    TokenValidationError,
    TokenVerificationError,
)

logger = logging.getLogger(__name__)


class TokenValidator:
    PUBLIC_KEYS_URL = "https://keys.marketplace.stackit.cloud/v1/resolve-customer/keys.json"

    def get_public_keys(self) -> dict:
        """Fetch and decode the marketplace public keys."""
        logger.info("🔑 Starting public keys fetch...")
        response = requests.get(self.PUBLIC_KEYS_URL, timeout=10)
        response.raise_for_status()
        return response.json()

    def verify_token(self, token: str, public_keys: dict) -> None:
        """Verify the JWT token signature."""
        logger.info("🔐 Verifying token signature...")
        try:
            headers = jwt.get_unverified_header(token)
            jwt.decode(token, public_keys[headers["kid"]], algorithms=["RS256"], issuer=self.PUBLIC_KEYS_URL)
        except Exception as e:
            raise TokenVerificationError(str(e))


def validate_token(token: Optional[str]) -> None:
    """Main token validation function."""
    if not token:
        raise TokenValidationError("Token cannot be None or empty")

    try:
        validator = TokenValidator()
        public_key_set = validator.get_public_keys()
        logger.info("✅ Public key successfully fetched")

        validator.verify_token(token, public_key_set)
        logger.info("✅ Token signature verified successfully")
    except requests.RequestException as e:
        raise TokenValidationError(f"Failed to fetch public keys: {str(e)}") from e
    except TokenVerificationError as e:
        raise TokenValidationError(f"Failed to verify token signature: {str(e)}") from e
    except Exception as e:
        raise TokenValidationError(
            f"Unexpected error during token validation: {str(e)}"
        ) from e
