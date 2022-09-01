package aptos

import (
    "context"
    "fmt"
    "github.com/hashicorp/go-hclog"
    "github.com/motoko9/aptos-go/common/jsonutil"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestClient_AccountBalance(t *testing.T) {
    client := NewClient(DevNet_RPC, hclog.Default())
    ledger, err := client.Ledger(context.Background())
    assert.NoError(t, err)
    {
        balance, err := client.AccountBalance(
            context.Background(),
            "0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f",
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
            "0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f",
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
            "0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f",
            BTCCoin,
            ledger.LedgerVersion)
        if err != nil {
            panic(err)
        }
        fmt.Printf("account BTC balance: %d\n", balance)
    }
}

func TestClient_AccountResources_Latest(t *testing.T) {
    client := NewClient(DevNet_RPC, hclog.Default())
    accountResources, err := client.AccountResources(context.Background(),
        "0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f", 0)
    if err != nil {
        panic(err)
    }
    fmt.Printf("account resources: \n")
    jsonutil.PrintJsonStringWithIndent(accountResources)
}
