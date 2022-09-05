package hello_blockchain

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/examples"
	"github.com/motoko9/aptos-go/rpcmodule"
	"github.com/motoko9/aptos-go/wallet"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_ReadFromResource(t *testing.T) {
	ctx := context.Background()

	fmt.Printf("moudle address: %s\n", moduleAddress)

	// user account
	userPriv := examples.BobPrivateKey
	userAddr := wallet.PublicKey2Address(userPriv.PublicKey())
	fmt.Printf("user address: %s\n", userAddr)

	// todo
	// need to test
	resourceType := fmt.Sprintf("%s::message::MessageHolder", moduleAddress)
	resource, err := client.AccountResourceByAddressAndType(ctx, userAddr, resourceType, 0)
	assert.Equal(t, nil, err)
	messageHold, ok := resource.Object.(*MessageHolder)
	assert.Equal(t, true, ok)
	fmt.Printf("account resource: %v\n", messageHold)
}

func Test_ReadFromModuleMethod(t *testing.T) {
	ctx := context.Background()

	// user account
	userPriv := examples.BobPrivateKey
	userAddr := wallet.PublicKey2Address(userPriv.PublicKey())
	fmt.Printf("user address: %s\n", userAddr)

	userAcc, err := client.Account(ctx, userAddr, 0)
	assert.Equal(t, false, err.IsError())

	payload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      fmt.Sprintf("%s::message::get_message", moduleAddress),
		TypeArguments: []string{},
		Arguments:     []interface{}{userAddr},
	}
	txPayload := &rpcmodule.TransactionPayload{
		Type:   rpcmodule.EntryFunctionPayload,
		Object: payload,
	}

	hash, err := client.SignAndSubmitTransaction(ctx, userAddr, userAcc.SequenceNumber, txPayload, userPriv)
	assert.Equal(t, false, err.IsError())
	fmt.Printf("submit transaction hash: %s\n", hash)

	time.Sleep(5 * time.Second)

	confirmed, err := client.ConfirmTransaction(ctx, hash)
	assert.Equal(t, false, err.IsError())
	fmt.Printf("transfer confirmed: %v\n", confirmed)
}
