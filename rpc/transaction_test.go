package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_Transaction(t *testing.T) {
	client := New(DevNet_RPC)
	transaction, err := client.Transaction(context.Background(), "0x455253655184e2046c9ede2168c914e52465ad9dffbed94dd42e1089dcd1f066")
	if err != nil {
		panic(err)
	}
	transactionJson, _ := json.Marshal(transaction)
	fmt.Printf("transaction: %s\n", string(transactionJson))
}
