package rpc

import (
    "context"
    "encoding/hex"
    "encoding/json"
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

	resp, err:= cl.fetchClient.Get("/transactions").SetQueryParams(params).Execute()
	if err != nil {
		return nil, err
	}

    var transactions rpcmodule.Transactions
    if err = json.Unmarshal(resp.BodyBytes(), &transactions); err != nil {
        return nil, err
    }

	return &transactions, nil
}

func (cl *Client) TransactionByHash(ctx context.Context, hash string) (*rpcmodule.Transaction, error) {
    resp, err := cl.fetchClient.Get(fmt.Sprintf("/transactions/by_hash/%v", hash)).Execute()
    if err != nil {
        return nil, err
    }

    var transaction rpcmodule.Transaction
    if err = json.Unmarshal(resp.BodyBytes(), &transaction); err != nil {
        return nil, err
    }
    return &transaction, nil
}

func (cl *Client) TransactionByVersion(ctx context.Context, version uint64) (*rpcmodule.Transaction, error) {
    resp, err := cl.fetchClient.Get(fmt.Sprintf("/transactions/by_version/%v", version)).Execute()
    if err != nil {
        return nil, err
    }

    var transaction rpcmodule.Transaction
    if err = json.Unmarshal(resp.BodyBytes(), &transaction); err != nil {
        return nil, err
    }
    return &transaction, nil
}

func (cl *Client) TransactionEncodeSubmission(ctx context.Context, tx *rpcmodule.EncodeSubmissionRequest) (string, error) {
    resp, err := cl.fetchClient.Post("/transactions/encode_submission").
        SetJSONBody(tx).Execute()
    if err != nil {
        return "", err
    }

    var raw string
    if err = json.Unmarshal(resp.BodyBytes(), &raw); err != nil {
        return "", err
    }
    return raw, nil
}

func (cl *Client) EncodeSubmission(ctx context.Context, tx *rpcmodule.EncodeSubmissionRequest) ([]byte, error) {
	resp, err:= cl.fetchClient.Post("/transactions/encode_submission").SetJSONBody(tx).Execute()
	if err != nil {
		return nil, err
	}

    var raw string
    if err = json.Unmarshal(resp.BodyBytes(), &raw); err != nil {
        return nil, err
    }

	hexMessage := raw[2:]
	message, err := hex.DecodeString(hexMessage)
	if err != nil {
		return nil, err
	}
	return message, nil
}


func (cl *Client) SubmitTransaction(ctx context.Context, tx *rpcmodule.SubmitTransactionRequest) (string, error) {
    resp, err := cl.fetchClient.Post("/transactions").SetJSONBody(tx).Execute()
    if err != nil {
        return "", err
    }

    var transaction rpcmodule.TransactionPendingTransaction
    if err = json.Unmarshal(resp.BodyBytes(), &transaction); err != nil {
        return "", err
    }
    return transaction.Hash, nil
}

func (cl *Client) EstimateGasPrice(ctx context.Context) (uint64, error) {
    resp, err:= cl.fetchClient.Get("/estimate_gas_price").Execute()
    if err != nil {
        return 0, err
    }

    var gasEstimate rpcmodule.GasEstimate
    if err = json.Unmarshal(resp.BodyBytes(), &gasEstimate); err != nil {
        return 0, err
    }

    return gasEstimate.GasEstimate, nil
}