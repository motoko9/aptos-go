# Aptos SDK library for Go

Go library to interface with Aptos JSON RPC.

## Installation
```bash
$ cd my-project
$ go get github.com/motoko9/aptos-go
```

## Demo
```go
package test

import (
    "context"
    "fmt"
    "github.com/motoko9/aptos-go/aptos"
    "github.com/motoko9/aptos-go/aptosmodule"
    "github.com/motoko9/aptos-go/rpc"
    "github.com/motoko9/aptos-go/rpcmodule"
    "github.com/motoko9/aptos-go/utils"
    "github.com/shopspring/decimal"
    "math/big"
    "testing"
)

type LiquidityPool struct {
    CoinX                *aptosmodule.Coin `json:"coin_x_reserve"`
    CoinY                *aptosmodule.Coin `json:"coin_y_reserve"`
    LatestBlockTimestamp uint64            `json:"last_block_timestamp,string"`
}

func LiquidityPoolCreator() interface{} {
    return &LiquidityPool{}
}

func TestResource(t *testing.T) {
    lpResourceType := "0x190d44266241744264b964a37b8f09863167a12d3e70cda39376cfb4e3561e12::liquidity_pool::LiquidityPool"
    rpcmodule.RegisterResourceObjectCreator(lpResourceType, LiquidityPoolCreator)
    client := aptos.New(rpc.MainNet_RPC, true)
    swapAddress := "0x05a97986a9d031c4567e15b797be516910cfcb4156312482efc6a19c0a30c948"
    res, aptosErr := client.AccountResources(context.Background(), swapAddress, 0)
    if aptosErr != nil {
        panic(aptosErr)
    }
    // register coin
    client.CustomizeCoin("self", "USDC", 6, "USDC", "0xd6d6372c8bde72a7ab825c00b9edd35e643fb94a61c55d9d94a9db3010098548::USDC::Coin")
    for _, item := range *res {
        m, t, err := utils.ExtractFromResource(item.Type)
        if err != nil {
            fmt.Printf("resource is invalid\n")
            continue
        }
        if m != lpResourceType {
            continue
        }
        if len(t) != 3 {
            fmt.Printf("liquidity pool is invalid\n")
            continue
        }

        if item.Object == nil {
            fmt.Printf("resource object is invalid\n")
            continue
        }

        lp := item.Object.(*LiquidityPool)
        if lp.CoinX.Value < 200000000 || lp.CoinY.Value < 200000000 {
            continue
        }

        coin0 := client.FindCoinByType(t[0])
        if coin0 == nil {
            fmt.Printf("coin %s is not supported\n", t[0])
            continue
        }

        coin1 := client.FindCoinByType(t[1])
        if coin1 == nil {
            fmt.Printf("coin %s is not supported\n", t[1])
            continue
        }

        fmt.Printf("(%s %s %s), (%s %s)\n",
            aptos.CoinAlias(coin0.Symbol, coin0.Source),
            aptos.CoinAlias(coin1.Symbol, coin1.Source),
            t[2],
            decimal.NewFromInt(int64(lp.CoinX.Value)).Div(decimal.NewFromBigInt(big.NewInt(1), int32(coin0.Decimals))).StringFixed(2),
            decimal.NewFromInt(int64(lp.CoinY.Value)).Div(decimal.NewFromBigInt(big.NewInt(1), int32(coin1.Decimals))).StringFixed(2),
        )
    }
}
```
## Example

* [account](./examples/account)
* [move basic](./examples/move_basic)
* [token basic](./examples/token)

