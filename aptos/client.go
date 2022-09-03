package aptos

import (
	"github.com/motoko9/aptos-go/rpc"
)

type Client struct {
	*rpc.Client
}

func New(endpoint string) *Client {
	client := rpc.New(endpoint)
	return &Client{
		client,
	}
}
