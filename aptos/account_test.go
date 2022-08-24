package aptos

import (
	"context"
	"fmt"
	"testing"
)

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
