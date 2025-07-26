# MCP Client

A Python client for the Fi MCP (Model-Controller-Presenter) server. This client provides a simple interface to interact with the MCP server and access various financial data tools.

## Features

- Simple interface to call all available MCP tools
- Handles authentication flow automatically
- Parses and returns structured data from the server
- Provides convenience methods for each tool
- Includes a test script to demonstrate usage

## Available Tools

The MCP server provides the following tools:

1. `fetch_net_worth` - Calculate comprehensive net worth using data from connected accounts
2. `fetch_credit_report` - Retrieve comprehensive credit report including scores, loans, etc.
3. `fetch_epf_details` - Retrieve detailed EPF (Employee Provident Fund) account information
4. `fetch_mf_transactions` - Retrieve detailed mutual fund transaction history
5. `fetch_bank_transactions` - Retrieve detailed bank transactions for each bank account
6. `fetch_stock_transactions` - Retrieve detailed Indian stock transactions

## Tool Response Formats

Each tool returns data in a specific format. Below are the details of the response formats for each tool:

### fetch_net_worth

Returns comprehensive information about a user's net worth, including assets, liabilities, and detailed breakdowns.

**Response Structure:**
```json
{
  "netWorthResponse": {
    "assetValues": [
      {
        "netWorthAttribute": "ASSET_TYPE_MUTUAL_FUND",
        "value": {
          "currencyCode": "INR",
          "units": "169919"
        }
      }
    ],
    "totalNetWorthValue": {
      "currencyCode": "INR",
      "units": "1721734"
    }
  },
  "mfSchemeAnalytics": {
    "schemeAnalytics": [
      {
        "schemeDetail": {
          "amc": "NIPPON_INDIA_MUTUAL_FUND",
          "nameData": {
            "longName": "Nippon India Corporate Bond Fund - Direct Growth"
          }
        },
        "enrichedAnalytics": {
          "analytics": {
            "schemeDetails": {
              "currentValue": {
                "currencyCode": "INR",
                "units": "59058"
              },
              "investedValue": {
                "currencyCode": "INR",
                "units": "52000"
              },
              "XIRR": 8.15
            }
          }
        }
      }
    ]
  },
  "accountDetailsBulkResponse": {
    "accountDetailsMap": {}
  }
}
```

### fetch_credit_report

Returns detailed credit report information including credit score, account details, and inquiry history.

**Response Structure:**
```json
{
  "creditReports": [
    {
      "creditReportData": {
        "userMessage": {
          "userMessageText": "Normal Response"
        },
        "creditProfileHeader": {
          "reportDate": "20240521",
          "reportTime": "104515"
        },
        "currentApplication": {
          "currentApplicationDetails": {
            "enquiryReason": "5",
            "amountFinanced": "0",
            "durationOfAgreement": "0"
          }
        },
        "creditAccount": {
          "creditAccountSummary": {
            "account": {
              "creditAccountTotal": "3",
              "creditAccountActive": "3",
              "creditAccountDefault": "1"
            },
            "totalOutstandingBalance": {
              "outstandingBalanceSecured": "110000",
              "outstandingBalanceSecuredPercentage": "36",
              "outstandingBalanceUnSecured": "196000"
            }
          },
          "creditAccountDetails": [
            {
              "subscriberName": "Axis Bank",
              "portfolioType": "R",
              "accountType": "10",
              "openDate": "20220120",
              "creditLimitAmount": "75000"
            }
          ]
        },
        "score": {
          "bureauScore": "621",
          "bureauScoreConfidenceLevel": "M"
        }
      },
      "vendor": "EXPERIAN"
    }
  ]
}
```

### fetch_epf_details

Returns detailed information about Employee Provident Fund accounts.

**Response Structure:**
```json
{
  "uanAccounts": [
    {
      "phoneNumber": {},
      "rawDetails": {
        "est_details": [
          {
            "est_name": "GLOBAL LOGISTICS SOLUTIONS",
            "member_id": "MHPN*****************",
            "office": "(RO)PUNE",
            "doj_epf": "10-07-2017",
            "doe_epf": "30-04-2020",
            "pf_balance": {
              "net_balance": "250000",
              "employee_share": {
                "credit": "115000",
                "balance": "115000"
              },
              "employer_share": {
                "credit": "115000",
                "balance": "115000"
              }
            }
          }
        ],
        "overall_pf_balance": {
          "pension_balance": "185000",
          "current_pf_balance": "635000"
        }
      }
    }
  ]
}
```

### fetch_mf_transactions

Returns detailed transaction history for mutual fund investments.

**Response Structure:**
```json
{
  "mfTransactions": [
    {
      "isin": "INF179KB1HS3",
      "schemeName": "Nippon India Corporate Bond Fund - Direct Growth",
      "folioId": "2002011001",
      "txns": [
        [1, "2022-04-01", 18.50, 432.4324, 8000],
        [1, "2022-05-01", 18.56, 430.8937, 8000]
      ]
    }
  ],
  "schemaDescription": "A list of mutual fund investments. Each 'txns' field is a list of data arrays with schema: [ orderType (1 for BUY and 2 for SELL), transactionDate, purchasePrice, purchaseUnits, transactionAmount ]."
}
```

### fetch_bank_transactions

Returns detailed transaction history for bank accounts.

**Response Structure:**
```json
{
  "schemaDescription": "A list of bank transactions. Each 'txns' field is a list of data arrays with schema: [transactionAmount, transactionNarration, transactionDate, transactionType (1 for CREDIT, 2 for DEBIT, 3 for OPENING, 4 for INTEREST, 5 for TDS, 6 for INSTALLMENT, 7 for CLOSING and 8 for OTHERS), transactionMode, currentBalance].",
  "bankTransactions": [
    {
      "bank": "State Bank of India",
      "txns": [
        ["145000", "SALARY CREDIT - MINISTRY OF FINANCE - JULY 2024", "2024-07-01", 1, "NEFT", "385000"],
        ["10000", "NEFT DR-ICIC0003344-TRANSFER FOR RD", "2024-07-02", 2, "FT", "375000"]
      ]
    },
    {
      "bank": "ICICI Bank",
      "txns": [
        ["10000", "NEFT CR-SBIN0000556-FUNDS FROM SBI", "2024-07-02", 1, "FT", "12500"],
        ["10000", "AUTO DEBIT - RD INSTALLMENT A/C XXXXXX3344", "2024-07-03", 6, "ACH", "2500"]
      ]
    }
  ]
}
```

### fetch_stock_transactions

Returns detailed transaction history for stock investments.

**Response Structure:**
```json
{
  "schemaDescription": "A list of stock transactions. Each 'txns' field is a list of data arrays with schema: [transactionType (1 for BUY, 2 for SELL, 3 for BONUS, 4 for SPLIT), transactionDate, quantity, navValue]. nav value may not be present in some of the transactions",
  "stockTransactions": [
    {
      "isin": "INE0BWS23018",
      "txns": [
        [1, "2023-05-04", 100],
        [1, "2023-05-04", 170]
      ]
    },
    {
      "isin": "INF204KB14I5",
      "txns": [
        [1, "2023-05-04", 100, 10.51]
      ]
    }
  ]
}
```

## Installation

1. Make sure you have Python 3.6+ installed
2. Install the required packages:
   ```
   pip install requests
   ```

## Usage

### Basic Usage

```python
# Import the MCPClient class from the mcp_client module
from mcp_client import MCPClient

# Create a client
client = MCPClient()

# Call a specific tool
bank_transactions = client.fetch_bank_transactions()

# Display the result
# In your actual code, you would use print() or other methods to display the data
```

### Command Line Usage

You can also use the client from the command line:

```bash
python mcp_client.py fetch_bank_transactions
```

### Testing All Tools

To test all available tools, run the test script:

```bash
python test_mcp_client.py
```

To test a specific tool:

```bash
python test_mcp_client.py fetch_net_worth
```

## Authentication

The client handles authentication automatically. When you call a tool for the first time, you'll be prompted to visit a login URL. After logging in, press Enter to continue, and the client will retry the tool call with the authenticated session.

## Configuration

You can configure the client by passing parameters to the constructor:

```python
# Import the MCPClient class from the mcp_client module
from mcp_client import MCPClient

# Custom session ID and server URL
client = MCPClient(
    session_id="your-session-id",
    server_url="https://your-mcp-server.com/mcp/stream"
)
```

## Example: Accessing Bank Transactions

```python
# Import the necessary modules
from mcp_client import MCPClient

# Create a client
client = MCPClient()

# Fetch bank transactions
bank_transactions = client.fetch_bank_transactions()

# Access the schema description
schema_description = bank_transactions.get("schemaDescription")
# Display the schema description
# Example output: "A list of bank transactions. Each 'txns' field is a list of data arrays..."

# Access the bank transactions
for bank in bank_transactions.get("bankTransactions", []):
    bank_name = bank.get("bank")
    transactions = bank.get("txns", [])
    # Display bank information
    # Example output: "Bank: State Bank of India"
    # Example output: "Number of transactions: 14"

    # Process the first transaction
    if transactions:
        first_txn = transactions[0]
        amount = first_txn[0]
        narration = first_txn[1]
        date = first_txn[2]
        # Display transaction information
        # Example output: "First transaction: 145000 on 2024-07-01 - SALARY CREDIT - MINISTRY OF FINANCE - JULY 2024"
```

## Error Handling

The client includes basic error handling for HTTP errors and authentication issues. If a tool call fails, an exception will be raised with details about the error.

## Contributing

Feel free to contribute to this client by submitting issues or pull requests.
