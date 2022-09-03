package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_Account(t *testing.T) {
	account, err := client.Account(context.Background(), "0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f", 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("account: \n")
	accountJson, _ := json.MarshalIndent(account, "", "    ")
	fmt.Printf(string(accountJson))
}

func TestClient_AccountResources(t *testing.T) {
	ledger, err := client.Ledger(context.Background())
	if err != nil {
		panic(err)
	}
	accountResources, err := client.AccountResources(context.Background(),
		"0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f", ledger.LedgerVersion)
	if err != nil {
		panic(err)
	}
	fmt.Printf("account resources: \n")
	accountResourcesJson, _ := json.MarshalIndent(accountResources, "", "    ")
	fmt.Printf(string(accountResourcesJson))
}

func TestClient_AccountResources_Latest(t *testing.T) {
	accountResources, err := client.AccountResources(context.Background(),
		"0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f", 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("account resources: \n")
	accountResourcesJson, _ := json.MarshalIndent(accountResources, "", "    ")
	fmt.Printf(string(accountResourcesJson))
}

func TestClient_AccountResourceByAddressAndType(t *testing.T) {
	ledger, err := client.Ledger(context.Background())
	if err != nil {
		panic(err)
	}

	accountResource, err := client.AccountResourceByAddressAndType(
		context.Background(),
		"0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f",
		"0x1::coin::CoinStore<0x1::aptos_coin::AptosCoin>",
		ledger.LedgerVersion)
	if err != nil {
		panic(err)
	}
	fmt.Printf("account resource: \n")
	accountResourceJson, _ := json.MarshalIndent(accountResource, "", "    ")
	fmt.Printf(string(accountResourceJson))
}

func TestClient_AccountModules(t *testing.T) {
	ledger, err := client.Ledger(context.Background())
	if err != nil {
		panic(err)
	}
	accountModules, err := client.AccountModules(context.Background(),
		"0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f", ledger.LedgerVersion)
	if err != nil {
		panic(err)
	}

	fmt.Printf("account modules: \n")
	accountModulesJson, _ := json.MarshalIndent(accountModules, "", "    ")
	fmt.Printf(string(accountModulesJson))
}

func TestClient_AccountModuleByAddressAndName(t *testing.T) {
	ledger, err := client.Ledger(context.Background())
	if err != nil {
		panic(err)
	}
	accountModule, err := client.AccountModuleByAddressAndName(context.Background(),
		"0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f", "message", ledger.LedgerVersion)
	if err != nil {
		panic(err)
	}

	fmt.Printf("account module: \n")
	accountModuleJson, _ := json.MarshalIndent(accountModule, "", "    ")
	fmt.Printf(string(accountModuleJson))
}
