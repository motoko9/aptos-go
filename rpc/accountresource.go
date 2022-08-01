package rpc

import (
	"context"
	"encoding/json"
	"fmt"
)

type Guid struct {
	Id struct {
		Addr        string `json:"addr"`
		CreationNum uint64 `json:"creation_num,string"`
	} `json:"id"`
}
type Coin struct {
	Value uint64 `json:"value,string"`
}

type CoinEvents struct {
	Counter uint64 `json:"counter,string"`
	Guid    Guid   `json:"guid"`
}

type CoinStore struct {
	Coin           Coin       `json:"coin"`
	DepositEvents  CoinEvents `json:"deposit_events"`
	WithdrawEvents CoinEvents `json:"withdraw_events"`
}

type AccountResource struct {
	// only support CoinStore type
	// todo
	T    string    `json:"type"`
	Data CoinStore `json:"data"`
	//Data json.RawMessage `json:"data"`
}

type AccountResources []AccountResource

func (cl *Client) AccountResources(ctx context.Context, address string, version uint64) (*AccountResources, error) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	result, code, err := cl.Get("/accounts/"+address+"/resources", params)
	if err != nil || code != 200 {
		return nil, err
	}
	var accountResources AccountResources
	if err = json.Unmarshal(result, &accountResources); err != nil {
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
	result, code, err := cl.Get("/accounts/"+address+"/resource/"+t, params)
	if err != nil || code != 200 {
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
	coin, ok := CoinType[coin]
	if !ok {
		return 0, fmt.Errorf("coin %s is not supported", coin)
	}
	resourceType := fmt.Sprintf("0x1::coin::CoinStore<%s>", coin)
	//
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	result, code, err := cl.Get("/accounts/"+address+"/resource/"+resourceType, params)
	if err != nil || code != 200 {
		return 0, err
	}
	var accountResource AccountResource
	if err = json.Unmarshal(result, &accountResource); err != nil {
		return 0, err
	}
	return accountResource.Data.Coin.Value, nil
}
