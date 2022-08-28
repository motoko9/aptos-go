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

func TestNewUsdtAccount(t *testing.T) {
	ctx := context.Background()

	// new account
	wallet := wallet.New()
	wallet.Save("account_usdt")
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

func TestReadUsdtAccount(t *testing.T) {
	ctx := context.Background()

	// new account
	wallet, err := wallet.NewFromKeygenFile("account_usdt")
	if err != nil {
		panic(err)
	}
	address := wallet.Address()
	fmt.Printf("address: %s\n", address)
	fmt.Printf("public key: %s\n", wallet.PublicKey().String())
	fmt.Printf("private key: %s\n", wallet.PrivateKey.String())

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

func TestCoinPublish(t *testing.T) {
	ctx := context.Background()

	// coin account
	wallet, err := wallet.NewFromKeygenFile("account_usdt")
	if err != nil {
		panic(err)
	}
	address := wallet.Address()
	fmt.Printf("coin publish address: %s\n", wallet.Address())

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// read move byte code
	content, err := ioutil.ReadFile("./usdt.mv")
	if err != nil {
		panic(err)
	}

	// publish message
	txHash, err := client.PublishMoveModuleLegacy(ctx, address, content, wallet)
	if err != nil {
		panic(err)
	}
	//
	fmt.Printf("publish move rpcmodule transaction hash: %s\n", txHash)

	//
	confirmed, err := client.ConfirmTransaction(ctx, txHash)
	if err != nil {
		panic(err)
	}
	fmt.Printf("publish move rpcmodule transaction confirmed: %v\n", confirmed)
}
