package hello_blockchain

import (
	"github.com/motoko9/aptos-go/rpcmodule"
)

type MessageChangeEvents struct {
	Counter string `json:"counter"`
	Guid    struct {
		Id struct {
			Addr        string `json:"addr"`
			CreationNum string `json:"creation_num"`
		} `json:"id"`
	} `json:"guid"`
}

type MessageHolder struct {
	Message             string              `json:"message"`
	MessageChangeEvents MessageChangeEvents `json:"message_change_events"`
}

func MessageHolderCreator() interface{} {
	return &MessageHolder{}
}

func init() {
	rpcmodule.RegisterResourceObjectCreator("0x1::account::Account", MessageHolderCreator)
}
