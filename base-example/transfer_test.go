package base_example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
)

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
	ledger, err := client.Ledger(ctx)
	if err != nil {
		panic(err)
	}

	// check from account
	{
		balance, err := client.AccountBalance(ctx, addressFrom, aptos.AptosCoin, ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("from account balance: %d\n", balance)
	}

	// check to account
	{
		balance, err := client.AccountBalance(ctx, addressTo, aptos.AptosCoin, ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("to account balance: %d\n", balance)
	}

	// from account
	accountFrom, err := client.Account(ctx, addressFrom, 0)
	if err != nil {
		panic(err)
	}

	transaction, err := client.TransferCoinMsg(addressFrom, accountFrom.SequenceNumber, aptos.AptosCoin, uint64(1000), addressTo)
	if err != nil {
		panic(err)
	}

	// sign message
	signData, err := client.EncodeSubmission(ctx, transaction)
	if err != nil {
		panic(err)
	}

	// sign
	signature, err := walletFrom.Sign(signData)
	if err != nil {
		panic(err)
	}

	// add signature
	transaction.Signature = &rpc.Signature{
		T: "ed25519_signature",
		//PublicKey: fromAccount.AuthenticationKey,
		PublicKey: "0x" + walletFrom.PublicKey().String(),
		Signature: "0x" + hex.EncodeToString(signature),
	}

	// submit
	tx, err := client.SubmitTransaction(ctx, transaction)
	if err != nil {
		panic(err)
	}
	//
	fmt.Printf("transfer hash: %s\n", tx.Hash)

	//
	confirmed, err := client.ConfirmTransaction(ctx, tx.Hash)
	if err != nil {
		panic(err)
	}
	fmt.Printf("transfer confirmed: %v\n", confirmed)

	// check account balance after transfer
	// transfer has confirmed, but balance is not update
	// todo
	{
		balance, err := client.AccountBalance(ctx, addressFrom, aptos.AptosCoin, ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("from account balance: %d\n", balance)
	}
	{
		balance, err := client.AccountBalance(ctx, addressTo, aptos.AptosCoin, ledger.LedgerVersion)
		if err != nil {
			panic(err)
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
	ledger, err := client.Ledger(ctx)
	if err != nil {
		panic(err)
	}

	// check from account
	{
		balance, err := client.AccountBalance(ctx, addressFrom, aptos.AptosCoin, ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("from account balance: %d\n", balance)
	}

	// check to account
	{
		balance, err := client.AccountBalance(ctx, addressTo, aptos.AptosCoin, ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("to account balance: %d\n", balance)
	}

	tx, err := client.TransferCoin(ctx, addressFrom, aptos.AptosCoin, uint64(1000), addressTo, walletFrom)
	if err != nil {
		panic(err)
	}
	//
	fmt.Printf("transfer hash: %s\n", tx.Hash)

	//
	confirmed, err := client.ConfirmTransaction(ctx, tx.Hash)
	if err != nil {
		panic(err)
	}
	fmt.Printf("transfer confirmed: %v\n", confirmed)

	// check account balance after transfer
	// transfer has confirmed, but balance is not update
	// todo
	{
		balance, err := client.AccountBalance(ctx, addressFrom, aptos.AptosCoin, ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("from account balance: %d\n", balance)
	}
	{
		balance, err := client.AccountBalance(ctx, addressTo, aptos.AptosCoin, ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("to account balance: %d\n", balance)
	}
}