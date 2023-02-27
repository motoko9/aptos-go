package move_example

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/rpcmodule"
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
	fmt.Printf("coin address: %s\n", coinAddress)

	// new rpc
	client := aptos.New(rpc.TestNet_RPC, false)

	//
	coinInfo, aptosErr := client.CoinInfo(context.Background(), aptos.CoinAlias("USDT", "wormhole"), 0)
	if aptosErr != nil {
		panic(aptosErr)
	}

	fmt.Printf("name: %s, symbol: %s, decimals: %d\n", coinInfo.Name, coinInfo.Symbol, coinInfo.Decimals)
}

func TestCoinInfo1(t *testing.T) {
	// token account
	coinWallet, err := wallet.NewFromKeygenFile("account_usdt")
	if err != nil {
		panic(err)
	}
	coinAddress := coinWallet.Address()
	fmt.Printf("coin address: %s\n", coinAddress)

	// recipient account
	userWallet, err := wallet.NewFromKeygenFile("account_mint")
	if err != nil {
		panic(err)
	}
	userAddress := userWallet.Address()
	fmt.Printf("recipient address: %s\n", userAddress)

	// new rpc
	client := aptos.New(rpc.TestNet_RPC, false)

	//
	raw, aptosErr := client.View(context.Background(), &rpcmodule.ViewRequest{
		Function:      "0x1::coin::balance",
		TypeArguments: []string{client.FindCoinBySymbolSource("USDT", "wormhole").T},
		Arguments:     []interface{}{userAddress},
	})
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("balance: %s\n", raw)

	//
	raw, aptosErr = client.View(context.Background(), &rpcmodule.ViewRequest{
		Function:      "0x1::coin::name",
		TypeArguments: []string{client.FindCoinBySymbolSource("USDT", "wormhole").T},
		Arguments:     []interface{}{},
	})
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("name: %s\n", raw)

	//
	raw, aptosErr = client.View(context.Background(), &rpcmodule.ViewRequest{
		Function:      "0x1::coin::symbol",
		TypeArguments: []string{client.FindCoinBySymbolSource("USDT", "wormhole").T},
		Arguments:     []interface{}{},
	})
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("symbol: %s\n", raw)
}
