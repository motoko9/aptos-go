package base_example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/faucet"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/rpcmodule"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
	"time"
)

func TestNewToAccount(t *testing.T) {
	ctx := context.Background()

	// new account
	wallet := wallet.New()
	wallet.Save("account_to")
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
	client := aptos.New(aptos.DevNet_RPC)

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

func TestReadToAccount(t *testing.T) {
	ctx := context.Background()

	// new account
	wallet, err := wallet.NewFromKeygenFile("account_to")
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

func TestTransfer_raw(t *testing.T) {
	ctx := context.Background()

	// account
	walletFrom, err := wallet.NewFromKeygenFile("account_example")
	if err != nil {
		panic(err)
	}
	addressFrom := walletFrom.Address()
	fmt.Printf("from address: %s\n", addressFrom)

	walletTo, err := wallet.NewFromKeygenFile("account_to")
	if err != nil {
		panic(err)
	}
	addressTo := walletTo.Address()
	fmt.Printf("to address: %s\n", addressTo)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// latest ledger
	ledger, aptosErr := client.Ledger(ctx)
	if aptosErr != nil {
		panic(aptosErr)
	}

	// check from account
	{
		balance, aptosErr := client.AccountBalance(ctx, addressFrom, aptos.AptosCoin, ledger.LedgerVersion)
		if aptosErr != nil {
			panic(aptosErr)
		}
		fmt.Printf("from account balance: %d\n", balance)
	}

	// check to account
	{
		balance, aptosErr := client.AccountBalance(ctx, addressTo, aptos.AptosCoin, ledger.LedgerVersion)
		if aptosErr != nil {
			panic(aptosErr)
		}
		fmt.Printf("to account balance: %d\n", balance)
	}

	// from account
	accountFrom, aptosErr := client.Account(ctx, addressFrom, 0)
	if aptosErr != nil {
		panic(aptosErr)
	}

	payload, aptosErr := aptos.TransferCoinPayload(aptos.AptosCoin, uint64(1000), addressTo)
	if aptosErr != nil {
		panic(aptosErr)
	}

	encodeSubmissionReq := rpcmodule.EncodeSubmissionReq(addressFrom, accountFrom.SequenceNumber, payload)

	// sign message
	signData, aptosErr := client.EncodeSubmission(ctx, encodeSubmissionReq)
	if aptosErr != nil {
		panic(aptosErr)
	}

	// sign
	signature, err := walletFrom.Sign(signData)
	if err != nil {
		panic(err)
	}

	// add signature
	submitReq := rpcmodule.SubmitTransactionReq(encodeSubmissionReq, rpcmodule.Signature{
		Type: rpcmodule.Ed25519Signature,
		Object: rpcmodule.SignatureEd25519Signature{
			Type:      rpcmodule.Ed25519Signature,
			PublicKey: "0x" + walletFrom.PublicKey().String(),
			Signature: "0x" + hex.EncodeToString(signature),
		},
	})

	// submit
	txHash, aptosErr := client.SubmitTransaction(ctx, submitReq)
	if aptosErr != nil {
		panic(aptosErr)
	}
	//
	fmt.Printf("transfer hash: %s\n", txHash)

	//
	confirmed, aptosErr := client.ConfirmTransaction(ctx, txHash)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("transfer confirmed: %v\n", confirmed)

	// check account balance after transfer
	// transfer has confirmed, but balance is not update
	// todo
	{
		balance, aptosErr := client.AccountBalance(ctx, addressFrom, aptos.AptosCoin, ledger.LedgerVersion)
		if aptosErr != nil {
			panic(aptosErr)
		}
		fmt.Printf("from account balance: %d\n", balance)
	}
	{
		balance, aptosErr := client.AccountBalance(ctx, addressTo, aptos.AptosCoin, ledger.LedgerVersion)
		if aptosErr != nil {
			panic(aptosErr)
		}
		fmt.Printf("to account balance: %d\n", balance)
	}
}

func TestTransfer(t *testing.T) {
	ctx := context.Background()

	// account
	walletFrom, err := wallet.NewFromKeygenFile("account_example")
	if err != nil {
		panic(err)
	}
	addressFrom := walletFrom.Address()
	fmt.Printf("from address: %s\n", addressFrom)

	walletTo, err := wallet.NewFromKeygenFile("account_to")
	if err != nil {
		panic(err)
	}
	addressTo := walletTo.Address()
	fmt.Printf("to address: %s\n", addressTo)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// latest ledger
	ledger, aptosErr := client.Ledger(ctx)
	if aptosErr != nil {
		panic(aptosErr)
	}

	// check from account
	{
		balance, aptosErr := client.AccountBalance(ctx, addressFrom, aptos.AptosCoin, ledger.LedgerVersion)
		if aptosErr != nil {
			panic(aptosErr)
		}
		fmt.Printf("from account balance: %d\n", balance)
	}

	// check to account
	{
		balance, aptosErr := client.AccountBalance(ctx, addressTo, aptos.AptosCoin, ledger.LedgerVersion)
		if aptosErr != nil {
			panic(aptosErr)
		}
		fmt.Printf("to account balance: %d\n", balance)
	}

	txHash, aptosErr := client.TransferCoin(ctx, addressFrom, aptos.AptosCoin, uint64(1000), addressTo, walletFrom)
	if aptosErr != nil {
		panic(aptosErr)
	}
	//
	fmt.Printf("transfer hash: %s\n", txHash)

	//
	confirmed, aptosErr := client.ConfirmTransaction(ctx, txHash)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("transfer confirmed: %v\n", confirmed)

	// check account balance after transfer
	// transfer has confirmed, but balance is not update
	// todo
	{
		balance, aptosErr := client.AccountBalance(ctx, addressFrom, aptos.AptosCoin, ledger.LedgerVersion)
		if aptosErr != nil {
			panic(aptosErr)
		}
		fmt.Printf("from account balance: %d\n", balance)
	}
	{
		balance, aptosErr := client.AccountBalance(ctx, addressTo, aptos.AptosCoin, ledger.LedgerVersion)
		if aptosErr != nil {
			panic(aptosErr)
		}
		fmt.Printf("to account balance: %d\n", balance)
	}
}
