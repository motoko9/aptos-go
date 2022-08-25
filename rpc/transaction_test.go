package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_Transactions(t *testing.T) {
	client := New(DevNet_RPC)
	transaction, err := client.Transactions(context.Background(), 1000, 1)
	if err != nil {
		panic(err)
	}
	transactionJson, _ := json.MarshalIndent(transaction, "", "    ")
	fmt.Printf("transaction: %s\n", string(transactionJson))
}

func TestClient_TransactionByHash(t *testing.T) {
	client := New(DevNet_RPC)
	transaction, err := client.TransactionByHash(context.Background(), "0xfd20a0e92e10470ed582d8d3bc562a628c412f67116eee37182bdee4df9dbcdf")
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
