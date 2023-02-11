package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Transactions(t *testing.T) {
	transaction, err := client.Transactions(context.Background(), 425348794, 5)
	if err != nil {
		panic(err)
	}
	transactionJson, _ := json.MarshalIndent(transaction, "", "    ")
	fmt.Printf("transaction: %s\n", string(transactionJson))
}

func Test_TransactionByHash(t *testing.T) {
	transaction, err := client.TransactionByHash(context.Background(), "0xe77db8c8612ffc9c7ac779817c674dcc95f1dbf9ccd830aa4da51de2511ef725")
	if err != nil {
		panic(err)
	}
	transactionJson, _ := json.MarshalIndent(transaction, "", "    ")
	fmt.Printf("transaction: %s\n", string(transactionJson))
}

func Test_TransactionByVersion(t *testing.T) {
	transaction, err := client.TransactionByVersion(context.Background(), 425348797)
	if err != nil {
		panic(err)
	}
	transactionJson, _ := json.MarshalIndent(transaction, "", "    ")
	fmt.Printf("transaction: %s\n", string(transactionJson))
}

func Test_TransactionsByAccount(t *testing.T) {
	transaction, err := client.TransactionsByAccount(context.Background(), "0x74f3bbe39c7e2793a2e5445ee0336c9ac3191534762b41dcfc1054ad077ccc7c", 0, 5)
	if err != nil {
		panic(err)
	}
	transactionJson, _ := json.MarshalIndent(transaction, "", "    ")
	fmt.Printf("transaction: %s\n", string(transactionJson))
}

func Test_EstimateGasPrice(t *testing.T) {
	gasEstimate, err := client.EstimateGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("estimate gas price: %d\n", gasEstimate)
}
