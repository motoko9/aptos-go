package rpc

import (
	"context"
	"encoding/json"
)

type Ledger struct {
	ChainId             uint64 `json:"chain_id"`
	Epoch               uint64 `json:"epoch"`
	LedgerVersion       uint64 `json:"ledger_version,string"`
	OldestLedgerVersion string `json:"oldest_ledger_version"`
	LedgerTimestamp     uint64 `json:"ledger_timestamp,string"`
	NodeRole            string `json:"node_role"`
}

func (cl *Client) Ledger(ctx context.Context) (*Ledger, error) {
	result, code, err := cl.Get("", nil)
	if err != nil || code != 200 {
		return nil, err
	}
	var ledger Ledger
	if err = json.Unmarshal(result, &ledger); err != nil {
		return nil, err
	}
	return &ledger, nil
}
