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
	coinAccount, err := client.Account(ctx, coinAddress, 0)
	if err != nil {
		panic(err)
	}

	//
	payload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          "entry_function_payload",
		Function:      "0x1::managed_coin::initialize",
		TypeArguments: []string{fmt.Sprintf("%s::usdt::USDTCoin", coinAddress)},
		Arguments: []interface{}{
			hex.EncodeToString([]byte("usdt")),
			hex.EncodeToString([]byte("USDT")),
			"6",
			false,
		},
	}
	encodeSubmissionReq, err := rpcmodule.EncodeSubmissionReq(coinAddress, coinAccount.SequenceNumber, rpcmodule.TransactionPayload{
		Type:   "entry_function_payload",
		Object: payload,
	})
	if err != nil {
		panic(err)
	}

	// sign message
	signData, err := client.EncodeSubmission(ctx, encodeSubmissionReq)
	if err != nil {
		panic(err)
	}

	// sign
	signature, err := coinWallet.Sign(signData)
	if err != nil {
		panic(err)
	}

	// add signature
	submitReq, err := rpcmodule.SubmitTransactionReq(encodeSubmissionReq, rpcmodule.AccountSignature{
		Type: "ed25519_signature",
		Object: rpcmodule.AccountSignatureEd25519Signature{
			Type:      "ed25519_signature",
			PublicKey: "0x" + coinWallet.PublicKey().String(),
			Signature: "0x" + hex.EncodeToString(signature),
		},
	})

	// submit
	txHash, err := client.SubmitTransaction(ctx, submitReq)
	if err != nil {
		panic(err)
	}
	//
	fmt.Printf("transaction hash: %s\n", txHash)

	//
	confirmed, err := client.ConfirmTransaction(ctx, txHash)
	if err != nil {
		panic(err)
	}
	fmt.Printf("transaction confirmed: %v\n", confirmed)
}
