package faucet

import (
	"fmt"
	"testing"
)

func TestClient_Account(t *testing.T) {
	hashs, err := FundAccount("0xb0e4f55ea10ba1604f943ceb665c698ebc72c19fef44173572ab11107c1d6b30", 1000000)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", hashs)
}
