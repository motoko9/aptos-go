package hello_blockchain

import (
    "context"
    "fmt"
    "github.com/motoko9/aptos-go/examples"
    "github.com/motoko9/aptos-go/rpcmodule"
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

func Test_Write(t *testing.T) {
    ctx := context.Background()

    fmt.Printf("moudle address: %s\n", moduleAddress)

    userPriv := examples.BobPrivateKey
    userAddr := userPriv.PublicKey().Address()
    fmt.Printf("user address: %s\n", userAddr)

    userAcc, err := client.Account(ctx, userAddr, 0)
    assert.NoError(t, err, "user account is not found")

    message := "hello world!"
    payload := rpcmodule.TransactionPayloadEntryFunctionPayload{
        Type:          rpcmodule.EntryFunctionPayload,
        Function:      fmt.Sprintf("%s::message::set_message", moduleAddress),
        TypeArguments: []string{},
        Arguments:     []interface{}{message},
    }
    txPayload := &rpcmodule.TransactionPayload{
        Type:   rpcmodule.EntryFunctionPayload,
        Object: payload,
    }

    req, err := rpcmodule.EncodeSubmissionReq(userAddr, userAcc.SequenceNumber, txPayload)
    assert.NoError(t, err)

    hash, err := client.SignAndSubmitTransaction(ctx, req, userPriv)
    assert.NoError(t, err, "submit transaction failed")
    fmt.Printf("submit transaction hash: %s\n", hash)

    time.Sleep(5 * time.Second)

    confirmed, err := client.ConfirmTransaction(ctx, hash)
    assert.NoError(t, err, "confirm transaction failed")
    fmt.Printf("transfer confirmed: %v\n", confirmed)
}
