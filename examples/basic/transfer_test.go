package basic

import (
    "context"
    "fmt"
    "github.com/hashicorp/go-hclog"
    "github.com/motoko9/aptos-go/aptos"
    "github.com/motoko9/aptos-go/crypto"
    "github.com/motoko9/aptos-go/examples"
    "github.com/motoko9/aptos-go/rpc"
    "github.com/motoko9/aptos-go/rpcmodule"
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

var client *aptos.Client
var priv crypto.PrivateKey
var to string

func init() {
    client = aptos.NewClient(rpc.DevNet_RPC, hclog.Default())
    priv = examples.PrivateKey
    to = examples.AlicePrivateKey.PublicKey().Address()
}

func Test_Transfer_Aptos(t *testing.T) {
    ctx := context.Background()

    // account
    privFrom := priv
    pubKFrom := privFrom.PublicKey()
    addressFrom := pubKFrom.Address()
    fmt.Printf("from address: %s\n", addressFrom)

    // latest ledger
    ledger, err := client.Ledger(ctx)
    assert.NoError(t, err)

    // check from account
    {
        balance, err := client.AccountBalance(ctx, addressFrom, aptos.AptosCoin, ledger.LedgerVersion)
        assert.NoError(t, err)
        fmt.Printf("from account balance: %d\n", balance)
    }

    //check to account
    {
        balance, err := client.AccountBalance(ctx, to, aptos.AptosCoin, ledger.LedgerVersion)
        assert.NoError(t, err)
        fmt.Printf("to account balance: %d\n", balance)
    }

    hash, err := client.TransferCoin(ctx, addressFrom, aptos.AptosCoin, uint64(100), to, privFrom)
    assert.NoError(t, err)
    fmt.Printf("submit transaction hash: %s\n", hash)

    time.Sleep(5 * time.Second)

    confirmed, err := client.ConfirmTransaction(ctx, hash)
    assert.NoError(t, err, "confirm transaction failed")
    fmt.Printf("transfer confirmed: %v\n", confirmed)

    // update latest ledger to update balance result
    // or else balance will not update
    ledger, err = client.Ledger(ctx)
    assert.NoError(t, err)
    {
        balance, err := client.AccountBalance(ctx, addressFrom, aptos.AptosCoin, ledger.LedgerVersion)
        if err != nil {
            panic(err)
        }
        fmt.Printf("from account balance: %d\n", balance)
    }
    {
        balance, err := client.AccountBalance(ctx, to, aptos.AptosCoin, ledger.LedgerVersion)
        if err != nil {
            panic(err)
        }
        fmt.Printf("to account balance: %d\n", balance)
    }
}

// TODO Binary Canonical Serialization (BCS) to get raw data to sign
func Test_CreateSigningMessage(t *testing.T) {
    privFrom := priv
    pubKFrom := privFrom.PublicKey()
    addressFrom := pubKFrom.Address()
    accountFrom, e1 := client.Account(context.Background(), addressFrom, 0)
    assert.NoError(t, e1)

    addressTo := "0x4c80f1fe097f290528975c49ae8c64ce0c3cf673a16876471962910f4ecea74e"

    payload, err := aptos.TransferCoinPayload(aptos.AptosCoin, uint64(100), addressTo)
    assert.NoError(t, err)

    req, err := rpcmodule.EncodeSubmissionReq(addressFrom, accountFrom.SequenceNumber, payload)
    assert.NoError(t, err)
    str, err := client.TransactionEncodeSubmission(context.Background(), req)
    assert.NoError(t, err)
    fmt.Println(str)

    // TODO
    //bytes, err := lcs.Marshal(tx)
    //assert.NoError(t, err)
    //fmt.Printf("raw data: %v\n", hex.EncodeToString(bytes))
}
