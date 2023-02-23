package rpc

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/rpcmodule"
	"testing"
)

func Test_View(t *testing.T) {
	viewReq := &rpcmodule.ViewRequest{
		Function:      "0x1::coin::balance",
		TypeArguments: []string{"0x1::aptos_coin::AptosCoin"},
		Arguments:     []interface{}{"0x74f3bbe39c7e2793a2e5445ee0336c9ac3191534762b41dcfc1054ad077ccc7c"},
	}
	r, err := client.View(context.Background(), viewReq)
	if err != nil {
		panic(err)
	}
	fmt.Printf("view result: %s\n", r)
}
