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

func (cl *Client) SetHeaders(headers map[string]string) {
	cl.Client.SetHeaders(headers)
}


