package faucet

import (
	"fmt"
	"testing"
)

func TestClient_Account(t *testing.T) {
	hashes, err := FundAccount("0x697c173eeb917c95a382b60f546eb73a4c6a2a7b2d79e6c56c87104f9c04345f", 1000000)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", hashes)
}
