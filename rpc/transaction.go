package rpc

import (
	"context"
	"encoding/hex"
	"encoding/json"
)

type Signature struct {
	T         string `json:"type"`
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
}

type Module struct {
	ByteCode string `json:"bytecode"`
}

type Payload struct {
	T             string        `json:"type"`
	Function      string        `json:"function,omitempty"`
	TypeArguments []string      `json:"type_arguments"` //todo maybe need to omitempty, but move function call is needed event if empty
	Arguments     []interface{} `json:"arguments"`      //todo maybe need to omitempty, but move function call is needed event if empty
	Modules       []Module      `json:"modules,omitempty"`
}

type Transaction struct {
	T                       string     `json:"type,omitempty"`
	Hash                    string     `json:"hash,omitempty"`
	Sender                  string     `json:"sender"`
	SequenceNumber          uint64     `json:"sequence_number,string"`
	MaxGasAmount            uint64     `json:"max_gas_amount,string"`
	GasUnitPrice            uint64     `json:"gas_unit_price,string"`
	GasCurrencyCode         string     `json:"gas_currency_code,omitempty"`
	ExpirationTimestampSecs uint64     `json:"expiration_timestamp_secs,string"`
	Payload                 *Payload   `json:"payload"`
	Signature               *Signature `json:"signature,omitempty"`
}

func (cl *Client) Transaction(ctx context.Context, hash string) (*Transaction, error) {
	var transaction Transaction
	code, err := cl.Get(ctx, "/transactions/"+hash, nil, &transaction)
	if err != nil || code != 200 {
		return nil, err
	}
	return &transaction, nil
}

type SignMessageResult struct {
	Message string `json:"message"`
}

func (cl *Client) SignMessage(ctx context.Context, tx *Transaction) ([]byte, error) {
	requestBody, err := json.Marshal(tx)
	if err != nil {
		return nil, err
	}
	var signMessage SignMessageResult
	code, err := cl.Post(ctx, "/transactions/signing_message", nil, requestBody, &signMessage)
	if err != nil || code != 200 {
		return nil, err
	}
	//
	hexMessage := signMessage.Message[2:]
	message, err := hex.DecodeString(hexMessage)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (cl *Client) SubmitTransaction(ctx context.Context, tx *Transaction) (*Transaction, error) {
	requestBody, err := json.Marshal(tx)
	if err != nil {
		return nil, err
	}
	var transaction Transaction
	code, err := cl.Post(ctx, "/transactions", nil, requestBody, &transaction)
	if err != nil || (code != 200 && code != 202) {
		return nil, err
	}
	//
	return &transaction, nil
}
