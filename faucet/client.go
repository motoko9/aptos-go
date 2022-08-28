package faucet

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/rpcmodule"
)

type FundAccountResult struct {
}

func FundAccount(address string, amount uint64) ([]string, *rpcmodule.AptosError) {
	client := rpc.New("https://faucet.devnet.aptoslabs.com")
	var hashes []string
	err, aptosErr := client.Post(context.Background(), "/mint", map[string]string{
		"amount":  fmt.Sprintf("%d", amount),
		"address": address,
	}, nil, &hashes)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		return nil, aptosErr
	}
	return hashes, nil
}
