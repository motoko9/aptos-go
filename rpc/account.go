package rpc

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) Account(ctx context.Context, address string, version uint64) (*rpcmodule.AccountData, *rpcmodule.AptosError) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var account rpcmodule.AccountData
	err, aptosErr := cl.Get(ctx, "/accounts/"+address, params, &account)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		return nil, aptosErr
	}
	return &account, nil
}

func (cl *Client) AccountResources(ctx context.Context, address string, version uint64) (*rpcmodule.MoveResources, *rpcmodule.AptosError) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var moveResources rpcmodule.MoveResources
	err, aptosErr := cl.Get(ctx, "/accounts/"+address+"/resources", params, &moveResources)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		return nil, aptosErr
	}
	return &moveResources, nil
}

func (cl *Client) AccountResourceByAddressAndType(ctx context.Context, address string, t string, version uint64) (*rpcmodule.MoveResource, *rpcmodule.AptosError) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var moveResource rpcmodule.MoveResource
	err, aptosErr := cl.Get(ctx, "/accounts/"+address+"/resource/"+t, params, &moveResource)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		return nil, aptosErr
	}
	return &moveResource, nil
}

func (cl *Client) AccountModules(ctx context.Context, address string, version uint64) (*rpcmodule.MoveModules, *rpcmodule.AptosError) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var moveModules rpcmodule.MoveModules
	err, aptosErr := cl.Get(ctx, "/accounts/"+address+"/modules", params, &moveModules)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		return nil, aptosErr
	}
	return &moveModules, nil
}

func (cl *Client) AccountModuleByAddressAndName(ctx context.Context, address string, name string, version uint64) (*rpcmodule.MoveModule, *rpcmodule.AptosError) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var moveModule rpcmodule.MoveModule
	err, aptosErr := cl.Get(ctx, "/accounts/"+address+"/rpcmodule/"+name, params, &moveModule)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		return nil, aptosErr
	}
	return &moveModule, nil
}
