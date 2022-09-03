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
		return 0, &rpcmodule.AptosError{
			Message:     fmt.Sprintf("coin %s is not supported", coin),
			ErrorCode:   "400",
			VmErrorCode: 0,
		}
	}

	resourceType := fmt.Sprintf("0x1::coin::CoinStore<%s>", coin)
	resource, err := cl.AccountResourceByAddressAndType(ctx, address, resourceType, version)
	if err != nil {
		// resource not found, so balance is zero
		if err.ErrorCode == rpcmodule.ResourceNotFound {
			return 0, nil
		}
		return 0, err
	}
	coinStore, ok := resource.Object.(*aptosmodule.CoinStore)
	if !ok {
		return 0, &rpcmodule.AptosError{
			Message:     fmt.Sprintf("address %s resouce is invalid", address),
			ErrorCode:   "400",
			VmErrorCode: 0,
		}
	}
	return coinStore.Coin.Value, nil
}
