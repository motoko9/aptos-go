package move_example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/faucet"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
	"time"
)

func TestTransfer(t *testing.T) {
	ctx := context.Background()

	// coin account
	coinWallet, err := wallet.NewFromKeygenFile("account_example")
	if err != nil {
		panic(err)
	}
	coinAddress := coinWallet.Address()
	fmt.Printf("coin address: %s\n", coinAddress)

	// recipient account
	fromWallet, err := wallet.NewFromKeygenFile("account_recipient")
	if err != nil {
		panic(err)
	}
	fromAddress := fromWallet.Address()
	fmt.Printf("from address: %s\n", fromAddress)

	// fund (max: 20000)
	fundAmount := uint64(20000)
	{
		hashes, err := faucet.FundAccount(fromAddress, fundAmount)
		if err != nil {
			panic(err)
		}
		fmt.Printf("fund txs: %v\n", hashes)
	}

	// to account
	toWallet, err := wallet.NewFromKeygenFile("account_to")
	if err != nil {
		panic(err)
	}
	toAddress := toWallet.Address()
	fmt.Printf("to address: %s\n", toAddress)

	// fund (max: 20000)
	{
		hashes, err := faucet.FundAccount(toAddress, fundAmount)
		if err != nil {
			panic(err)
		}
		fmt.Printf("fund txs: %v\n", hashes)
	}

	time.Sleep(time.Second * 5)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// coin account
	fromAccount, err := client.Account(ctx, fromAddress, 0)
	if err != nil {
		panic(err)
	}

	//
	transferAmount := uint64(50000000)
	payload := rpc.EntryFunctionPayload{
		T:             "entry_function_payload",
		Function:      "0x1::coin::transfer",
		TypeArguments: []string{fmt.Sprintf("%s::usdt::USDTCoin", coinAddress)},
		Arguments: []interface{}{
			toAddress,
			fmt.Sprintf("%d", transferAmount),
		},
	}
	transaction := rpc.Transaction{
		T:                       "",
		Hash:                    "",
		Sender:                  fromAddress,
		SequenceNumber:          fromAccount.SequenceNumber,
		MaxGasAmount:            uint64(2000),
		GasUnitPrice:            uint64(1),
		GasCurrencyCode:         "",
		ExpirationTimestampSecs: uint64(time.Now().Unix() + 600), // now + 10 minutes
		Payload:                 &payload,
		Signature:               nil,
	}

	// sign message
	signData, err := client.EncodeSubmission(ctx, &transaction)
	if err != nil {
		panic(err)
	}

	// sign
	signature, err := fromWallet.Sign(signData)
	if err != nil {
		panic(err)
	}

	// add signature
	transaction.Signature = &rpc.Signature{
		T: "ed25519_signature",
		//PublicKey: fromAccount.AuthenticationKey,
		PublicKey: "0x" + fromWallet.PublicKey().String(),
		Signature: "0x" + hex.EncodeToString(signature),
	}

	// submit
	tx, err := client.SubmitTransaction(ctx, &transaction)
	if err != nil {
		panic(err)
	}
	//
	fmt.Printf("transaction hash: %s\n", tx.Hash)

	//
	confirmed, err := client.ConfirmTransaction(ctx, tx.Hash)
	if err != nil {
		panic(err)
	}
	fmt.Printf("transaction confirmed: %v\n", confirmed)
}
