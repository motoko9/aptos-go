package rpc

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) EventsByKey(ctx context.Context, key string) (*rpcmodule.Events, *rpcmodule.AptosError) {
	url := fmt.Sprintf("/events/%s", key)
	var events rpcmodule.Events
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Get(url).Execute(&events, &aptosError)
	if aptosError.IsError() {
		return nil, &aptosError
	}
	return &events, nil
}

func (cl *Client) EventsByHandle(ctx context.Context, address string, handle string, field string) (*rpcmodule.Events, *rpcmodule.AptosError) {
	url := fmt.Sprintf("/accounts/%s/events/%s/%s", address, handle, field)
	var events rpcmodule.Events
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Get(url).Execute(&events, &aptosError)
	if aptosError.IsError() {
		return nil, &aptosError
	}
	return &events, nil
}
