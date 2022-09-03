package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Block(t *testing.T) {
	block, err := client.BlockByHeight(context.Background(), 8655089, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("block: \n")
	blockJson, _ := json.MarshalIndent(block, "", "    ")
	fmt.Printf(string(blockJson))
}
