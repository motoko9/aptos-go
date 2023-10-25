package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet, err := NewFromKey("priviteKey")
	if err != nil {
		panic(err)
	}
	publickey := wallet.PublicKey().String()
	address := wallet.Address()
	fmt.Printf("public key = %s, address = %s\n", publickey, address)

}
