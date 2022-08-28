package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_Block(t *testing.T) {
	client := New(DevNet_RPC)
	block, err := client.Block(context.Background(), 79591900, true)
	if err != nil {
		panic(err)
	}
	blockJson, _ := json.MarshalIndent(block, "", "    ")
	fmt.Printf("block: %s\n", string(blockJson))
}
