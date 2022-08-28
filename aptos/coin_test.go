package aptos

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_CoinInfo(t *testing.T) {
	client := New(DevNet_RPC)
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
