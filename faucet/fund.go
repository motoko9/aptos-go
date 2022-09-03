package faucet

import (
	"fmt"
	"github.com/motoko9/aptos-go/fetchclient"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func FundAccount(address string, amount uint64) ([]string, *rpcmodule.AptosError) {
	fetchClient := fetchclient.NewClientWithEndpoint("https://faucet.devnet.aptoslabs.com")
	var hashes []string
	var aptosError rpcmodule.AptosError
	fetchClient.Post("/mint").
		SetQueryParams(map[string]string{
			"amount":  fmt.Sprintf("%d", amount),
			"address": address,
		}).Execute(&hashes, &aptosError)
	if aptosError.IsError() {
		return nil, &aptosError
	}
	return hashes, nil
}
