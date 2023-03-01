package rpc

import (
	"github.com/motoko9/aptos-go/fetchclient"
	"github.com/motoko9/aptos-go/rpcmodule"
)

type Client struct {
	fetchClient *fetchclient.Client
	rsp rpcmodule.General
}

func New(endpoint string) *Client {
	return &Client{
		fetchClient: fetchclient.NewClientWithEndpoint(endpoint),
	}
}
