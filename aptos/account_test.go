package aptos

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/motoko9/aptos-go/rpc"
	"testing"
)

func TestClient_AccountBalance(t *testing.T) {
	client := New(rpc.TestNet_RPC, false)
	ledger, err := client.Ledger(context.Background())
	if err != nil {
		panic(err)
	}
	{
		balance, err := client.AccountBalance(
			context.Background(),
			"0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f",
			CoinAlias("APT", "aptos"),
			ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("account Aptos balance: %d\n", balance)
	}
	{
		balance, err := client.AccountBalance(
			context.Background(),
			"0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f",
			CoinAlias("USDT", "wormhole"),
			ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("account USDT balance: %d\n", balance)
	}
	{
		balance, err := client.AccountBalance(
			context.Background(),
			"0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f",
			CoinAlias("BTC", "wormhole"),
			ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("account BTC balance: %d\n", balance)
	}
}

func TestClient_AccountResources_Latest(t *testing.T) {
	client := New(rpc.TestNet_RPC, false)
	accountResources, err := client.AccountResources(context.Background(),
		"0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f", "", 0, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("account resources: \n")
	accountResourcesJson, _ := json.MarshalIndent(accountResources, "", "    ")
	fmt.Printf(string(accountResourcesJson))
}
