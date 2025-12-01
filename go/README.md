# STACKIT Marketplace Reference Code in Golang

This Golang reference code exemplifies how to integrate with the STACKIT Marketplace API as a vendor, showing the essential flow of authenticating and handling marketplace subscriptions.

## Overview

This example application demonstrates three key marketplace integration steps:

1. **Token Validation**: Validates marketplace tokens using public key cryptography
2. **Customer Resolution**: Retrieves customer information from subscription tokens
3. **Subscription Approval**: Processes and approves marketplace subscriptions

## Prerequisites

- Go 1.25 or higher
- A STACKIT Service Account Key
- A marketplace token (`x-stackit-marketplace`) received via the vendor's redirect URL

## Getting Started

1. Clone this repository
2. Set up authentication with the STACKIT Go SDK (see [Using the STACKIT Go SDK](#using-the-stackit-go-sdk) below)
3. Set the following environment variables:

   - `MP_REF_CODE_REDIRECT_TOKEN`: The `x-stackit-marketplace` token received during redirect
   - `MP_REF_CODE_VENDOR_PROJECT_ID`: Your vendor STACKIT project ID associated with the product

4. Run the application using either:

```bash
make run
```

or

```bash
go run cmd/main.go
```

## Flow Explanation

The application demonstrates the following flow:

1. **Marketplace token validation**: Validates the marketplace token by:
   - Fetching the public key from STACKIT
   - Validating the token format
   - Verifying the token signature
2. **Customer Resolution**: Uses the validated token to resolve customer information
3. **Subscription Approval**: Approves the marketplace subscription

## Project Structure

```bash
.
├── cmd/
│   └── main.go                         # Application entry point
├── internal/
│   ├── step_1_validate_token.go        # Token validation logic
│   ├── step_2_resolve_customer.go      # Customer resolution logic
│   └── step_3_approve_subscription.go  # Subscription approval logic
├── utils/                              # Utility functions
│   └── utils.go
```

## Development

The project includes several make targets for development:

```bash
make build    # Build the application
make test     # Run tests
make lint     # Run linter
make fmt      # Format code
make clean    # Clean build artifacts
```

## Using the STACKIT Go SDK

The examples in this repository use the official [STACKIT Go SDK](https://github.com/stackitcloud/stackit-sdk-go) to interact with the Marketplace API.

Authentication is handled automatically by the SDK using its default configuration. For details, see the SDK's [Authentication documentation](https://github.com/stackitcloud/stackit-sdk-go?tab=readme-ov-file#authentication).

A **potential** way to authenticate is by setting one of these environment variables:

- `STACKIT_SERVICE_ACCOUNT_KEY`: Your service account key
- `STACKIT_SERVICE_ACCOUNT_KEY_PATH`: Path to a file containing your service account key
- `STACKIT_SERVICE_ACCOUNT_TOKEN`: Your service account access token

## Contributing

Feel free to submit issues and enhancement requests!
