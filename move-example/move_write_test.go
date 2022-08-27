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

func TestMoveWrite(t *testing.T) {
	ctx := context.Background()

	// move Module account
	moveModule, err := wallet.NewFromKeygenFile("account_move_publish")
	if err != nil {
		panic(err)
	}
	moduleAddress := moveModule.Address()
	fmt.Printf("move rpcmodule address: %s\n", moduleAddress)

	// user account
	wallet := wallet.New()
	wallet.Save("account_user")
	address := wallet.Address()
	fmt.Printf("user address: %s\n", address)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// from account
	account, err := client.Account(ctx, address, 0)
	if err != nil {
		panic(err)
	}

	//
	message := []byte("hello world!")
	payload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          "entry_function_payload",
		Function:      fmt.Sprintf("%s::Message::set_message", moduleAddress),
		TypeArguments: []string{},
		Arguments:     []interface{}{hex.EncodeToString(message)},
	}
	encodeSubmissionReq, err := rpcmodule.EncodeSubmissionReq(
		address, account.SequenceNumber, rpcmodule.TransactionPayload{
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
	signature, err := wallet.Sign(signData)
	if err != nil {
		panic(err)
	}

	submitReq, err := rpcmodule.SubmitTransactionReq(encodeSubmissionReq, rpcmodule.AccountSignature{
		Type: "ed25519_signature",
		Object: rpcmodule.AccountSignatureEd25519Signature{
			Type:      "ed25519_signature",
			PublicKey: "0x" + wallet.PublicKey().String(),
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
