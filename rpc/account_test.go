package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_Account(t *testing.T) {
	client := New(DevNet_RPC)
	account, err := client.Account(context.Background(), "0xdb8e31e499902c188ecd9786862a98f00a09fd1d7257ac9a5a154341318d0aa9")
	if err != nil {
		panic(err)
	}
	accountJson, _ := json.Marshal(account)
	fmt.Printf("account: %s\n", string(accountJson))
}
