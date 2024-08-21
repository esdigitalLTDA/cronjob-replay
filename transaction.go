package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func transferFromTreasury(client *ethclient.Client, bridgeWallet common.Address) (string, error) {
	privateKey := os.Getenv("TREASURY_PRIVATE_KEY")
	if privateKey == "" {
		return "", fmt.Errorf("treasury private key not set")
	}

	// Parse the private key
	treasuryKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}

	// Get the nonce for the treasury wallet
	treasuryAddress := crypto.PubkeyToAddress(treasuryKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), treasuryAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %v", err)
	}

	// Get the gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to suggest gas price: %v", err)
	}

	// Get the transfer amount from the environment
	transferAmountWei := os.Getenv("TRANSFER_AMOUNT_WEI")
	amount := new(big.Int)
	amount, ok := amount.SetString(transferAmountWei, 10)
	if !ok {
		return "", fmt.Errorf("invalid transfer amount")
	}

	// Set up the transaction
	gasLimit := uint64(21000) // For standard ETH transfer
	tx := types.NewTransaction(nonce, bridgeWallet, amount, gasLimit, gasPrice, nil)

	// Sign the transaction
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get network ID: %v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), treasuryKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %v", err)
	}

	// Return the transaction hash
	return signedTx.Hash().Hex(), nil
}
