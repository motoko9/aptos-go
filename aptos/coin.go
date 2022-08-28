package aptos

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/motoko9/aptos-go/aptosmodule"
	"github.com/motoko9/aptos-go/rpcmodule"
	"strings"
)

const (
	AptosCoin = "Aptos"
	BTCCoin   = "BTC"
	USDTCoin  = "USDT"
)

// only for devnet, mainnet is diffierent
// todo
var CoinType = map[string]string{
	"Aptos": "0x1::aptos_coin::AptosCoin",
	"BTC":   "0x43417434fd869edee76cca2a4d2301e528a1551b1d719b75c350c3c97d15b8b9::coins::BTC",
	"USDT":  "0x1685cdc9a83c3da34c59208f34bddb3217f63bfbe0c393f04462d1ba06465d08::usdt::USDT",
}

func AddressFromCoinType(coinType string) string {
	items := strings.Split(coinType, "::")
	if len(items) != 3 {
		return ""
	}
	return items[0]
}

func (cl *Client) CoinInfo(ctx context.Context, coin string, version uint64) (*aptosmodule.CoinInfo, *rpcmodule.AptosError) {
	coinType, ok := CoinType[coin]
	if !ok {
		return nil, rpcmodule.AptosErrorFromError(fmt.Errorf("coin %s is not supported", coin))
	}
	//
	coinAddress := AddressFromCoinType(coinType)
	coinInfoResourceType := fmt.Sprintf("0x1::coin::CoinInfo<%s>", coinType)
	coinInfoResource, aptosErr := cl.AccountResourceByAddressAndType(ctx, coinAddress, coinInfoResourceType, version)
	if aptosErr != nil {
		return nil, aptosErr
	}
	//
	var coinInfo aptosmodule.CoinInfo
	if err := json.Unmarshal(coinInfoResource.Data, &coinInfo); err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	return &coinInfo, nil
}
