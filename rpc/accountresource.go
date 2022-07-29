package rpc

import (
	"context"
	"encoding/json"
	"fmt"
)

type AccountCoin struct {
	Value uint64 `json:"value,string"`
}

type AccountData struct {
	Coin AccountCoin `json:"coin"`
}

type AccountResource struct {
	T    string      `json:"type"`
	Data AccountData `json:"data"`
}

type AccountResourcesResult []AccountResource

func (cl *Client) AccountResources(ctx context.Context, address string, version uint64) (*AccountResourcesResult, error) {
	result, err := cl.Get("/accounts/"+address+"/resources", map[string]string{
		"version": fmt.Sprintf("%d", version),
	})
	if err != nil {
		return nil, err
	}
	var accountResources AccountResourcesResult
	if err = json.Unmarshal(result, &accountResources); err != nil {
		return nil, err
	}
	return &accountResources, nil
}

func (cl *Client) AccountResourceByAddressAndType(ctx context.Context, address string, t string, version uint64) (*AccountResource, error) {
	result, err := cl.Get("/accounts/"+address+"/resource/"+t, map[string]string{
		"version": fmt.Sprintf("%d", version),
	})
	if err != nil {
		return nil, err
	}
	var accountResource AccountResource
	if err = json.Unmarshal(result, &accountResource); err != nil {
		return nil, err
	}
	return &accountResource, nil
}

func (cl *Client) AccountBalance(ctx context.Context, address string, coin string, version uint64) (uint64, error) {
	// how to get other coin balance
	// todo
	resouceType := fmt.Sprintf("0x1::coin::CoinStore<0x1::aptos_coin::AptosCoin>")
	result, err := cl.Get("/accounts/"+address+"/resource/"+resouceType, map[string]string{
		"version": fmt.Sprintf("%d", version),
	})
	if err != nil {
		return 0, err
	}
	var accountResource AccountResource
	if err = json.Unmarshal(result, &accountResource); err != nil {
		return 0, err
	}
	return accountResource.Data.Coin.Value, nil
}
