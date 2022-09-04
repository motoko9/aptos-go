package hello_blockchain

import (
	"github.com/motoko9/aptos-go/aptosmodule"
	"github.com/motoko9/aptos-go/rpcmodule"
)

type MessageHolder struct {
	Message             string              `json:"message"`
	MessageChangeEvents aptosmodule.Events `json:"message_change_events"`
}

func MessageHolderCreator() interface{} {
	return &MessageHolder{}
}

func init() {
	rpcmodule.RegisterResourceObjectCreator("0xbb04c2079bc5611345689582eabab626732411b909045f8326d2b4980eac9b07::message::MessageHolder", MessageHolderCreator)
}
