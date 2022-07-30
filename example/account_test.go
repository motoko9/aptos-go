package example

import (
	"fmt"
	"github.com/motoko9/aptos-go/utils"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
)

func TestAccount(t *testing.T) {
	// new account
	wallet := wallet.New()
	wallet.Save("account_1000")
	address := utils.Pubkey2Address(wallet.PublicKey())
	fmt.Printf("address: %s\n", address)
}
