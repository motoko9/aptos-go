package rpc

import (
	"context"
	"encoding/json"
)

type Account struct {
	SequenceNumber    uint64 `json:"sequence_number,string"`
	AuthenticationKey string `json:"authentication_key"`
}

func (cl *Client) Account(ctx context.Context, address string) (*Account, error) {
	result, code, err := cl.Get("/accounts/"+address, nil)
	if err != nil || code != 200 {
		return nil, err
	}
	var account Account
	if err = json.Unmarshal(result, &account); err != nil {
		return nil, err
	}
	return &account, nil
}
