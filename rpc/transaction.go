package rpc

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) Transactions(ctx context.Context, start, limit int64) (*rpcmodule.Transactions, error) {
	var params map[string]string
	if start > 0 && limit > 0 {
		params = make(map[string]string)
		params["start"] = fmt.Sprintf("%d", start)
		params["limit"] = fmt.Sprintf("%d", limit)
	}
	var transactions rpcmodule.Transactions
	code, err := cl.Get(ctx, "/transactions", params, &transactions)
	if err != nil || code != 200 {
		return nil, err
	}
	return &transactions, nil
}

func (cl *Client) TransactionByHash(ctx context.Context, hash string) (*rpcmodule.Transaction, error) {
	var transaction rpcmodule.Transaction
	code, err := cl.Get(ctx, "/transactions/by_hash/"+hash, nil, &transaction)
	if err != nil || code != 200 {
		return nil, err
	}
	return &transaction, nil
}

func (cl *Client) TransactionByVersion(ctx context.Context, version uint64) (*rpcmodule.Transaction, error) {
	var transaction rpcmodule.Transaction
	code, err := cl.Get(ctx, "/transactions/by_version/"+fmt.Sprintf("%d", version), nil, &transaction)
	if err != nil || code != 200 {
		return nil, err
	}
	return &transaction, nil
}

func (cl *Client) EncodeSubmission(ctx context.Context, tx *rpcmodule.EncodeSubmissionRequest) ([]byte, error) {
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

func (cl *Client) SubmitTransaction(ctx context.Context, tx *rpcmodule.SubmitTransactionRequest) (string, error) {
	var transaction rpcmodule.TransactionPendingTransaction
	code, err := cl.Post(ctx, "/transactions", nil, tx, &transaction)
	if err != nil || (code != 200 && code != 202) {
		return "", err
	}
	return transaction.Hash, nil
}

func (cl *Client) EstimateGasPrice(ctx context.Context) (uint64, error) {
	var gasEstimate rpcmodule.GasEstimate
	code, err := cl.Get(ctx, "/estimate_gas_price", nil, &gasEstimate)
	if err != nil || code != 200 {
		return 0, err
	}
	return gasEstimate.GasEstimate, nil
}
