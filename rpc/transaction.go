package rpc

import (
	"context"
	"encoding/json"
)

type Signature struct {
	T         string `json:"type"`
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
}

type Payload struct {
	T             string   `json:"type"`
	Function      string   `json:"function"`
	TypeArguments []string `json:"type_arguments"`
	Arguments     []string `json:"arguments"`
}

type TransactionResult struct {
	T                       string    `json:"type"`
	Hash                    string    `json:"hash"`
	Sender                  string    `json:"sender"`
	SequenceNumber          uint64    `json:"sequence_number,string"`
	MaxGasAmount            uint64    `json:"max_gas_amount,string"`
	GasUnitPrice            uint64    `json:"gas_unit_price,string"`
	GasCurrencyCode         string    `json:"gas_currency_code"`
	ExpirationTimestampSecs uint64    `json:"expiration_timestamp_secs,string"`
	Payload                 Payload   `json:"payload"`
	Signature               Signature `json:"signature"`
}

func (cl *Client) Transaction(ctx context.Context, hash string) (*TransactionResult, error) {
	result, err := cl.Get("/transactions/"+hash, nil)
	if err != nil {
		return nil, err
	}
	var transaction TransactionResult
	if err = json.Unmarshal(result, &transaction); err != nil {
		return nil, err
	}
	return &transaction, nil
}
