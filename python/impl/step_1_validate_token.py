import base64
import logging
from typing import Any, Optional

import jwt
import requests

from .utils.exceptions import TokenValidationError, TokenVerificationError

logger = logging.getLogger(__name__)


class TokenValidator:
    PUBLIC_KEY_URL = "https://keys.marketplace.stackit.cloud/v2/resolve-customer.pub"

    def get_public_key(self) -> bytes:
        """Fetch and decode the marketplace public key."""
        logger.info("🔑 Starting public key fetch...")
        response = requests.get(self.PUBLIC_KEY_URL, timeout=10)
        response.raise_for_status()
        return base64.b64decode(response.content)

    def verify_token(self, token: str, public_key: bytes) -> dict[str, Any]:
        """Verify the JWT token signature."""
        logger.info("🔐 Verifying token signature...")
        try:
            return jwt.decode(token, public_key, algorithms=["RS256"])
        except Exception as e:
            raise TokenVerificationError(str(e))


def validate_token(token: Optional[str]) -> None:
    """Main token validation function."""
    if not token:
        raise TokenValidationError("Token cannot be None or empty")

    try:
        validator = TokenValidator()
        public_key = validator.get_public_key()
        logger.info("✅ Public key successfully fetched")

        validator.verify_token(token, public_key)
        logger.info("✅ Token signature verified successfully")
    except requests.RequestException as e:
        raise TokenValidationError(f"Failed to fetch public key: {str(e)}") from e
    except TokenVerificationError as e:
        raise TokenValidationError(f"Failed to verify token signature: {str(e)}") from e
    except Exception as e:
        raise TokenValidationError(
            f"Unexpected error during token validation: {str(e)}"
        ) from e
