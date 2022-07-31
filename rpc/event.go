package rpc

import (
	"context"
	"encoding/json"
)

type TransferEvent struct {
	Amount uint64 `json:"amount,string"`
}

type Event struct {
	Key            string `json:"key"`
	SequenceNumber uint64 `json:"sequence_number,string"`
	// only support TransferEvent type
	// todo
	T    string          `json:"type"`
	Data TransferEvent   `json:"data"`
	//Data json.RawMessage `json:"data"`
}

type Events []Event

func (cl *Client) EventsByKey(ctx context.Context, key string) (*Events, error) {
	result, code, err := cl.Get("/events/"+key, nil)
	if err != nil || code != 200 {
		return nil, err
	}
	var events Events
	if err = json.Unmarshal(result, &events); err != nil {
		return nil, err
	}
	return &events, nil
}

func (cl *Client) EventsByHandle(ctx context.Context, address string, handle string, field string) (*Events, error) {
	result, code, err := cl.Get("/accounts/"+address+"/events/"+handle+"/"+field, nil)
	if err != nil || code != 200 {
		return nil, err
	}
	var events Events
	if err = json.Unmarshal(result, &events); err != nil {
		return nil, err
	}
	return &events, nil
}
