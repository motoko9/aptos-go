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
			fmt.Printf("liquidity pool is invalid")
			continue
		}

		if item.Object == nil {
			fmt.Printf("resource object is invalid")
			continue
		}

		lp := item.Object.(*LiquidityPool)
		if lp.CoinX.Value < 200000000 || lp.CoinY.Value < 200000000 {
			continue
		}

		name0, r := client.GetCoinName(t[0])
		if !r {
			name0 = t[0]
		}

		name1, r := client.GetCoinName(t[1])
		if !r {
			name1 = t[1]
		}
		fmt.Printf("(%s %s %s), (%d %d)\n", name0, name1, t[2], lp.CoinX.Value, lp.CoinY.Value)
	}
}
```
## Example

* [account](./examples/account)
* [move basic](./examples/move_basic)
* [token basic](./examples/token)

