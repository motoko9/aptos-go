package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Transaction(t *testing.T) {
	transaction, err := client.TransactionByHash(context.Background(), "0x8da4d51ea5d07edd7059b9a6665c46c2050ef0d31d70b5be7f5884349fb5ca03")
	if err != nil {
		panic(err)
	}
	transactionJson, _ := json.MarshalIndent(transaction, "", "    ")
	fmt.Printf("transaction: %s\n", string(transactionJson))
}

func Test_TransactionByHash(t *testing.T) {
	transaction, err := client.TransactionByHash(context.Background(), "0xa78f68c3479e80e0bb4823e4d19956a311a2213c2afd39181ffe75be01d004d2")
	if err != nil {
		panic(err)
	}
	transactionJson, _ := json.MarshalIndent(transaction, "", "    ")
	fmt.Printf("transaction: %s\n", string(transactionJson))
}

func Test_TransactionByVersion(t *testing.T) {
	transaction, err := client.TransactionByVersion(context.Background(), 1000)
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
