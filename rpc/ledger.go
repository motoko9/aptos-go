package rpc

import (
	"context"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) Ledger(ctx context.Context) (*rpcmodule.Ledger, *rpcmodule.AptosError) {
	var ledger rpcmodule.Ledger
	err, aptosErr := cl.Get(ctx, "", nil, &ledger)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		return nil, aptosErr
	}
	return &ledger, nil
}
