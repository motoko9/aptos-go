package rpc

import (
	"context"
)

type Account struct {
	SequenceNumber    uint64 `json:"sequence_number,string"`
	AuthenticationKey string `json:"authentication_key"`
}

func (cl *Client) Account(ctx context.Context, address string) (*Account, error) {
	var account Account
	code, err := cl.Get(ctx, "/accounts/"+address, nil, &account)
	if err != nil || code != 200 {
		return nil, err
	}
	return &account, nil
}
