package aptos

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/rpc"
	"time"
)

func (cl *Client) TransactionPending(ctx context.Context, hash string) (bool, error) {
	var transaction rpc.Transaction
	code, err := cl.Get(ctx, "/transactions/"+hash, nil, &transaction)
	if code == -1 {
		return false, err
	}
	if code == 404 {
		// resource not found, maybe transaction is not on chain
		return true, nil
	}
	if code == 200 {
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

func (cl *Client) PublishMoveModule(ctx context.Context, account string, sequenceNumber uint64, module []byte) (*rpc.Transaction, error) {
	publishPayload := rpc.Payload{
		T: "module_bundle_payload",
		Modules: []rpc.Module{
			{
				ByteCode: "0x" + hex.EncodeToString(module),
			},
		},
	}
	publish := rpc.Transaction{
		T:                       "",
		Hash:                    "",
		Sender:                  account,
		SequenceNumber:          sequenceNumber,
		MaxGasAmount:            uint64(2000),
		GasUnitPrice:            uint64(1),
		GasCurrencyCode:         "",
		ExpirationTimestampSecs: uint64(time.Now().Unix() + 600), // now + 10 minutes
		Payload:                 &publishPayload,
		Signature:               nil,
	}
	return &publish, nil
}

func (cl *Client) TransferCoin(ctx context.Context, from string, sequenceNumber uint64, coin string, amount uint64, receipt string) (*rpc.Transaction, error) {
	// transfer
	coin, ok := CoinType[coin]
	if !ok {
		return nil, fmt.Errorf("coin %s is not supported", coin)
	}
	transferPayload := rpc.Payload{
		Function:      "0x1::coin::transfer",
		Arguments:     []interface{}{receipt, fmt.Sprintf("%d", amount)},
		T:             "script_function_payload",
		TypeArguments: []string{coin},
	}
	transaction := rpc.Transaction{
		T:                       "",
		Hash:                    "",
		Sender:                  from,
		SequenceNumber:          sequenceNumber,
		MaxGasAmount:            uint64(2000),
		GasUnitPrice:            uint64(1),
		GasCurrencyCode:         "",
		ExpirationTimestampSecs: uint64(time.Now().Unix() + 600), // now + 10 minutes
		Payload:                 &transferPayload,
		Signature:               nil,
	}
	return &transaction, nil
}

func (cl *Client) RegisterRecipient(ctx context.Context, from string, sequenceNumber uint64, coin string) (*rpc.Transaction, error) {
	// transfer
	coin, ok := CoinType[coin]
	if !ok {
		return nil, fmt.Errorf("coin %s is not supported", coin)
	}
	transferPayload := rpc.Payload{
		Function:      "0x1::coins::register",
		Arguments:     []interface{}{},
		T:             "script_function_payload",
		TypeArguments: []string{coin},
	}
	transaction := rpc.Transaction{
		T:                       "",
		Hash:                    "",
		Sender:                  from,
		SequenceNumber:          sequenceNumber,
		MaxGasAmount:            uint64(2000),
		GasUnitPrice:            uint64(1),
		GasCurrencyCode:         "",
		ExpirationTimestampSecs: uint64(time.Now().Unix() + 600), // now + 10 minutes
		Payload:                 &transferPayload,
		Signature:               nil,
	}
	return &transaction, nil
}
