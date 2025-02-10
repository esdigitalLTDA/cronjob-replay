package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func checkAndTransfer(client *ethclient.Client, bridgeWallet, treasuryWallet common.Address, minBalance int64, network string) {
	// Check bridge wallet balance
	balance, err := client.BalanceAt(context.Background(), bridgeWallet, nil)
	if err != nil {
		log.Printf("Error fetching %s bridge wallet balance: %v", network, err)
		return
	}

	balanceInEther := weiToEther(balance)
	fmt.Printf("%s bridge wallet balance: %s ETH\n", network, balanceInEther.String())

	if balance.Cmp(big.NewInt(minBalance)) < 0 {
		fmt.Printf("Balance below minimum on %s, checking treasury balance...\n", network)

		// Check treasury wallet balance
		treasuryBalance, err := client.BalanceAt(context.Background(), treasuryWallet, nil)
		if err != nil {
			log.Printf("Error fetching treasury balance on %s: %v", network, err)
			return
		}

		treasuryBalanceInEther := weiToEther(treasuryBalance)

		if treasuryBalance.Cmp(big.NewInt(0)) <= 0 {
			// Notify Slack that treasury is out of balance
			sendSlackNotification(fmt.Sprintf(
				":warning: *ALERT* :warning:\n"+
					"The *Treasury Wallet* on *%s* is out of balance!\n"+
					"🔴 Current balance: *0 ETH*\n"+
					"⚠️ Please investigate and fund the wallet immediately.",
				network,
			))
			return
		}

		// Perform transfer from treasury to bridge
		txHash, err := transferFromTreasury(client, bridgeWallet)
		if err != nil {
			log.Printf("Error transferring funds on %s: %v", network, err)
			return
		}

		// Notify Slack about the transfer
		sendSlackNotification(fmt.Sprintf(
			":money_with_wings: *Funds Transferred Successfully!*\n"+
				"Network: *%s*\n"+
				"💰 Transferred from *Treasury Wallet* to *Bridge Wallet*.\n"+
				"💲 Treasury Balance Before Transfer: *%s ETH*\n"+
				"📜 Transaction Hash: `%s`\n"+
				"✅ Please verify the transaction on the blockchain.",
			network,
			treasuryBalanceInEther.String(),
			txHash,
		))
	}
}

func weiToEther(wei *big.Int) *big.Float {
	ether := new(big.Float).SetInt(wei)
	return new(big.Float).Quo(ether, big.NewFloat(1e18))
}
