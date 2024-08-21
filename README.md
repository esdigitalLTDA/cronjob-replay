#### Prerequisites

- **Go version 1.23 or higher** is required if you are running the cronjob locally.

#### Environment Variables

Create a `.env` file in the root directory of the project with the following variables:

```env
# General Settings
CHECK_INTERVAL_HOURS=6              # Interval in hours to check the bridge wallet balance
MIN_BALANCE=4500000                 # Minimum balance threshold in wei to trigger a transfer
TRANSFER_AMOUNT_WEI=1000000000000000000  # Amount to transfer in wei (adjust as needed)

# Ethereum Settings
ETH_NODE_URL=https://your-ethereum-rpc-url  # Ethereum network RPC URL (e.g., Infura)

# Theta Settings
THETA_NODE_URL=https://your-theta-rpc-url   # Theta network RPC URL

# Wallet Settings
BRIDGE_WALLET_ADDRESS=0xYourBridgeWalletAddress   # Address of the bridge wallet
TREASURY_WALLET_ADDRESS=0xYourTreasuryWalletAddress  # Address of the treasury wallet
TREASURY_PRIVATE_KEY=your_treasury_private_key_here  # Private key of the treasury wallet

# Slack Notifications
SLACK_WEBHOOK_URL=https://hooks.slack.com/services/your/slack/webhook/url  # Slack Webhook URL for notifications
```

#### Running the Cronjob Locally

To run the cronjob locally, navigate to the root directory of the project and use the following command:

```bash
go run .
```

#### Running the Cronjob with Docker

1. **Build the Docker Image:**

   In the root directory of your project (where the `Dockerfile` is located), run the following command to build the Docker image:

   ```bash
   docker build -t cronjob-replay .
   ```

   This command will create a Docker image with the name `cronjob-replay`.

2. **Run the Docker Container:**

   After building the Docker image, you can run the container using the following command:

   ```bash
   docker run --env-file .env cronjob-replay
   ```

### Notes:
- **Environment Variables:** The `--env-file .env` option passes the environment variables from the `.env` file into the Docker container.