package move_example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/aptosmodule"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
)

func TestCoinInfo(t *testing.T) {
	// coin account
	coinWallet, err := wallet.NewFromKeygenFile("account_usdc")
	if err != nil {
		panic(err)
	}
	coinAddress := coinWallet.Address()
	fmt.Printf("coin address: %s\n", coinAddress)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	//
	resourceType := fmt.Sprintf("%s::usdc::CoinInfo", coinAddress)
	resource, aptosErr := client.AccountResourceByAddressAndType(context.Background(), coinAddress, resourceType, 0)
	if aptosErr != nil {
		panic(aptosErr)
	}

	coinInfo, ok := resource.Object.(*aptosmodule.CoinInfo)
	if !ok {
		panic("resource type is invalid")
	}
	// need to decode, because init with hex
	name, _ := hex.DecodeString(coinInfo.Name)
	symbol, _ := hex.DecodeString(coinInfo.Symbol)
	fmt.Printf("name: %s, symbol: %s, decimals: %d\n", name, symbol, coinInfo.Decimals)
}
