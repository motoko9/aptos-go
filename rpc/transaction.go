package rpc

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) Transactions(ctx context.Context, start, limit int64) (*rpcmodule.Transactions, *rpcmodule.AptosError) {
	var params map[string]string
	if start > 0 && limit > 0 {
		params = make(map[string]string)
		params["start"] = fmt.Sprintf("%d", start)
		params["limit"] = fmt.Sprintf("%d", limit)
	}
	var transactions rpcmodule.Transactions
	err, aptosErr := cl.Get(ctx, "/transactions", params, &transactions)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		return nil, aptosErr
	}
	return &transactions, nil
}

func (cl *Client) TransactionByHash(ctx context.Context, hash string) (*rpcmodule.Transaction, *rpcmodule.AptosError) {
	var transaction rpcmodule.Transaction
	err, aptosErr := cl.Get(ctx, "/transactions/by_hash/"+hash, nil, &transaction)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		return nil, aptosErr
	}
	return &transaction, nil
}

func (cl *Client) TransactionByVersion(ctx context.Context, version uint64) (*rpcmodule.Transaction, *rpcmodule.AptosError) {
	var transaction rpcmodule.Transaction
	err, aptosErr := cl.Get(ctx, "/transactions/by_version/"+fmt.Sprintf("%d", version), nil, &transaction)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		return nil, aptosErr
	}
	return &transaction, nil
}

func (cl *Client) EncodeSubmission(ctx context.Context, tx *rpcmodule.EncodeSubmissionRequest) ([]byte, *rpcmodule.AptosError) {
	var encodedSubmission string
	err, aptosErr := cl.Post(ctx, "/transactions/encode_submission", nil, tx, &encodedSubmission)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		return nil, aptosErr
	}
	//
	hexMessage := encodedSubmission[2:]
	message, err := hex.DecodeString(hexMessage)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	return message, nil
}

func (cl *Client) SubmitTransaction(ctx context.Context, tx *rpcmodule.SubmitTransactionRequest) (string, *rpcmodule.AptosError) {
	var transaction rpcmodule.TransactionPendingTransaction
	err, aptosErr := cl.Post(ctx, "/transactions", nil, tx, &transaction)
	if err != nil {
		return "", rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		return "", aptosErr
	}
	return transaction.Hash, nil
}

func (cl *Client) EstimateGasPrice(ctx context.Context) (uint64, *rpcmodule.AptosError) {
	var gasEstimate rpcmodule.GasEstimate
	err, aptosErr := cl.Get(ctx, "/estimate_gas_price", nil, &gasEstimate)
	if err != nil {
		return 0, rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		return 0, aptosErr
	}
	return gasEstimate.GasEstimate, nil
}
