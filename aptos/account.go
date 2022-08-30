package aptos

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptosmodule"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) AccountBalance(ctx context.Context, address string, coin string, version uint64) (uint64, *rpcmodule.AptosError) {
	// how to get other coin balance
	// todo
	coin, ok := CoinType[coin]
	if !ok {
		return 0, rpcmodule.AptosErrorFromError(fmt.Errorf("coin %s is not supported", coin))
	}
	//
	resourceType := fmt.Sprintf("0x1::coin::CoinStore<%s>", coin)
	accountResource, aptosErr := cl.AccountResourceByAddressAndType(ctx, address, resourceType, version)
	if aptosErr != nil {
		// resource not found, so balance is zero
		if aptosErr.ErrorCode == rpcmodule.ResourceNotFound {
			return 0, nil
		}
		return 0, aptosErr
	}
	coinStore := accountResource.Object.(*aptosmodule.CoinStore)
	return coinStore.Coin.Value, nil
}
