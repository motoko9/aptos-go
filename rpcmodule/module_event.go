package rpcmodule

import (
	"encoding/json"
)

type Events []Event

type Event struct {
	Key            string          `json:"key"`
	SequenceNumber uint64          `json:"sequence_number,string"`
	Type           string          `json:"type"`
	Data           json.RawMessage `json:"data"`
}
