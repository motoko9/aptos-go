package example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/faucet"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/utils"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
	"time"
)

func TestTransfer(t *testing.T) {
	ctx := context.Background()

	// new account
	wallet := wallet.New()
	address := utils.Pubkey2Address(wallet.PublicKey())
	fmt.Printf("address: %s\n", address)

	// fund (max: 20000)
	fundAmount := uint64(20000)
	hashs, err := faucet.FundAccount(address, fundAmount)
	if err != nil {
		panic(err)
	}
	fmt.Printf("fund txs: %v\n", hashs)

	//
	time.Sleep(time.Second * 5)

	// new rpc
	client := rpc.New(rpc.DevNet_RPC)

	// latest ledger
	ledger, err := client.Ledger(ctx)
	if err != nil {
		panic(err)
	}

	// check from account
	from := address
	{
		balance, err := client.AccountBalance(ctx, from, "AptosCoin", ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("from account balance: %d\n", balance)
	}

	// check to account
	to := "0x228f33506ef0c00a39db928664d9bbddf83cb1c0ca5df1a2c85f9ff619870f73"
	{
		balance, err := client.AccountBalance(ctx, to, "AptosCoin", ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("to account balance: %d\n", balance)
	}

	// from account
	fromAccount, err := client.Account(ctx, from)
	if err != nil {
		panic(err)
	}

	// transfer
	transferAmount := uint64(1000)
	transferPayload := rpc.Payload{
		Function:      "0x1::coin::transfer",
		Arguments:     []string{to, fmt.Sprintf("%d", transferAmount)},
		T:             "script_function_payload",
		TypeArguments: []interface{}{"0x1::aptos_coin::AptosCoin"},
	}
	// now + 10 minutes
	maxGasAmount := uint64(2000)
	gasUnitPrice := uint64(1)
	expirationTimestampSecs := uint64(time.Now().Unix() + 600)

	signData, err := client.SignMessage(ctx, from, fromAccount.SequenceNumber, maxGasAmount, gasUnitPrice, expirationTimestampSecs, transferPayload)
	if err != nil {
		panic(err)
	}

	// sign
	signature, err := wallet.PrivateKey.Sign(signData)
	if err != nil {
		panic(err)
	}

	// submit
	tx, err := client.SubmitTransaction(ctx, from, fromAccount.SequenceNumber, maxGasAmount, gasUnitPrice, expirationTimestampSecs, transferPayload, rpc.Signature{
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
		balance, err := client.AccountBalance(ctx, from, "AptosCoin", ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("from account balance: %d\n", balance)
	}
	{
		balance, err := client.AccountBalance(ctx, to, "AptosCoin", ledger.LedgerVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("to account balance: %d\n", balance)
	}
}
