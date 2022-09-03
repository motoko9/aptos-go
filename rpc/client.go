package rpc

import (
    "github.com/hashicorp/go-hclog"
	"github.com/motoko9/aptos-go/fetchclient"
)

type Client struct {
    fetchClient *fetchclient.Client
}

func New(endpoint string) *Client {
    l := hclog.Default().Named("rpc-client")
    return &Client{
        fetchClient: fetchclient.NewClientWithEndpoint(endpoint, l),
    }
}
