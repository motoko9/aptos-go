package rpc

import (
	"context"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) Ledger(ctx context.Context) (*rpcmodule.Ledger, error) {
	var ledger rpcmodule.Ledger
	code, err := cl.Get(ctx, "", nil, &ledger)
	if err != nil || code != 200 {
		return nil, err
	}
	return &ledger, nil
}
