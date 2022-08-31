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
	moveModule, err := wallet.NewFromKeygenFile("account_helloworld")
	if err != nil {
		panic(err)
	}
	moduleAddress := moveModule.Address()
	fmt.Printf("move module address: %s\n", moduleAddress)

	// user account
	userWallet, err := wallet.NewFromKeygenFile("account_user")
	if err != nil {
		panic(err)
	}
	address := userWallet.Address()
	fmt.Printf("user address: %s\n", address)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// from account
	account, aptosErr := client.Account(ctx, address, 0)
	if aptosErr != nil {
		panic(aptosErr)
	}

	//
	message := []byte("hello world!")
	payload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          "entry_function_payload",
		Function:      fmt.Sprintf("%s::helloworld::set_message", moduleAddress),
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
	signData, aptosErr := client.EncodeSubmission(ctx, encodeSubmissionReq)
	if aptosErr != nil {
		panic(aptosErr)
	}

	// sign
	signature, err := userWallet.Sign(signData)
	if err != nil {
		panic(err)
	}

	submitReq, err := rpcmodule.SubmitTransactionReq(encodeSubmissionReq, rpcmodule.Signature{
		Type: "ed25519_signature",
		Object: rpcmodule.SignatureEd25519Signature{
			Type:      "ed25519_signature",
			PublicKey: "0x" + userWallet.PublicKey().String(),
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
