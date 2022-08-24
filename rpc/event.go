package rpc

import (
	"context"
	"encoding/json"
)

type Event struct {
	Key            string `json:"key"`
	SequenceNumber uint64 `json:"sequence_number,string"`
	// only support TransferEvent type
	// todo
	T    string          `json:"type"`
	//Data interface{}   `json:"data"`
	Data json.RawMessage `json:"data"`
}

type Events []Event

func (cl *Client) EventsByKey(ctx context.Context, key string) (*Events, error) {
	var events Events
	code, err := cl.Get(ctx, "/events/"+key, nil, &events)
	if err != nil || code != 200 {
		return nil, err
	}
	return &events, nil
}

func (cl *Client) EventsByHandle(ctx context.Context, address string, handle string, field string) (*Events, error) {
	var events Events
	code, err := cl.Get(ctx, "/accounts/"+address+"/events/"+handle+"/"+field, nil, &events)
	if err != nil || code != 200 {
		return nil, err
	}
	return &events, nil
}
