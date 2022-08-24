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

func TestClient_AccountResources(t *testing.T) {
	client := New(DevNet_RPC)
	ledger, err := client.Ledger(context.Background())
	if err != nil {
		panic(err)
	}
	accountResources, err := client.AccountResources(context.Background(),
		"0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f", ledger.LedgerVersion)
	if err != nil {
		panic(err)
	}
	accountResourcesJson, _ := json.MarshalIndent(accountResources, "", "    ")
	fmt.Printf("account resources: %s\n", string(accountResourcesJson))
}

func TestClient_AccountResources_Latest(t *testing.T) {
	client := New(DevNet_RPC)
	accountResources, err := client.AccountResources(context.Background(),
		"0xdb8e31e499902c188ecd9786862a98f00a09fd1d7257ac9a5a154341318d0aa9", 0)
	if err != nil {
		panic(err)
	}
	accountResourcesJson, _ := json.MarshalIndent(accountResources, "", "    ")
	fmt.Printf("account resources: %s\n", string(accountResourcesJson))
}

func TestClient_AccountResourceByAddressAndType(t *testing.T) {
	client := New(DevNet_RPC)
	ledger, err := client.Ledger(context.Background())
	if err != nil {
		panic(err)
	}
	accountResource, err := client.AccountResourceByAddressAndType(
		context.Background(),
		"0xdb8e31e499902c188ecd9786862a98f00a09fd1d7257ac9a5a154341318d0aa9",
		"0x1::coin::CoinStore<0x1::aptos_coin::AptosCoin>",
		ledger.LedgerVersion)
	if err != nil {
		panic(err)
	}
	accountResourceJson, _ := json.MarshalIndent(accountResource, "", "    ")
	fmt.Printf("account resource: %s\n", string(accountResourceJson))
}
