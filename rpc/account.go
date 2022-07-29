package rpc

import (
	"context"
	"encoding/json"
)

type AccountResult struct {
	SequenceNumber    uint64 `json:"sequence_number,string"`
	AuthenticationKey string `json:"authentication_key"`
}

func (cl *Client) Account(ctx context.Context, address string) (*AccountResult, error) {
	result, err := cl.Get("/accounts/"+address, nil)
	if err != nil {
		return nil, err
	}
	var account AccountResult
	if err = json.Unmarshal(result, &account); err != nil {
		return nil, err
	}
	return &account, nil
}
