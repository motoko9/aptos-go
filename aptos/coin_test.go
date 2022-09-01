package aptos

import (
    "context"
    "fmt"
    "github.com/hashicorp/go-hclog"
    "testing"
)

func TestClient_CoinInfo(t *testing.T) {
    client := NewClient(DevNet_RPC, hclog.Default())
    {
        aptosCoinInfo, err := client.CoinInfo(context.Background(), AptosCoin, 0)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%v\n", aptosCoinInfo)
    }
    {
        usdtCoinInfo, err := client.CoinInfo(context.Background(), USDTCoin, 0)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%v\n", usdtCoinInfo)
    }
}
