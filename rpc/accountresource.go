package rpc

import (
	"context"
	"encoding/json"
	"fmt"
)

type AccountResource struct {
	// only support CoinStore type
	// todo
	T    string    `json:"type"`
	//Data interface{} `json:"data"`
	Data json.RawMessage `json:"data"`
}

type AccountResources []AccountResource

func (cl *Client) AccountResources(ctx context.Context, address string, version uint64) (*AccountResources, error) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var accountResources AccountResources
	code, err := cl.Get(ctx, "/accounts/"+address+"/resources", params, &accountResources)
	if err != nil || code != 200 {
		return nil, err
	}
	return &accountResources, nil
}

func (cl *Client) AccountResourceByAddressAndType(ctx context.Context, address string, t string, version uint64) (*AccountResource, error) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var accountResource AccountResource
	code, err := cl.Get(ctx, "/accounts/"+address+"/resource/"+t, params, &accountResource)
	if err != nil || code != 200 {
		return nil, err
	}
	return &accountResource, nil
}
