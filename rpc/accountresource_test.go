package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_AccountResources(t *testing.T) {
	client := New(DevNet_RPC)
	ledger, err := client.Ledger(context.Background())
	if err != nil {
		panic(err)
	}
	accountResources, err := client.AccountResources(context.Background(),
		"0xdb8e31e499902c188ecd9786862a98f00a09fd1d7257ac9a5a154341318d0aa9", ledger.LedgerVersion)
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

func TestClient_AccountBalance(t *testing.T) {
	client := New(DevNet_RPC)
	ledger, err := client.Ledger(context.Background())
	if err != nil {
		panic(err)
	}
	{
		balance, err := client.AccountBalance(
			context.Background(),
			"0xdb8e31e499902c188ecd9786862a98f00a09fd1d7257ac9a5a154341318d0aa9",
			AptosCoin,
			ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("account Aptos balance: %d\n", balance)
	}
	{
		balance, err := client.AccountBalance(
			context.Background(),
			"0xdb8e31e499902c188ecd9786862a98f00a09fd1d7257ac9a5a154341318d0aa9",
			USDTCoin,
			ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("account USDT balance: %d\n", balance)
	}
	{
		balance, err := client.AccountBalance(
			context.Background(),
			"0xdb8e31e499902c188ecd9786862a98f00a09fd1d7257ac9a5a154341318d0aa9",
			BTCCoin,
			ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("account BTC balance: %d\n", balance)
	}
}
