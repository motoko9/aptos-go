package move_example

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
	"time"
)

func TestNewReceiptAccount(t *testing.T) {
	// new account
	receiptWallet := wallet.New()
	receiptWallet.SaveToKeygenFile("account_receipt")
	receiptAddress := receiptWallet.Address()
	fmt.Printf("receipt address: %s\n", receiptAddress)

	// new rpc
	client := aptos.New(rpc.TestNet_RPC, false)

	//
	faultWallet, err := wallet.NewFromKeygenFile("./../account_fault")
	if err != nil {
		panic(err)
	}
	faultAddress := faultWallet.Address()
	fmt.Printf("fault address: %s\n", faultAddress)

	hash, aptosErr := client.CreateAccount(context.Background(), faultAddress, receiptAddress, faultWallet)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("create receipt account transaction: %s\n", hash)

	// fund
	txHash, aptosErr := client.TransferCoin(context.Background(), faultAddress, aptos.AptosCoin, 100000000, receiptAddress, faultWallet)
	if aptosErr != nil {
		panic(err)
	}
	fmt.Printf("fund receipt account transacion: %s\n", txHash)

	// register
	_, aptosErr = client.RegisterRecipient(context.Background(), receiptAddress, aptos.USDTCoin, receiptWallet)
	if aptosErr != nil {
		panic(aptosErr)
	}

	//
	time.Sleep(time.Second * 5)

	// latest ledger
	ledger, aptosErr := client.Ledger(context.Background())
	if aptosErr != nil {
		panic(aptosErr)
	}

	// check account
	balance, aptosErr := client.AccountBalance(context.Background(), receiptAddress, aptos.AptosCoin, ledger.LedgerVersion)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("account balance: %d\n", balance)

	// check account
	balance, aptosErr = client.AccountBalance(context.Background(), receiptAddress, aptos.USDTCoin, ledger.LedgerVersion)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("account balance: %d\n", balance)
}

func TestReadToAccount(t *testing.T) {
	// new account
	wallet, err := wallet.NewFromKeygenFile("account_receipt")
	if err != nil {
		panic(err)
	}
	address := wallet.Address()
	fmt.Printf("address: %s\n", address)

	// new rpc
	client := aptos.New(rpc.TestNet_RPC, false)

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

	// check account
	balance, aptosErr = client.AccountBalance(context.Background(), address, aptos.USDTCoin, ledger.LedgerVersion)
	if err != nil {
		panic(aptosErr)
	}
	fmt.Printf("account balance: %d\n", balance)
}

func TestTransfer(t *testing.T) {
	ctx := context.Background()

	// token account
	coinWallet, err := wallet.NewFromKeygenFile("account_usdt")
	if err != nil {
		panic(err)
	}
	coinAddress := coinWallet.Address()
	fmt.Printf("token address: %s\n", coinAddress)

	// recipient account
	fromWallet, err := wallet.NewFromKeygenFile("account_mint")
	if err != nil {
		panic(err)
	}
	fromAddress := fromWallet.Address()
	fmt.Printf("from address: %s\n", fromAddress)

	// to account
	toWallet, err := wallet.NewFromKeygenFile("account_receipt")
	if err != nil {
		panic(err)
	}
	toAddress := toWallet.Address()
	fmt.Printf("to address: %s\n", toAddress)

	// new rpc
	client := aptos.New(rpc.TestNet_RPC, false)
	transferAmount := uint64(50000000)
	txHash, aptosErr := client.TransferCoin(ctx, fromAddress, aptos.USDTCoin, transferAmount, toAddress, fromWallet)
	if aptosErr != nil {
		panic(aptosErr)
	}
	//
	fmt.Printf("transaction hash: %s\n", txHash)

	//
	confirmed, aptosErr := client.ConfirmTransaction(ctx, txHash)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("transaction confirmed: %v\n", confirmed)
}
