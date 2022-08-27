package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_Transactions(t *testing.T) {
	client := New(DevNet_RPC)
	transaction, err := client.Transactions(context.Background(), 1000, 10)
	if err != nil {
		panic(err)
	}
	transactionJson, _ := json.MarshalIndent(transaction, "", "    ")
	fmt.Printf("transaction: %s\n", string(transactionJson))
}

func TestClient_TransactionByHash(t *testing.T) {
	client := New(DevNet_RPC)
	transaction, err := client.TransactionByHash(context.Background(), "0x5e57d7a9c9c82b91ebb1ccfaeef748bb2cf159a5cb665b34c95fc0673c693eed")
	if err != nil {
		panic(err)
	}
	transactionJson, _ := json.MarshalIndent(transaction, "", "    ")
	fmt.Printf("transaction: %s\n", string(transactionJson))
}

func TestClient_TransactionByVersion(t *testing.T) {
	client := New(DevNet_RPC)
	transaction, err := client.TransactionByVersion(context.Background(), 1000)
	if err != nil {
		panic(err)
	}
	transactionJson, _ := json.MarshalIndent(transaction, "", "    ")
	fmt.Printf("transaction: %s\n", string(transactionJson))
}

func TestClient_EstimateGasPrice(t *testing.T) {
	client := New(DevNet_RPC)
	gasEstimate, err := client.EstimateGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("estimate gas price: %d\n", gasEstimate)
}
