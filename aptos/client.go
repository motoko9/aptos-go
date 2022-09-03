package aptos

import (
	"github.com/hashicorp/go-hclog"
	"github.com/motoko9/aptos-go/rpc"
)

type Client struct {
	*rpc.Client
}

func New(endpoint string, logger hclog.Logger) *Client {
	client := rpc.New(endpoint, logger)
	return &Client{
		client,
	}
}
