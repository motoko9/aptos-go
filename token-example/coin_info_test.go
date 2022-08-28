package move_example

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
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
	resource, err := client.AccountResourceByAddressAndType(context.Background(), coinAddress, resourceType, 0)
	if err != nil {
		panic(err)
	}

	//
	type CoinInfo struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Decimals uint64 `json:"decimals,string"`
		Supply   uint64 `json:"supply,string"`
	}
	if resource.Type != resourceType {
		panic(" resource type error")
	}
	var coinInfo CoinInfo
	if err = json.Unmarshal(resource.Data, &coinInfo); err != nil {
		panic(err)
	}

	// need to decode, because init with hex
	name, _ := hex.DecodeString(coinInfo.Name)
	symbol, _ := hex.DecodeString(coinInfo.Symbol)
	fmt.Printf("name: %s, symbol: %s, decimals: %d, supply: %d\n", name, symbol, coinInfo.Decimals, coinInfo.Supply)
}
