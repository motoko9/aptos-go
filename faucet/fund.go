package faucet

import (
    "encoding/json"
    "fmt"
    "github.com/hashicorp/go-hclog"
	"github.com/motoko9/aptos-go/fetchclient"
)

func FundAccount(address string, amount uint64) ([]string, error) {
    fetchClient := fetchclient.NewClientWithEndpoint("https://faucet.devnet.aptoslabs.com", hclog.Default())
    resp, err := fetchClient.Post("/mint").
        SetQueryParams(map[string]string{
            "amount":  fmt.Sprintf("%d", amount),
            "address": address,
        }).Execute()
    if err != nil {
        return nil, err
    }
    var hashes []string
    if err = json.Unmarshal(resp.BodyBytes(), &hashes); err != nil {
        return nil, err
    }
    return hashes, nil
}
