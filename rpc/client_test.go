package rpc

var client *Client

func init() {
	client = New(TestNet_RPC)
}
