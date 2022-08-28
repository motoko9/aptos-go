package move_example

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
)

func TestCoinInfo(t *testing.T) {
	// coin account
	coinWallet, err := wallet.NewFromKeygenFile("account_usdt")
	if err != nil {
		panic(err)
	}
	coinAddress := coinWallet.Address()
	fmt.Printf("coin address: %s\n", coinAddress)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	//
	coinInfo, err := client.CoinInfo(context.Background(), aptos.USDTCoin, 0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("name: %s, symbol: %s, decimals: %d\n", coinInfo.Name, coinInfo.Symbol, coinInfo.Decimals)
}
