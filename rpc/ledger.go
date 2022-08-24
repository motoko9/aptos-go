package rpc

import (
	"context"
)

type Ledger struct {
	ChainId             uint64 `json:"chain_id"`
	Epoch               uint64 `json:"epoch,string"`
	LedgerVersion       uint64 `json:"ledger_version,string"`
	OldestLedgerVersion uint64 `json:"oldest_ledger_version,string"`
	BlockHeight         uint64 `json:"block_height,string"`
	OldestBlockHeight   uint64 `json:"oldest_block_height,string"`
	LedgerTimestamp     uint64 `json:"ledger_timestamp,string"`
	NodeRole            string `json:"node_role"`
}

func (cl *Client) Ledger(ctx context.Context) (*Ledger, error) {
	var ledger Ledger
	code, err := cl.Get(ctx, "", nil, &ledger)
	if err != nil || code != 200 {
		return nil, err
	}
	return &ledger, nil
}
