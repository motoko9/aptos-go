package aptos

import (
	"github.com/hashicorp/go-hclog"
	"github.com/motoko9/aptos-go/rpc"
)

type Client struct {
	*rpc.Client
}

func NewClient(endpoint string, logger hclog.Logger) *Client {
	client := rpc.NewClient(endpoint, logger)
	return &Client{
		client,
	}
}
