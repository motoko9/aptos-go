package coin

import (
    "context"
    "fmt"
    "github.com/hashicorp/go-hclog"
    "github.com/motoko9/aptos-go/aptos"
    "github.com/motoko9/aptos-go/crypto"
    "github.com/motoko9/aptos-go/examples"
    "github.com/motoko9/aptos-go/rpc"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "testing"
    "time"
)

var client *aptos.Client
var priv crypto.PrivateKey

func init() {
    client = aptos.NewClient(rpc.DevNet_RPC, hclog.Default())
    priv = examples.PrivateKey
}

func Test_Deploy(t *testing.T) {
    ctx := context.Background()

    privFrom := priv
    pubKFrom := privFrom.PublicKey()
    deployAddress := pubKFrom.Address()
    fmt.Printf("coin deploy address: %s\n", deployAddress)

    // read move byte code
    content, err := ioutil.ReadFile("./moon_coin.mv")
    assert.NoError(t, err, "read mv failed")

    hash, err := client.PublishMoveModuleLegacy(ctx, deployAddress, content, privFrom)
    assert.NoError(t, err)
    fmt.Printf("submit transaction hash: %s\n", hash)

    time.Sleep(5 * time.Second)

    confirmed, err := client.ConfirmTransaction(ctx, hash)
    assert.NoError(t, err, "confirm transaction failed")
    fmt.Printf("publish move module transaction confirmed: %v\n", confirmed)
}
