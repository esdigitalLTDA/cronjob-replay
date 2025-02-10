#### Prerequisites

- **Go version 1.23 or higher** is required if you are running the cronjob locally.

---

#### Environment Variables

Create a `.env` file in the root directory of the project with the following variables:

```env
# General Settings
CHECK_INTERVAL_HOURS=6                    # Interval in hours to check the bridge wallet balance
MIN_BALANCE=4500000                       # Minimum balance threshold in wei to trigger a transfer
TRANSFER_AMOUNT_WEI=1000000000000000000     # Amount to transfer in wei (adjust as needed)

# Ethereum Settings
ETH_NODE_URL=https://your-ethereum-rpc-url  # Ethereum network RPC URL (e.g., Infura)

# Theta Settings
THETA_NODE_URL=https://your-theta-rpc-url    # Theta network RPC URL

# Wallet Settings
BRIDGE_WALLET_ADDRESS=0xYourBridgeWalletAddress    # Address of the bridge wallet
TREASURY_WALLET_ADDRESS=0xYourTreasuryWalletAddress  # Address of the treasury wallet
TREASURY_PRIVATE_KEY=your_treasury_private_key_here  # Private key of the treasury wallet

# Slack Notifications
SLACK_WEBHOOK_URL=https://hooks.slack.com/services/your/slack/webhook/url  # Slack Webhook URL for notifications
SLACK_CHANNEL="#alerts"
```

---

#### Running the Cronjob Locally

To run the cronjob locally, navigate to the root directory of the project and use the following command:

```bash
go run .
```

---

#### Running the Cronjob with Docker

1. **Build the Docker Image**

   In the root directory of your project (where the `Dockerfile` is located), run:

   ```bash
   docker build -t cronjob-replay .
   ```

   This command creates a Docker image named `cronjob-replay`.

2. **Run the Docker Container**

   After building the image, start the container using:

   ```bash
   docker run --env-file .env cronjob-replay
   ```

   > **Note:** The `--env-file .env` option passes the environment variables from the `.env` file into the container.

---

#### Running the Cronjob with Docker Compose

1. **Build and Run with Docker Compose**

   In the root directory of your project, run:

   ```bash
   docker compose up --build -d
   ```

   This command will:
   - Build the Docker image (if needed),
   - Start the container in detached mode,
   - Automatically load environment variables from your `.env` file as specified in the Compose file.

---

### Summary

- **Local Development:** Use `go run .` to run the cronjob directly with Go.
- **Docker Run:** Build the image with `docker build` and start it with `docker run --env-file .env cronjob-replay`.
- **Docker Compose:** Simply run `docker compose up --build -d` for a simplified multi-container setup (or single container, in this case), with the environment automatically loaded.