package rpc

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) Account(ctx context.Context, address string, version uint64) (*rpcmodule.AccountData, error) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var account rpcmodule.AccountData
	code, err := cl.Get(ctx, "/accounts/"+address, params, &account)
	if err != nil || code != 200 {
		return nil, err
	}
	return &account, nil
}

func (cl *Client) AccountResources(ctx context.Context, address string, version uint64) (*rpcmodule.MoveResources, error) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var moveResources rpcmodule.MoveResources
	code, err := cl.Get(ctx, "/accounts/"+address+"/resources", params, &moveResources)
	if err != nil || code != 200 {
		return nil, err
	}
	return &moveResources, nil
}

func (cl *Client) AccountResourceByAddressAndType(ctx context.Context, address string, t string, version uint64) (*rpcmodule.MoveResource, error) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var moveResource rpcmodule.MoveResource
	code, err := cl.Get(ctx, "/accounts/"+address+"/resource/"+t, params, &moveResource)
	if err != nil || code != 200 {
		return nil, err
	}
	return &moveResource, nil
}

func (cl *Client) AccountModules(ctx context.Context, address string, version uint64) (*rpcmodule.MoveModules, error) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var moveModules rpcmodule.MoveModules
	code, err := cl.Get(ctx, "/accounts/"+address+"/modules", params, &moveModules)
	if err != nil || code != 200 {
		return nil, err
	}
	return &moveModules, nil
}

func (cl *Client) AccountModuleByAddressAndName(ctx context.Context, address string, name string, version uint64) (*rpcmodule.MoveModule, error) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var moveModule rpcmodule.MoveModule
	code, err := cl.Get(ctx, "/accounts/"+address+"/rpcmodule/"+name, params, &moveModule)
	if err != nil || code != 200 {
		return nil, err
	}
	return &moveModule, nil
}
