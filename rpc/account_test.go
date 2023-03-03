package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_Account(t *testing.T) {
	account, err := client.Account(context.Background(), "0x74f3bbe39c7e2793a2e5445ee0336c9ac3191534762b41dcfc1054ad077ccc7c", 0)
	if err != nil {
		panic(err)
	}
	accountJson, _ := json.MarshalIndent(account, "", "    ")
	fmt.Printf(string(accountJson))
}

func TestClient_AccountResources(t *testing.T) {
	ledger, err := client.Ledger(context.Background())
	if err != nil {
		panic(err)
	}
	accountResources, err := client.AccountResources(context.Background(),
		"0x74f3bbe39c7e2793a2e5445ee0336c9ac3191534762b41dcfc1054ad077ccc7c", "", 0, ledger.LedgerVersion)
	if err != nil {
		panic(err)
	}
	accountResourcesJson, _ := json.MarshalIndent(accountResources, "", "    ")
	fmt.Printf(string(accountResourcesJson))
}

func TestClient_AccountResources_Latest(t *testing.T) {
	accountResources, err := client.AccountResources(context.Background(),
		"0x74f3bbe39c7e2793a2e5445ee0336c9ac3191534762b41dcfc1054ad077ccc7c", "", 0, 0)
	if err != nil {
		panic(err)
	}
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
		"0x74f3bbe39c7e2793a2e5445ee0336c9ac3191534762b41dcfc1054ad077ccc7c",
		"0x1::token::CoinStore<0x1::aptos_coin::AptosCoin>",
		ledger.LedgerVersion)
	if err != nil {
		panic(err)
	}
	accountResourceJson, _ := json.MarshalIndent(accountResource, "", "    ")
	fmt.Printf(string(accountResourceJson))
}

func TestClient_AccountModules(t *testing.T) {
	ledger, err := client.Ledger(context.Background())
	if err != nil {
		panic(err)
	}
	accountModules, err := client.AccountModules(context.Background(),
		"0x74f3bbe39c7e2793a2e5445ee0336c9ac3191534762b41dcfc1054ad077ccc7c", ledger.LedgerVersion)
	if err != nil {
		panic(err)
	}
	accountModulesJson, _ := json.MarshalIndent(accountModules, "", "    ")
	fmt.Printf(string(accountModulesJson))
}

func TestClient_AccountModuleByAddressAndName(t *testing.T) {
	ledger, err := client.Ledger(context.Background())
	if err != nil {
		panic(err)
	}
	accountModule, err := client.AccountModuleByAddressAndName(context.Background(),
		"0x74f3bbe39c7e2793a2e5445ee0336c9ac3191534762b41dcfc1054ad077ccc7c", "message", ledger.LedgerVersion)
	if err != nil {
		panic(err)
	}

	accountModuleJson, _ := json.MarshalIndent(accountModule, "", "    ")
	fmt.Printf(string(accountModuleJson))
}
