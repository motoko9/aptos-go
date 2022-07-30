package faucet

import (
	"encoding/json"
	"fmt"
	"github.com/motoko9/aptos-go/rpc"
)

type FundAccountResult struct {
}

func FundAccount(address string, amount uint64) ([]string, error) {
	client := rpc.New("https://faucet.devnet.aptoslabs.com")
	result, code, err := client.Post("/mint", map[string]string{
		"amount":  fmt.Sprintf("%d", amount),
		"address": address,
	}, nil)
	if err != nil || code != 200 {
		return nil, err
	}
	var hashes []string
	if err = json.Unmarshal(result, &hashes); err != nil {
		return nil, err
	}
	return hashes, nil
}
