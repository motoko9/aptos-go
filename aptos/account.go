package aptos

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptosmodule"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) AccountBalance(ctx context.Context, address string, coin string, version uint64) (uint64, error) {
	// how to get other coin balance
	// todo
	coin, ok := CoinType[coin]
	if !ok {
		return 0, rpcmodule.ClientErrorCtor(400, fmt.Sprintf("coin %s is not supported", coin))
	}

	var coinStore aptosmodule.CoinStore
	resourceType := fmt.Sprintf("0x1::coin::CoinStore<%s>", coin)
	err := cl.AccountResourceByAddressAndType(ctx, address, resourceType, version, &coinStore)
	if err != nil {
		// resource not found, so balance is zero
		if err.Error() == rpcmodule.ResourceNotFound {
			return 0, nil
		}
		return 0, err
	}
	return coinStore.Coin.Value, nil
}
