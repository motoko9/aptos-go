package tokenswap_example

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

func TestCreatePool(t *testing.T) {
	ctx := context.Background()

	// swap Module account
	swapWallet, err := wallet.NewFromKeygenFile("account_swap")
	if err != nil {
		panic(err)
	}
	address := swapWallet.Address()
	fmt.Printf("move rpcmodule address: %s\n", address)

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
	payload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          "entry_function_payload",
		Function:      fmt.Sprintf("%s::swap::create_pool", address),
		TypeArguments: []string{coin1, coin2},
		Arguments:     []interface{}{},
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
	signature, err := swapWallet.Sign(signData)
	if err != nil {
		panic(err)
	}

	// add signature
	submitReq, err := rpcmodule.SubmitTransactionReq(encodeSubmissionReq, rpcmodule.AccountSignature{
		Type: "ed25519_signature",
		Object: rpcmodule.AccountSignatureEd25519Signature{
			Type:      "ed25519_signature",
			PublicKey: "0x" + swapWallet.PublicKey().String(),
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
