package tokenswap_example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
	"time"
)

func TestCreatePool(t *testing.T) {
	ctx := context.Background()

	// swap Module account
	swapWallet, err := wallet.NewFromKeygenFile("account_swap")
	if err != nil {
		panic(err)
	}
	address := swapWallet.Address()
	fmt.Printf("move module address: %s\n", address)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// from account
	account, err := client.Account(ctx, address, 0)
	if err != nil {
		panic(err)
	}

	// create pool
	coin1 := aptos.CoinType[aptos.AptosCoin]
	coin2 := aptos.CoinType[aptos.USDTCoin]
	payload := rpc.EntryFunctionPayload{
		T:             "entry_function_payload",
		Function:      fmt.Sprintf("%s::swap::create_pool", address),
		TypeArguments: []string{coin1, coin2},
		Arguments:     []interface{}{},
	}
	transaction := rpc.Transaction{
		T:                       "",
		Hash:                    "",
		Sender:                  address,
		SequenceNumber:          account.SequenceNumber,
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
	signature, err := swapWallet.Sign(signData)
	if err != nil {
		panic(err)
	}

	// add signature
	transaction.Signature = &rpc.Signature{
		T: "ed25519_signature",
		//PublicKey: fromAccount.AuthenticationKey,
		PublicKey: "0x" + swapWallet.PublicKey().String(),
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
