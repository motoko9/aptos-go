package rpc

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/common/jsonutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Block(t *testing.T) {
	block, err := client.BlockByHeight(context.Background(), 79591900, true)
	assert.NoError(t, err)
	fmt.Printf("block: \n")
	jsonutil.PrintJsonStringWithIndent(block)
}
