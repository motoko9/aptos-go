package rpc

import (
	"context"
	"encoding/json"
	"fmt"
)

type Block struct {
	BlockHeight    uint64         `json:"block_height,string"`
	BlockHash      string         `json:"block_hash"`
	BlockTimestamp uint64         `json:"block_timestamp,string"`
	FirstVersion   uint64         `json:"first_version,string"`
	LastVersion    uint64         `json:"last_version,string"`
	Transactions   []*Transaction `json:"transactions"`
}

func (cl *Client) Block(ctx context.Context, height uint64, withTransactions bool) (*Block, error) {
	params := make(map[string]string)
	params["with_transactions"] = "false"
	if withTransactions {
		params["with_transactions"] = "true"
	}

	result, code, err := cl.Get("/blocks/by_height/"+fmt.Sprintf("%d", height), params)
	if err != nil || code != 200 {
		return nil, err
	}
	var block Block
	if err = json.Unmarshal(result, &block); err != nil {
		return nil, err
	}
	return &block, nil
}
