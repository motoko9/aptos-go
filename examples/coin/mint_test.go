package coin

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/crypto"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_MintTo(t *testing.T) {
	ctx := context.Background()

	// coin account
	privFrom := priv
	pubKFrom := privFrom.PublicKey()
	coinAddress := pubKFrom.Address()
	fmt.Printf("coin address: %s\n", coinAddress)

	// recipient address generate by register_recipient_test
	privTo, _ := crypto.NewPrivateKeyFromFile("addr_registered.key")
	recipientAddress := privTo.PublicKey().Address()

	mintAmount := uint64(1000000000)
	hash, err := client.MintCoin(
		context.Background(),
		coinAddress,
		fmt.Sprintf("%s::moon_coin::MoonCoin", coinAddress),
		recipientAddress,
		mintAmount,
		privFrom)
	assert.Equal(t, nil, err)
	fmt.Printf("submit transaction hash: %s\n", hash)

	time.Sleep(5 * time.Second)

	confirmed, err := client.ConfirmTransaction(ctx, hash)
	assert.Equal(t, nil, err)
	fmt.Printf("transaction confirmed: %v\n", confirmed)
}
