package rpc

import (
    "context"
    "encoding/json"
    "github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) Ledger(ctx context.Context) (*rpcmodule.Ledger, error) {
    resp, err := cl.fetchClient.Get("").Execute()
    if err != nil {
        return nil, err
    }
    var ledger rpcmodule.Ledger
    if err = json.Unmarshal(resp.BodyBytes(), &ledger); err != nil {
        return nil, err
    }

    return &ledger, nil
}
