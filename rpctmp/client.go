package rpctmp

import (
	"github.com/motoko9/aptos-go/fetchclient"
)

type Client struct {
	fetchClient *fetchclient.Client
}

func New(endpoint string) *Client {
	return &Client{
		fetchClient: fetchclient.NewClientWithEndpoint(endpoint),
	}
}
