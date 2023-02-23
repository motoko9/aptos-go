package base_example

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
	"time"
)

func TestNewExampleAccount(t *testing.T) {
	// new account
	exampleWallet := wallet.New()
	exampleWallet.SaveToKeygenFile("account_example")
	exampleAddress := exampleWallet.Address()
	fmt.Printf("example address: %s\n", exampleAddress)

	// new rpc
	client := aptos.New(rpc.TestNet_RPC)

	//
	faultWallet, err := wallet.NewFromKeygenFile("./../account_fault")
	if err != nil {
		panic(err)
	}
	faultAddress := faultWallet.Address()
	fmt.Printf("fault address: %s\n", faultAddress)
	fmt.Printf("fault public key: %s\n", faultWallet.PublicKey().String())
	fmt.Printf("fault private key: %s\n", faultWallet.PrivateKey.String())

	// create account on aptos
	txHash, aptosErr := client.CreateAccount(context.Background(), faultAddress, exampleAddress, faultWallet)
	if aptosErr != nil {
		panic(err)
	}
	fmt.Printf("create example account transacion: %s\n", txHash)

	//
	time.Sleep(time.Second * 5)

	// fund
	txHash, aptosErr = client.TransferCoin(context.Background(), faultAddress, aptos.AptosCoin, 100000000, exampleAddress, faultWallet)
	if aptosErr != nil {
		panic(err)
	}
	fmt.Printf("fund example account transacion: %s\n", txHash)

	//
	time.Sleep(time.Second * 5)

	// latest ledger
	ledger, aptosErr := client.Ledger(context.Background())
	if aptosErr != nil {
		panic(aptosErr)
	}

	// check account
	balance, aptosErr := client.AccountBalance(context.Background(), exampleAddress, aptos.AptosCoin, ledger.LedgerVersion)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("account balance: %d\n", balance)
}

func TestReadExampleAccount(t *testing.T) {
	// new account
	exampleWallet, err := wallet.NewFromKeygenFile("account_example")
	if err != nil {
		panic(err)
	}
	exampleAddress := exampleWallet.Address()
	fmt.Printf("example address: %s\n", exampleAddress)

	// new rpc
	client := aptos.New(rpc.TestNet_RPC)

	// latest ledger
	ledger, aptosErr := client.Ledger(context.Background())
	if aptosErr != nil {
		panic(aptosErr)
	}

	// check account
	balance, aptosErr := client.AccountBalance(context.Background(), exampleAddress, aptos.AptosCoin, ledger.LedgerVersion)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("account balance: %d\n", balance)
}
