package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func Test_BlockByHeight(t *testing.T) {
	block, err := client.BlockByHeight(context.Background(), 57770440, true)
	if err != nil {
		panic(err)
	}
	blockJson, _ := json.MarshalIndent(block, "", "    ")
	fmt.Printf(string(blockJson))
}

func Test_BlockByVersion(t *testing.T) {
	block, err := client.BlockByVersion(context.Background(), 425348794, true)
	if err != nil {
		panic(err)
	}
	blockJson, _ := json.MarshalIndent(block, "", "    ")
	fmt.Printf(string(blockJson))
}
