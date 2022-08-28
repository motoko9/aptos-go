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

func TestMint(t *testing.T) {
	ctx := context.Background()

	// coin account
	coinWallet, err := wallet.NewFromKeygenFile("account_usdc")
	if err != nil {
		panic(err)
	}
	coinAddress := coinWallet.Address()
	fmt.Printf("coin address: %s\n", coinAddress)

	// recipient account
	userWallet, err := wallet.NewFromKeygenFile("account_recipient")
	if err != nil {
		panic(err)
	}
	address := userWallet.Address()
	fmt.Printf("recipient address: %s\n", address)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// coin account
	coinAccount, err := client.Account(ctx, coinAddress, 0)
	if err != nil {
		panic(err)
	}

	//
	mintAmount := uint64(1000000000)
	payload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          "entry_function_payload",
		Function:      fmt.Sprintf("%s::usdc::mint", coinAddress),
		TypeArguments: []string{},
		Arguments: []interface{}{
			address,
			fmt.Sprintf("%d", mintAmount),
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
	if err != nil {
		panic(err)
	}

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
