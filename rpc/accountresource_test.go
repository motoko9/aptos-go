package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_AccountResources(t *testing.T) {
	client := New(DevNet_RPC)
	accountResources, err := client.AccountResources(context.Background(), "0xdb8e31e499902c188ecd9786862a98f00a09fd1d7257ac9a5a154341318d0aa9", 352973)
	if err != nil {
		panic(err)
	}
	accountResourcesJson, _ := json.Marshal(accountResources)
	fmt.Printf("account resources: %s\n", string(accountResourcesJson))
}

func TestClient_AccountResourceByAddressAndType(t *testing.T) {
	client := New(DevNet_RPC)
	accountResource, err := client.AccountResourceByAddressAndType(
		context.Background(),
		"0xdb8e31e499902c188ecd9786862a98f00a09fd1d7257ac9a5a154341318d0aa9",
		"0x1::coin::CoinStore<0x1::aptos_coin::AptosCoin>",
		352973)
	if err != nil {
		panic(err)
	}
	accountResourceJson, _ := json.Marshal(accountResource)
	fmt.Printf("account resource: %s\n", string(accountResourceJson))
}

func TestClient_AccountBalance(t *testing.T) {
	client := New(DevNet_RPC)
	balance, err := client.AccountBalance(
		context.Background(),
		"0xdb8e31e499902c188ecd9786862a98f00a09fd1d7257ac9a5a154341318d0aa9",
		"USDT",
		352973)
	if err != nil {
		panic(err)
	}
	fmt.Printf("account balance: %d\n", balance)
}
