package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_Ledger(t *testing.T) {
	client := New(DevNet_RPC)
	ledger, err := client.Ledger(context.Background())
	if err != nil {
		panic(err)
	}
	ledgerJson, _ := json.Marshal(ledger)
	fmt.Printf("ledger: %s\n", string(ledgerJson))
}
