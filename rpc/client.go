package rpc

import (
    "github.com/hashicorp/go-hclog"
    "github.com/motoko9/aptos-go/httpclient"
)

type Client struct {
    fetchClient *httpclient.Client
}

func New(endpoint string) *Client {
    l := hclog.Default().Named("rpc-client")
    return &Client{
        fetchClient: httpclient.NewClientWithEndpoint(endpoint, l),
    }
}
