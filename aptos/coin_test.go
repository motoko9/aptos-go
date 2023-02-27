package aptos

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/rpc"
	"testing"
)

func TestClient_CoinInfo(t *testing.T) {
	client := New(rpc.TestNet_RPC, false)
	{
		aptosCoinInfo, err := client.CoinInfo(context.Background(), CoinAlias("APT", "aptos"), 0)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", aptosCoinInfo)
	}
	{
		usdtCoinInfo, err := client.CoinInfo(context.Background(), CoinAlias("USDT", "wormhole"), 0)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", usdtCoinInfo)
	}
}
