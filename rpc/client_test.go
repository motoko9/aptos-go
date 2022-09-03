package rpc

var client *Client

func init() {
	client = New(DevNet_RPC)
}
