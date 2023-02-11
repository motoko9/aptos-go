package rpcmodule

type Block struct {
	BlockHeight    uint64        `json:"block_height,string"`
	BlockHash      string        `json:"block_hash"`
	BlockTimestamp uint64        `json:"block_timestamp,string"`
	FirstVersion   uint64        `json:"first_version,string"`
	LastVersion    uint64        `json:"last_version,string"`
	Transactions   []Transaction `json:"transactions"`
}

type Ledger struct {
	ChainId             uint64 `json:"chain_id"`
	Epoch               uint64 `json:"epoch,string"`
	LedgerVersion       uint64 `json:"ledger_version,string"`
	OldestLedgerVersion uint64 `json:"oldest_ledger_version,string"`
	BlockHeight         uint64 `json:"block_height,string"`
	OldestBlockHeight   uint64 `json:"oldest_block_height,string"`
	LedgerTimestamp     uint64 `json:"ledger_timestamp,string"`
	NodeRole            string `json:"node_role"`
	GitHash             string `json:"git_hash"`
}
