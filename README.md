<div align="center">
<br>
<img src=".github/images/stackit-logo.svg" alt="STACKIT logo" width="50%"/>
<br>
<br>
</div>

# STACKIT Marketplace Reference Code

This repository contains reference implementations for **vendors** looking to integrate with the **STACKIT Marketplace**, demonstrating the subscription approval flow and token validation process.

## What is the STACKIT Marketplace?

The [STACKIT Marketplace](https://docs.stackit.cloud/stackit/en/stackit-marketplace-322109923.html) is a digital platform connecting STACKIT customers with third-party digital products, serving as the technological foundation for rapidly expanding the STACKIT portfolio and building industry-specific cloud offerings. It provides access to a diverse range of products, from IaaS and SaaS to licenses, professional services, and datasets.

## Supported Languages

The integration examples are available in multiple languages:

- [Go](/go) - Go implementation
- [Python](/python) - Python implementation

Each language implementation demonstrates the same core functionality while following language-specific best practices.

## Getting Started

1. Choose your preferred language implementation from the folders above
2. Follow the language-specific README for setup instructions
3. Review the implementation details in the code
4. Adapt the examples to your specific use case

### Integration Flow

1. **Marketplace token validation**: Validates the marketplace token by:
   1. Fetching the public key from STACKIT
   2. Validating the token format
   3. Verifying the token signature
2. **Customer Resolution**: Uses the validated token to resolve customer information
3. **Subscription Approval**: Approves the marketplace subscription

### Key concepts

- **Redirect Token**: The Redirect Token, or simply token, is the JWT Token sent to the vendors when the STACKIT Marketplace redirects consumers to the vendor's `Sign Up/Login` page, which happens in the process of creating a subscription (see [SaaS Integration](https://docs.stackit.cloud/stackit/en/4-2-saas-frontend-integration-339673765.html) for more details).

## Additional Resources

- [Public Marketplace](https://marketplace.stackit.cloud/)
- [STACKIT Marketplace Documentation](https://docs.stackit.cloud/stackit/en/stackit-marketplace-322109923.html)
- [API Reference](https://docs.api.stackit.cloud/documentation/stackit-marketplace/version/v1)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
