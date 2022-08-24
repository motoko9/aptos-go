package faucet

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/rpc"
)

type FundAccountResult struct {
}

func FundAccount(address string, amount uint64) ([]string, error) {
	client := rpc.New("https://faucet.devnet.aptoslabs.com")
	var hashes []string
	code, err := client.Post(context.Background(), "/mint", map[string]string{
		"amount":  fmt.Sprintf("%d", amount),
		"address": address,
	}, nil, &hashes)
	if err != nil || code != 200 {
		return nil, err
	}
	return hashes, nil
}
