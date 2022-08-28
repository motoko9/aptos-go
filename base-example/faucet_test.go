package base_example

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/faucet"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
	"time"
)

func TestFaucet(t *testing.T) {
	ctx := context.Background()

	// new account
	wallet, err := wallet.NewFromKeygenFile("account_example")
	if err != nil {
		panic(err)
	}
	address := wallet.Address()
	fmt.Printf("address: %s\n", address)

	// fund (max: 20000)
	amount := uint64(20000)
	hashes, aptosErr := faucet.FundAccount(address, amount)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("fund txs: %v\n", hashes)

	//
	time.Sleep(time.Second * 5)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// latest ledger
	ledger, aptosErr := client.Ledger(ctx)
	if aptosErr != nil {
		panic(aptosErr)
	}

	// check account
	balance, aptosErr := client.AccountBalance(ctx, address, aptos.AptosCoin, ledger.LedgerVersion)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("account balance: %d\n", balance)
}
