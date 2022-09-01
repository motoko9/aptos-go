package rpc

import "github.com/hashicorp/go-hclog"

var client *Client

func init() {
    client = NewClient(DevNet_RPC, hclog.Default())
}
