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
	accountJson, err := json.MarshalIndent(account, "", "    ")
	if err != nil {
		panic(err)
	}
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
	accountResourcesJson, _ := json.MarshalIndent(accountResources, "", "    ")
	fmt.Printf("account resources: %s\n", string(accountResourcesJson))
}

func TestClient_AccountResources_Latest(t *testing.T) {
	accountResources, err := client.AccountResources(context.Background(),
		"0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f", 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("account resources: \n")
	accountJson, err := json.MarshalIndent(accountResources, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(accountJson))
}

func TestClient_AccountResourceByAddressAndType(t *testing.T) {
	ledger, err := client.Ledger(context.Background())
	if err != nil {
		panic(err)
	}

	var a interface{}
	err = client.AccountResourceByAddressAndType(
		context.Background(),
		"0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f",
		"0x1::coin::CoinStore<0x1::aptos_coin::AptosCoin>",
		ledger.LedgerVersion, &a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("account resource: \n")
	accountJson, err := json.MarshalIndent(accountResources, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(accountJson))
}

func TestClient_AccountModules(t *testing.T) {
	ledger, err := client.Ledger(context.Background())
	if err != nil {
		panic(err)
	}
	accountResources, err := client.AccountModules(context.Background(),
		"0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f", ledger.LedgerVersion)
	if err != nil {
		panic(err)
	}
	accountResourcesJson, _ := json.MarshalIndent(accountResources, "", "    ")
	fmt.Printf("account modules: %s\n", string(accountResourcesJson))
}

func TestClient_AccountModuleByAddressAndName(t *testing.T) {
	ledger, err := client.Ledger(context.Background())
	if err != nil {
		panic(err)
	}
	accountResources, err := client.AccountModuleByAddressAndName(context.Background(),
		"0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f", "message", ledger.LedgerVersion)
	if err != nil {
		panic(err.Error())
	}
	accountResourcesJson, _ := json.MarshalIndent(accountResources, "", "    ")
	fmt.Printf("account modules: %s\n", string(accountResourcesJson))
}
