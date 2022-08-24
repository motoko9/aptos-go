package aptos

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

func (cl *Client) AccountBalance(ctx context.Context, address string, coin string, version uint64) (uint64, error) {
	// how to get other coin balance
	// todo
	coin, ok := CoinType[coin]
	if !ok {
		return 0, fmt.Errorf("coin %s is not supported", coin)
	}
	resourceType := fmt.Sprintf("0x1::coin::CoinStore<%s>", coin)
	//
	accountResource, err := cl.AccountResourceByAddressAndType(ctx, address, resourceType, version)
	if err != nil {
		return 0, err
	}
	var coinStore CoinStore
	if err = json.Unmarshal(accountResource.Data, &coinStore); err != nil {
		return 0, err
	}
	return coinStore.Coin.Value, nil
}
