package example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/utils"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
	"time"
)

func TestTransfer(t *testing.T) {
	ctx := context.Background()

	// account
	walletFrom, err := wallet.NewFromKeygenFile("account_from")
	if err != nil {
		panic(err)
	}
	addressFrom := utils.Pubkey2Address(walletFrom.PublicKey())
	fmt.Printf("from address: %s\n", addressFrom)

	walletTo, err := wallet.NewFromKeygenFile("account_to")
	if err != nil {
		panic(err)
	}
	addressTo := utils.Pubkey2Address(walletTo.PublicKey())
	fmt.Printf("to address: %s\n", addressTo)

	// new rpc
	client := rpc.New(rpc.DevNet_RPC)

	// latest ledger
	ledger, err := client.Ledger(ctx)
	if err != nil {
		panic(err)
	}

	// check from account
	{
		balance, err := client.AccountBalance(ctx, addressFrom, "AptosCoin", ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("from account balance: %d\n", balance)
	}

	// check to account
	{
		balance, err := client.AccountBalance(ctx, addressTo, "AptosCoin", ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("to account balance: %d\n", balance)
	}

	// from account
	fromAccount, err := client.Account(ctx, addressFrom)
	if err != nil {
		panic(err)
	}

	// transfer
	transferAmount := uint64(1000)
	transferPayload := rpc.Payload{
		Function:      "0x1::coin::transfer",
		Arguments:     []string{addressTo, fmt.Sprintf("%d", transferAmount)},
		T:             "script_function_payload",
		TypeArguments: []interface{}{"0x1::aptos_coin::AptosCoin"},
	}
	// now + 10 minutes
	maxGasAmount := uint64(2000)
	gasUnitPrice := uint64(1)
	expirationTimestampSecs := uint64(time.Now().Unix() + 600)

	signData, err := client.SignMessage(ctx, addressFrom, fromAccount.SequenceNumber, maxGasAmount, gasUnitPrice, expirationTimestampSecs, transferPayload)
	if err != nil {
		panic(err)
	}

	// sign
	signature, err := walletFrom.PrivateKey.Sign(signData)
	if err != nil {
		panic(err)
	}

	// submit
	tx, err := client.SubmitTransaction(ctx, addressFrom, fromAccount.SequenceNumber, maxGasAmount, gasUnitPrice, expirationTimestampSecs, transferPayload, rpc.Signature{
		T:         "ed25519_signature",
		PublicKey: fromAccount.AuthenticationKey,
		Signature: "0x" + hex.EncodeToString(signature),
	})
	if err != nil {
		panic(err)
	}
	//
	fmt.Printf("transfer hash: %s\n", tx.Hash)

	time.Sleep(time.Second * 5)

	// check account balance after transfer
	{
		balance, err := client.AccountBalance(ctx, addressFrom, "AptosCoin", ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("from account balance: %d\n", balance)
	}
	{
		balance, err := client.AccountBalance(ctx, addressTo, "AptosCoin", ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("to account balance: %d\n", balance)
	}
}
