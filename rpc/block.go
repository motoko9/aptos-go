package rpc

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) Block(ctx context.Context, height uint64, withTransactions bool) (*rpcmodule.Block, *rpcmodule.AptosError) {
	params := make(map[string]string)
	params["with_transactions"] = "false"
	if withTransactions {
		params["with_transactions"] = "true"
	}

	var block rpcmodule.Block
	err, aptosErr := cl.Get(ctx, "/blocks/by_height/"+fmt.Sprintf("%d", height), params, &block)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		return nil, aptosErr
	}
	return &block, nil
}
