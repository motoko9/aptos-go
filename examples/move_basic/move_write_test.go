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
	// move Module account
	moveModule, err := wallet.NewFromKeygenFile("account_move")
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
	userAddress := userWallet.Address()
	fmt.Printf("user address: %s\n", userAddress)

	// new rpc
	client := aptos.New(rpc.TestNet_RPC)

	// from account
	userAccount, aptosErr := client.Account(context.Background(), userAddress, 0)
	if aptosErr != nil {
		panic(aptosErr)
	}

	//
	message := []byte("hello world!")
	payload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      fmt.Sprintf("%s::helloworld::set_message", moduleAddress),
		TypeArguments: []string{},
		Arguments:     []interface{}{hex.EncodeToString(message)},
	}
	payload1 := &rpcmodule.TransactionPayload{
		Type:   rpcmodule.EntryFunctionPayload,
		Object: payload,
	}

	txHash, aptosErr := client.SignAndSubmitTransaction(context.Background(), userAddress, userAccount.SequenceNumber, payload1, userWallet)
	if aptosErr != nil {
		panic(aptosErr)
	}

	//
	confirmed, aptosErr := client.ConfirmTransaction(context.Background(), txHash)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("transaction confirmed: %v\n", confirmed)
}
