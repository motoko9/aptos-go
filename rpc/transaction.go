package rpc

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Signature struct {
	T         string `json:"type"`
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
}

type Payload struct {
	T             string        `json:"type"`
	Function      string        `json:"function"`
	TypeArguments []interface{} `json:"type_arguments"`
	Arguments     []string      `json:"arguments"`
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
	result, code, err := cl.Get("/transactions/"+hash, nil)
	if err != nil || code != 200 {
		return nil, err
	}
	var transaction Transaction
	if err = json.Unmarshal(result, &transaction); err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (cl *Client) TransactionPending(ctx context.Context, hash string) (bool, error) {
	result, code, err := cl.Get("/transactions/"+hash, nil)
	if code == -1 {
		return false, err
	}
	if code == 404 {
		return true, nil
	}
	if code == 200 {
		var transaction Transaction
		if err = json.Unmarshal(result, &transaction); err != nil {
			return false, err
		}
		if transaction.T == "pending_transaction" {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, err
}

func (cl *Client) ConfirmTransaction(ctx context.Context, hash string) (bool, error) {
	counter := 0
	for counter < 100 {
		pending, err := cl.TransactionPending(ctx, hash)
		if err != nil {
			return false, err
		}
		if !pending {
			return true, nil
		}
		counter++
		time.Sleep(time.Second * 1)
	}
	return false, nil
}

type SignMessageResult struct {
	Message string `json:"message"`
}

func (cl *Client) SignMessage(ctx context.Context, tx *Transaction) ([]byte, error) {
	requestBody, err := json.Marshal(tx)
	if err != nil {
		return nil, err
	}
	result, code, err := cl.Post("/transactions/signing_message", nil, requestBody)
	if err != nil || code != 200 {
		return nil, err
	}
	var signMessage SignMessageResult
	if err = json.Unmarshal(result, &signMessage); err != nil {
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
	result, code, err := cl.Post("/transactions", nil, requestBody)
	if err != nil || (code != 200 && code != 202) {
		return nil, err
	}
	var transaction Transaction
	if err = json.Unmarshal(result, &transaction); err != nil {
		return nil, err
	}
	//
	return &transaction, nil
}
