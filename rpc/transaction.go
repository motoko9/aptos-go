package rpc

import (
	"context"
	"encoding/hex"
	"fmt"
)

type Signature struct {
	T         string `json:"type"`
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
}

type EntryFunctionPayload struct {
	T             string        `json:"type"`
	Function      string        `json:"function,omitempty"`
	TypeArguments []string      `json:"type_arguments"` //todo maybe need to omitempty, but move function call is needed event if empty
	Arguments     []interface{} `json:"arguments"`      //todo maybe need to omitempty, but move function call is needed event if empty
}

type Module struct {
	ByteCode string `json:"bytecode"`
}

type ModuleBundlePayload struct {
	T       string   `json:"type"`
	Modules []Module `json:"modules,omitempty"`
}

type Transaction struct {
	T                       string `json:"type,omitempty"`
	Hash                    string `json:"hash,omitempty"`
	Sender                  string `json:"sender"`
	SequenceNumber          uint64 `json:"sequence_number,string"`
	MaxGasAmount            uint64 `json:"max_gas_amount,string"`
	GasUnitPrice            uint64 `json:"gas_unit_price,string"`
	GasCurrencyCode         string `json:"gas_currency_code,omitempty"`
	ExpirationTimestampSecs uint64 `json:"expiration_timestamp_secs,string"`
	//SecondarySigners        []string    `json:"secondary_signers,omitempty"`
	Payload   interface{} `json:"payload"`
	Signature *Signature  `json:"signature,omitempty"`
}

type Transactions []Transaction

func (cl *Client) Transactions(ctx context.Context, start, limit int64) (*Transactions, error) {
	var params map[string]string
	if start > 0 && limit > 0 {
		params = make(map[string]string)
		params["start"] = fmt.Sprintf("%d", start)
		params["limit"] = fmt.Sprintf("%d", limit)
	}
	var transactions Transactions
	code, err := cl.Get(ctx, "/transactions", params, &transactions)
	if err != nil || code != 200 {
		return nil, err
	}
	return &transactions, nil
}

func (cl *Client) TransactionByHash(ctx context.Context, hash string) (*Transaction, error) {
	var transaction Transaction
	code, err := cl.Get(ctx, "/transactions/by_hash/"+hash, nil, &transaction)
	if err != nil || code != 200 {
		return nil, err
	}
	return &transaction, nil
}

func (cl *Client) TransactionByVersion(ctx context.Context, version uint64) (*Transaction, error) {
	var transaction Transaction
	code, err := cl.Get(ctx, "/transactions/by_version/"+fmt.Sprintf("%d", version), nil, &transaction)
	if err != nil || code != 200 {
		return nil, err
	}
	return &transaction, nil
}

type SignMessageResult struct {
	Message string `json:"message"`
}

func (cl *Client) EncodeSubmission(ctx context.Context, tx *Transaction) ([]byte, error) {
	var encodedSubmission string
	code, err := cl.Post(ctx, "/transactions/encode_submission", nil, tx, &encodedSubmission)
	if err != nil || code != 200 {
		return nil, err
	}
	//
	hexMessage := encodedSubmission[2:]
	message, err := hex.DecodeString(hexMessage)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (cl *Client) SubmitTransaction(ctx context.Context, tx *Transaction) (*Transaction, error) {
	var transaction Transaction
	code, err := cl.Post(ctx, "/transactions", nil, tx, &transaction)
	if err != nil || (code != 200 && code != 202) {
		return nil, err
	}
	//
	return &transaction, nil
}
