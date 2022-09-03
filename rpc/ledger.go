package rpc

import (
    "context"
    "github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) Ledger(ctx context.Context) (*rpcmodule.Ledger, *rpcmodule.AptosError) {
    url := ""
    var ledger rpcmodule.Ledger
    var aptosError rpcmodule.AptosError
    cl.fetchClient.Get(url).Execute(&ledger, &aptosError)
    if aptosError.IsError() {
        return nil, &aptosError
    }
    return &ledger, nil
}
