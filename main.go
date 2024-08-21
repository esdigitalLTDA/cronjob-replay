package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	checkIntervalHours, _ := strconv.Atoi(os.Getenv("CHECK_INTERVAL_HOURS"))
	minBalance, _ := strconv.ParseInt(os.Getenv("MIN_BALANCE"), 10, 64)
	bridgeWalletAddress := os.Getenv("BRIDGE_WALLET_ADDRESS")
	treasuryWalletAddress := os.Getenv("TREASURY_WALLET_ADDRESS")
	slackWebhookURL := os.Getenv("SLACK_WEBHOOK_URL")

	ethClient, err := ethclient.Dial(os.Getenv("ETH_NODE_URL"))
	if err != nil {
		log.Printf("Error connecting to Ethereum client: %v", err)
		return
	}

	thetaClient, err := ethclient.Dial(os.Getenv("THETA_NODE_URL"))
	if err != nil {
		log.Printf("Error connecting to Theta client: %v", err)
		return
	}

	bridgeWallet := common.HexToAddress(bridgeWalletAddress)
	treasuryWallet := common.HexToAddress(treasuryWalletAddress)

	for {
		// Ethereum balance check and transfer
		checkAndTransfer(ethClient, bridgeWallet, treasuryWallet, minBalance, slackWebhookURL, "Ethereum")

		// Theta balance check and transfer
		checkAndTransfer(thetaClient, bridgeWallet, treasuryWallet, minBalance, slackWebhookURL, "Theta")

		// Wait for the next check interval
		time.Sleep(time.Duration(checkIntervalHours) * time.Hour)
	}
}
