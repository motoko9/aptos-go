package rpc

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) BlockByHeight(ctx context.Context, height uint64, withTransactions bool) (*rpcmodule.Block, error) {
    params := make(map[string]string)
    params["with_transactions"] = "false"
    if withTransactions {
        params["with_transactions"] = "true"
    }

    resp, err := cl.fetchClient.Get("/blocks/by_height/" + fmt.Sprintf("%d", height)).SetQueryParams(params).Execute()
    if err != nil {
        return nil, err
    }
    var block rpcmodule.Block
    if err = json.Unmarshal(resp.BodyBytes(), &block); err != nil {
        return nil, err
    }
    return &block, nil
}

func (cl *Client) BlockByVersion(ctx context.Context, version uint64, withTransactions bool) (*rpcmodule.Block, error) {
    params := make(map[string]string)
    params["with_transactions"] = "false"
    if withTransactions {
        params["with_transactions"] = "true"
    }

    resp, err := cl.fetchClient.Get("/blocks/by_version/" + fmt.Sprintf("%d", version)).SetQueryParams(params).Execute()
    if err != nil {
        return nil, err
    }
    var block rpcmodule.Block
    if err = json.Unmarshal(resp.BodyBytes(), &block); err != nil {
        return nil, err
    }
    return &block, nil
}
