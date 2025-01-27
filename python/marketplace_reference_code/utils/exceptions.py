class MarketplaceError(Exception):
    """Base exception for all marketplace operations."""

    pass


class TokenValidationError(MarketplaceError):
    """Raised when token validation fails."""

    pass


class TokenVerificationError(MarketplaceError):
    """Raised when token verification fails."""

    pass


class CustomerResolutionError(MarketplaceError):
    """Raised when customer resolution fails."""

    pass


class SubscriptionApprovalError(MarketplaceError):
    """Raised when subscription approval operation fails."""

    pass
