package rpc

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) Transactions(ctx context.Context, start int64, limit int16) (*rpcmodule.Transactions, *rpcmodule.AptosError) {
	var params map[string]string
	if start > 0 && limit > 0 {
		params = make(map[string]string)
		params["start"] = fmt.Sprintf("%d", start)
		params["limit"] = fmt.Sprintf("%d", limit)
	}
	url := fmt.Sprintf("transactions")
	var transactions rpcmodule.Transactions
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Get(url).SetQueryParams(params).Execute(&transactions, &cl.rsp, &aptosError)
	if aptosError.IsError() {
		return nil, &aptosError
	}
	return &transactions, nil
}

func (cl *Client) TransactionByHash(ctx context.Context, hash string) (*rpcmodule.Transaction, *rpcmodule.AptosError) {
	url := fmt.Sprintf("/transactions/by_hash/%s", hash)
	var transaction rpcmodule.Transaction
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Get(url).Execute(&transaction, &cl.rsp, &aptosError)
	if aptosError.IsError() {
		return nil, &aptosError
	}
	return &transaction, nil
}

func (cl *Client) TransactionByVersion(ctx context.Context, version uint64) (*rpcmodule.Transaction, *rpcmodule.AptosError) {
	url := fmt.Sprintf("/transactions/by_version/%d", version)
	var transaction rpcmodule.Transaction
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Get(url).Execute(&transaction, &cl.rsp, &aptosError)
	if aptosError.IsError() {
		return nil, &aptosError
	}
	return &transaction, nil
}

func (cl *Client) TransactionsByAccount(ctx context.Context, address string, start int64, limit int16) (*rpcmodule.Transactions, *rpcmodule.AptosError) {
	var params map[string]string
	if start > 0 && limit > 0 {
		params = make(map[string]string)
		params["start"] = fmt.Sprintf("%d", start)
		params["limit"] = fmt.Sprintf("%d", limit)
	}
	url := fmt.Sprintf("accounts/%s/transactions", address)
	var transactions rpcmodule.Transactions
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Get(url).SetQueryParams(params).Execute(&transactions, &cl.rsp, &aptosError)
	if aptosError.IsError() {
		return nil, &aptosError
	}
	return &transactions, nil
}

func (cl *Client) EncodeSubmission(ctx context.Context, tx *rpcmodule.EncodeSubmissionRequest) ([]byte, *rpcmodule.AptosError) {
	url := fmt.Sprintf("/transactions/encode_submission")
	var raw string
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Post(url).SetJSONBody(tx).Execute(&raw, &cl.rsp, &aptosError)
	if aptosError.IsError() {
		return nil, &aptosError
	}
	hexMessage := raw[2:]
	message, err := hex.DecodeString(hexMessage)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	return message, nil
}

func (cl *Client) SubmitTransaction(ctx context.Context, tx *rpcmodule.SubmitTransactionRequest) (string, *rpcmodule.AptosError) {
	url := fmt.Sprintf("/transactions")
	var transaction rpcmodule.PendingTransactionRsp
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Post(url).SetJSONBody(tx).Execute(&transaction, &cl.rsp, &aptosError)
	if aptosError.IsError() {
		return "", &aptosError
	}
	return transaction.Hash, nil
}

func (cl *Client) SimulateTransaction(ctx context.Context, tx *rpcmodule.SubmitTransactionRequest) (rpcmodule.SimulateTransactionRsp, *rpcmodule.AptosError) {
	url := fmt.Sprintf("/transactions/simulate")
	var transaction rpcmodule.SimulateTransactionRsp
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Post(url).SetJSONBody(tx).Execute(&transaction, &cl.rsp, &aptosError)
	if aptosError.IsError() {
		return nil, &aptosError
	}
	return transaction, nil
}

func (cl *Client) EstimateGasPrice(ctx context.Context) (uint64, *rpcmodule.AptosError) {
	url := fmt.Sprintf("/estimate_gas_price")
	var gasEstimate rpcmodule.GasEstimate
	var aptosError rpcmodule.AptosError
	cl.fetchClient.Get(url).Execute(&gasEstimate, &cl.rsp, &aptosError)
	if aptosError.IsError() {
		return 0, &aptosError
	}
	return gasEstimate.GasEstimate, nil
}
