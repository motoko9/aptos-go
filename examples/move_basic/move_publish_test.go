package move_example

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"io/ioutil"
	"testing"
	"time"
)

func TestNewModuleAccount(t *testing.T) {
	// new account
	moveWallet := wallet.New()
	moveWallet.SaveToKeygenFile("account_move")
	moveAddress := moveWallet.Address()
	fmt.Printf("move address: %s\n", moveAddress)

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
	txHash, aptosErr := client.CreateAccount(context.Background(), faultAddress, moveAddress, faultWallet)
	if aptosErr != nil {
		panic(err)
	}
	fmt.Printf("create move account transacion: %s\n", txHash)

	//
	time.Sleep(time.Second * 10)

	// fund
	txHash, aptosErr = client.TransferCoin(context.Background(), faultAddress, aptos.AptosCoin, 100000000, moveAddress, faultWallet)
	if aptosErr != nil {
		panic(err)
	}
	fmt.Printf("fund move account transacion: %s\n", txHash)

	//
	time.Sleep(time.Second * 10)

	// latest ledger
	ledger, aptosErr := client.Ledger(context.Background())
	if aptosErr != nil {
		panic(aptosErr)
	}

	// check account
	balance, aptosErr := client.AccountBalance(context.Background(), moveAddress, aptos.AptosCoin, ledger.LedgerVersion)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("account balance: %d\n", balance)
}

func TestReadModuleAccount(t *testing.T) {
	// new account
	wallet, err := wallet.NewFromKeygenFile("account_move")
	if err != nil {
		panic(err)
	}
	address := wallet.Address()
	fmt.Printf("address: %s\n", address)
	fmt.Printf("public key: %s\n", wallet.PublicKey().String())
	fmt.Printf("private key: %s\n", wallet.PrivateKey.String())

	// new rpc
	client := aptos.New(rpc.TestNet_RPC)

	// latest ledger
	ledger, aptosErr := client.Ledger(context.Background())
	if aptosErr != nil {
		panic(aptosErr)
	}

	// check account
	balance, aptosErr := client.AccountBalance(context.Background(), address, aptos.AptosCoin, ledger.LedgerVersion)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("account balance: %d\n", balance)
}

// please publish move by aptos cli
func TestMovePublish(t *testing.T) {
	// move Module account
	wallet, err := wallet.NewFromKeygenFile("account_move")
	if err != nil {
		panic(err)
	}
	address := wallet.Address()
	fmt.Printf("account move publish address: %s\n", wallet.Address())

	// new rpc
	client := aptos.New(rpc.TestNet_RPC)

	// read move byte code
	content, err := ioutil.ReadFile("./example.mv")
	if err != nil {
		panic(err)
	}

	meta, err := ioutil.ReadFile("./package-metadata.bcs")
	if err != nil {
		panic(err)
	}

	// publish message
	txHash, aptosErr := client.PublishMoveModule(context.Background(), address, content, meta, wallet)
	if aptosErr != nil {
		panic(aptosErr)
	}
	//
	fmt.Printf("publish move module transaction hash: %s\n", txHash)

	//
	confirmed, aptosErr := client.ConfirmTransaction(context.Background(), txHash)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("publish move module transaction confirmed: %v\n", confirmed)
}
