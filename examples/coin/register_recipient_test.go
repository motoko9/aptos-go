package coin

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/crypto"
	"github.com/motoko9/aptos-go/faucet"
	"github.com/motoko9/aptos-go/wallet"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_RegisterRecipient(t *testing.T) {
	ctx := context.Background()

	// coin account
	privFrom := priv
	pubKFrom := privFrom.PublicKey()
	coinAddress := wallet.PublicKey2Address(pubKFrom)
	fmt.Printf("coin address: %s\n", coinAddress)

	// new account
	pubKTo, privTo, err := crypto.NewRandomPrivateKey()
	assert.NoError(t, err)
	addressTo := wallet.PublicKey2Address(pubKTo)
	fmt.Printf("recipient address: %s\n", addressTo)
	privTo.SaveToFile("addr_registered.key")

	// fund (max: 20000)
	amount := uint64(20000)
	hashes, aptosErr := faucet.FundAccount(addressTo, amount)
	assert.Equal(t, nil, aptosErr)
	fmt.Printf("fund txs: %v\n", hashes)

	time.Sleep(time.Second * 5)

	hash, aptosErr := client.RegisterRecipient(ctx, addressTo, aptos.MOONCoin, privTo)
	assert.Equal(t, nil, aptosErr)
	fmt.Printf("submit transaction hash: %s\n", hash)

	time.Sleep(5 * time.Second)

	confirmed, aptosErr := client.ConfirmTransaction(ctx, hash)
	assert.NoError(t, nil, aptosErr)
	fmt.Printf("transaction confirmed: %v\n", confirmed)
}
