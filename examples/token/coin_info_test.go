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
	// token account
	coinWallet, err := wallet.NewFromKeygenFile("account_usdt")
	if err != nil {
		panic(err)
	}
	coinAddress := coinWallet.Address()
	fmt.Printf("token address: %s\n", coinAddress)

	// new rpc
	client := aptos.New(rpc.TestNet_RPC)

	//
	coinInfo, aptosErr := client.CoinInfo(context.Background(), aptos.USDTCoin, 0)
	if aptosErr != nil {
		panic(aptosErr)
	}

	fmt.Printf("name: %s, symbol: %s, decimals: %d\n", coinInfo.Name, coinInfo.Symbol, coinInfo.Decimals)
}
