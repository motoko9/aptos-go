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
	coinWallet, err := wallet.NewFromKeygenFile("account_usdc")
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
		Type:          "entry_function_payload",
		Function:      fmt.Sprintf("%s::usdc::initialize", coinAddress),
		TypeArguments: []string{},
		Arguments: []interface{}{
			hex.EncodeToString([]byte("usdc")),
			hex.EncodeToString([]byte("USDC")),
			"6",
		},
	}
	encodeSubmissionReq, err := rpcmodule.EncodeSubmissionReq(
		coinAddress, coinAccount.SequenceNumber, rpcmodule.TransactionPayload{
			Type:   "entry_function_payload",
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
	submitReq, err := rpcmodule.SubmitTransactionReq(encodeSubmissionReq, rpcmodule.AccountSignature{
		Type: "ed25519_signature",
		Object: rpcmodule.AccountSignatureEd25519Signature{
			Type:      "ed25519_signature",
			PublicKey: "0x" + coinWallet.PublicKey().String(),
			Signature: "0x" + hex.EncodeToString(signature),
		},
	})
	if err != nil {
		panic(err)
	}

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
