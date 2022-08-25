package move_example

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

func TestCoinInitialize(t *testing.T) {
	ctx := context.Background()

	// coin account
	coinWallet, err := wallet.NewFromKeygenFile("account_example")
	if err != nil {
		panic(err)
	}
	coinAddress := coinWallet.Address()
	fmt.Printf("coin address: %s\n", coinAddress)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// from account
	coinAccount, err := client.Account(ctx, coinAddress, 0)
	if err != nil {
		panic(err)
	}

	//
	payload := rpc.EntryFunctionPayload{
		T:             "entry_function_payload",
		Function:      "0x1::managed_coin::initialize",
		TypeArguments: []string{fmt.Sprintf("%s::usdt::USDTCoin", coinAddress)},
		Arguments: []interface{}{
			hex.EncodeToString([]byte("usdt")),
			hex.EncodeToString([]byte("USDT")),
			"6",
			false,
		},
	}
	transaction := rpc.Transaction{
		T:                       "",
		Hash:                    "",
		Sender:                  coinAddress,
		SequenceNumber:          coinAccount.SequenceNumber,
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
	signature, err := coinWallet.Sign(signData)
	if err != nil {
		panic(err)
	}

	// add signature
	transaction.Signature = &rpc.Signature{
		T: "ed25519_signature",
		//PublicKey: fromAccount.AuthenticationKey,
		PublicKey: "0x" + coinWallet.PublicKey().String(),
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
