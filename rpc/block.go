package rpc

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) BlockByHeight(ctx context.Context, height uint64, withTransactions bool) (*rpcmodule.Block, *rpcmodule.AptosError) {
	params := make(map[string]string)
	params["with_transactions"] = "false"
	if withTransactions {
		params["with_transactions"] = "true"
	}
	url := fmt.Sprintf("/blocks/by_height/%d", height)
	var block rpcmodule.Block
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Get(url).SetQueryParams(params).Execute(&block, &aptosError)
	if aptosError.IsError() {
		return nil, &aptosError
	}
	return &block, nil
}

func (cl *Client) BlockByVersion(ctx context.Context, version uint64, withTransactions bool) (*rpcmodule.Block, *rpcmodule.AptosError) {
	params := make(map[string]string)
	params["with_transactions"] = "false"
	if withTransactions {
		params["with_transactions"] = "true"
	}
	url := fmt.Sprintf("/blocks/by_version/%d", version)
	var block rpcmodule.Block
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Get(url).SetQueryParams(params).Execute(&block, &aptosError)
	if aptosError.IsError() {
		return nil, &aptosError
	}
	return &block, nil
}
