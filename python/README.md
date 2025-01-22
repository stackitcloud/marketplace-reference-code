# STACKIT Marketplace Reference Code in Python

This Python reference code exemplifies how to integrate with the STACKIT Marketplace API as a vendor, showing the essential flow of authenticating and handling marketplace subscriptions.

## Overview

This example application demonstrates three key marketplace integration steps:

1. **Token Validation**: Validates marketplace tokens using public key cryptography
2. **Customer Resolution**: Retrieves customer information from subscription tokens
3. **Subscription Approval**: Processes and approves marketplace subscriptions

## Prerequisites

- Python 3.8 or higher
- A STACKIT Service Account Key
- A marketplace token (`x-stackit-marketplace`) received via the vendor's redirect URL

## Getting Started

1. Clone this repository
2. Create and activate a virtual environment (recommended):

```bash
   python -m venv venv
   source venv/bin/activate  # On Windows: venv\Scripts\activate
```

3. Set up authentication with the STACKIT Python SDK

4. Set the following environment variables:

   - `MP_REF_CODE_REDIRECT_TOKEN`: The `x-stackit-marketplace` token received during redirect
   - `MP_REF_CODE_VENDOR_PROJECT_ID`: Your vendor STACKIT project ID associated with the product

5. Run the application using either:

```bash
   make run
```

or

```bash
   python -m impl.main
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
├── impl
│   ├── __init__.py
│   ├── main.py                         # Application entry point
│   ├── step_1_validate_token.py        # Token validation logic
│   ├── step_2_resolve_customer.py      # Customer resolution logic
│   ├── step_3_approve_subscription.py  # Subscription approval logic
│   └── utils                           # Utility functions
│       ├── __init__.py
│       ├── exceptions.py
│       └── formatting.py
```

## Development

The project includes several make targets for development:

```bash
make lint # Run linter (flake8)
make fmt # Format code (black)
make clean # Clean build artifacts
make deps # Install dependencies
```

## Using the STACKIT Python SDK

The examples in this repository use the official [STACKIT Python SDK](https://github.com/stackitcloud/stackit-sdk-python) to interact with the Marketplace API.

Authentication is handled automatically by the SDK using its default configuration. For details, see the [SDK's authentication documentation](https://github.com/stackitcloud/stackit-sdk-python?tab=readme-ov-file#authorization).

A **potential** way to authenticate is by setting one of these environment variables:

- `STACKIT_SERVICE_ACCOUNT_KEY`: Your service account key
- `STACKIT_SERVICE_ACCOUNT_KEY_PATH`: Path to a file containing your service account key
- `STACKIT_SERVICE_ACCOUNT_TOKEN`: Your service account access token

## Contributing

Feel free to submit issues and enhancement requests!
