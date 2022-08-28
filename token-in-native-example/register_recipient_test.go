package move_example

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

func TestNewRecipientAccount(t *testing.T) {
	ctx := context.Background()

	// new account
	wallet := wallet.New()
	wallet.Save("account_recipient")
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

func TestReadRecipientAccount(t *testing.T) {
	ctx := context.Background()

	// new account
	wallet, err := wallet.NewFromKeygenFile("account_recipient")
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

func TestRegisterRecipient(t *testing.T) {
	ctx := context.Background()

	// coin account
	coinWallet, err := wallet.NewFromKeygenFile("account_usdt")
	if err != nil {
		panic(err)
	}
	coinAddress := coinWallet.Address()
	fmt.Printf("coin address: %s\n", coinAddress)

	// new account
	userWallet, err := wallet.NewFromKeygenFile("account_recipient")
	if err != nil {
		panic(err)
	}
	address := userWallet.Address()
	fmt.Printf("recipient address: %s\n", address)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)
	txHash, aptosErr := client.RegisterRecipient(ctx, address, aptos.USDTCoin, userWallet)
	if aptosErr != nil {
		panic(aptosErr)
	}

	//
	confirmed, aptosErr := client.ConfirmTransaction(ctx, txHash)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("transaction confirmed: %v\n", confirmed)
}
