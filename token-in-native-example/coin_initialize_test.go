package move_example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/rpcmodule"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
)

func TestCoinInitialize(t *testing.T) {
	ctx := context.Background()

	// coin account
	coinWallet, err := wallet.NewFromKeygenFile("account_usdt")
	if err != nil {
		panic(err)
	}
	coinAddress := coinWallet.Address()
	fmt.Printf("coin address: %s\n", coinAddress)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// from account
	coinAccount, aptosErr := client.Account(ctx, coinAddress, 0)
	if aptosErr != nil {
		panic(aptosErr)
	}

	//
	payload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      "0x1::managed_coin::initialize",
		TypeArguments: []string{fmt.Sprintf("%s::usdt::USDT", coinAddress)},
		Arguments: []interface{}{
			hex.EncodeToString([]byte("usdt")),
			hex.EncodeToString([]byte("USDT")),
			//"usdt",
			//"USDT",
			6,
			true,
		},
	}
	encodeSubmissionReq := rpcmodule.EncodeSubmissionReq(coinAddress, coinAccount.SequenceNumber, &rpcmodule.TransactionPayload{
		Type:   rpcmodule.EntryFunctionPayload,
		Object: payload,
	})
	if err != nil {
		panic(err)
	}

	// sign message
	signData, aptosErr := client.EncodeSubmission(ctx, encodeSubmissionReq)
	if aptosErr != nil {
		panic(aptosErr)
	}

	// sign
	signature, err := coinWallet.Sign(signData)
	if err != nil {
		panic(err)
	}

	// add signature
	submitReq := rpcmodule.SubmitTransactionReq(encodeSubmissionReq, rpcmodule.Signature{
		Type: rpcmodule.Ed25519Signature,
		Object: rpcmodule.SignatureEd25519Signature{
			Type:      rpcmodule.Ed25519Signature,
			PublicKey: "0x" + coinWallet.PublicKey().String(),
			Signature: "0x" + hex.EncodeToString(signature),
		},
	})

	// submit
	txHash, aptosErr := client.SubmitTransaction(ctx, submitReq)
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
