package rpc

import (
	"context"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) EventsByKey(ctx context.Context, key string) (*rpcmodule.Events, *rpcmodule.AptosError) {
	var events rpcmodule.Events
	err, aptosErr := cl.Get(ctx, "/events/"+key, nil, &events)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		return nil, aptosErr
	}
	return &events, nil
}

func (cl *Client) EventsByHandle(ctx context.Context, address string, handle string, field string) (*rpcmodule.Events, *rpcmodule.AptosError) {
	var events rpcmodule.Events
	err, aptosErr := cl.Get(ctx, "/accounts/"+address+"/events/"+handle+"/"+field, nil, &events)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		return nil, aptosErr
	}
	return &events, nil
}
