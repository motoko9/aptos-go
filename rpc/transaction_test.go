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
	transaction, err := client.TransactionByHash(context.Background(), "0xa78f68c3479e80e0bb4823e4d19956a311a2213c2afd39181ffe75be01d004d2")
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
