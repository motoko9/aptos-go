package move_example

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/faucet"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"io/ioutil"
	"testing"
	"time"
)

func TestNewModuleAccount(t *testing.T) {
	ctx := context.Background()

	// new account
	wallet := wallet.New()
	wallet.Save("account_helloworld")
	address := wallet.Address()
	fmt.Printf("address: %s\n", address)

	// fund (max: 20000)
	amount := uint64(20000)
	hashes, err := faucet.FundAccount(address, amount)
	if err != nil {
		panic(err)
	}
	fmt.Printf("fund txs: %v\n", hashes)

	//
	time.Sleep(time.Second * 5)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// latest ledger
	ledger, err := client.Ledger(ctx)
	if err != nil {
		panic(err)
	}

	// check account
	balance, err := client.AccountBalance(ctx, address, aptos.AptosCoin, ledger.LedgerVersion)
	if err != nil {
		panic(err)
	}
	fmt.Printf("account balance: %d\n", balance)
}

func TestReadModuleAccount(t *testing.T) {
	ctx := context.Background()

	// new account
	wallet, err := wallet.LoadFromKeygenFile("account_helloworld")
	if err != nil {
		panic(err)
	}
	address := wallet.Address()
	fmt.Printf("address: %s\n", address)
	fmt.Printf("public key: %s\n", wallet.PublicKey().String())
	fmt.Printf("private key: %s\n", wallet.PrivateKey.String())

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

// todo
// can not working
// please publish move by aptos cli
func TestMovePublish(t *testing.T) {
	ctx := context.Background()

	// move Module account
	wallet, err := wallet.NewFromKeygenFile("account_helloworld")
	if err != nil {
		panic(err)
	}
	address := wallet.Address()
	fmt.Printf("account move publish address: %s\n", wallet.Address())

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// read move byte code
	content, err := ioutil.ReadFile("./helloworld.mv")
	if err != nil {
		panic(err)
	}

	// publish message
	txHash, aptosErr := client.PublishMoveModuleLegacy(ctx, address, content, wallet)
	if aptosErr != nil {
		panic(aptosErr)
	}
	//
	fmt.Printf("publish move module transaction hash: %s\n", txHash)

	//
	confirmed, aptosErr := client.ConfirmTransaction(ctx, txHash)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("publish move module transaction confirmed: %v\n", confirmed)
}
