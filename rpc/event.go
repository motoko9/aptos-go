package rpc

import (
	"context"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) EventsByKey(ctx context.Context, key string) (*rpcmodule.Events, error) {
	var events rpcmodule.Events
	code, err := cl.Get(ctx, "/events/"+key, nil, &events)
	if err != nil || code != 200 {
		return nil, err
	}
	return &events, nil
}

func (cl *Client) EventsByHandle(ctx context.Context, address string, handle string, field string) (*rpcmodule.Events, error) {
	var events rpcmodule.Events
	code, err := cl.Get(ctx, "/accounts/"+address+"/events/"+handle+"/"+field, nil, &events)
	if err != nil || code != 200 {
		return nil, err
	}
	return &events, nil
}
