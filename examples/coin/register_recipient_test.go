package coin

import (
    "context"
    "fmt"
    "github.com/motoko9/aptos-go/aptos"
    "github.com/motoko9/aptos-go/crypto"
    "github.com/motoko9/aptos-go/faucet"
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

func Test_RegisterRecipient(t *testing.T) {
    ctx := context.Background()

    // coin account
    privFrom := priv
    pubKFrom := privFrom.PublicKey()
    coinAddress := pubKFrom.Address()
    fmt.Printf("coin address: %s\n", coinAddress)

    // new account
    pubKTo, privTo, err := crypto.NewRandomPrivateKey()
    assert.NoError(t, err)
    addressTo := pubKTo.Address()
    fmt.Printf("recipient address: %s\n", addressTo)
    privTo.Save("addr_registered.key")

    // fund (max: 20000)
    amount := uint64(20000)
    hashes, err := faucet.FundAccount(addressTo, amount)
    assert.NoError(t, err, "faucet fund account failed")
    fmt.Printf("fund txs: %v\n", hashes)

    time.Sleep(time.Second * 5)

    hash, err := client.RegisterRecipient(ctx, addressTo, aptos.MOONCoin, privTo)
    assert.NoError(t, err, "submit transaction failed")
    fmt.Printf("submit transaction hash: %s\n", hash)

    time.Sleep(5 * time.Second)

    confirmed, err := client.ConfirmTransaction(ctx, hash)
    assert.NoError(t, err, "confirmation transaction failed")
    fmt.Printf("transaction confirmed: %v\n", confirmed)
}
