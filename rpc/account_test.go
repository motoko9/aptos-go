package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_Account(t *testing.T) {
	client := New(DevNet_RPC)
	account, err := client.Account(context.Background(), "0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f")
	if err != nil {
		panic(err)
	}
	accountJson, _ := json.MarshalIndent(account, "", "    ")
	fmt.Printf("account: %s\n", string(accountJson))
}
