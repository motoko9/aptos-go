package aptos

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/motoko9/aptos-go/aptosmodule"
)

func (cl *Client) AccountBalance(ctx context.Context, address string, coin string, version uint64) (uint64, error) {
	// how to get other coin balance
	// todo
	coin, ok := CoinType[coin]
	if !ok {
		return 0, fmt.Errorf("coin %s is not supported", coin)
	}
	resourceType := fmt.Sprintf("0x1::coin::CoinStore<%s>", coin)
	//
	accountResource, err := cl.AccountResourceByAddressAndType(ctx, address, resourceType, version)
	if err != nil {
		return 0, err
	}
	var coinStore aptosmodule.CoinStore
	if err = json.Unmarshal(accountResource.Data, &coinStore); err != nil {
		return 0, err
	}
	return coinStore.Coin.Value, nil
}
