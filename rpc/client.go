package rpc

import (
    "github.com/hashicorp/go-hclog"
    "github.com/motoko9/aptos-go/web/fetch"
)

type Client struct {
    fetchClient *fetch.Client
    logger      hclog.Logger
}

func NewClient(endpoint string, logger hclog.Logger) *Client {
    l := logger.ResetNamed("rpc-fetchClient")
    return &Client{
        fetchClient: fetch.NewClientWithEndpoint(endpoint, l),
        logger:      l,
    }
}
