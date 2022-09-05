package coin

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/wallet"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Initialize(t *testing.T) {
	ctx := context.Background()

	privFrom := priv
	pubKFrom := privFrom.PublicKey()
	initializeAddress := wallet.PublicKey2Address(pubKFrom)
	fmt.Printf("coin initialize address: %s\n", initializeAddress)

	hash, aptosErr := client.InitializeCoin(
		context.Background(),
		initializeAddress,
		fmt.Sprintf("%s::moon_coin::MoonCoin", initializeAddress),
		hex.EncodeToString([]byte("Moon Coin")),
		hex.EncodeToString([]byte("MOON")),
		6,
		privFrom)
	assert.Equal(t, nil, aptosErr)
	fmt.Printf("submit transaction hash: %s\n", hash)

	time.Sleep(5 * time.Second)

	confirmed, aptosErr := client.ConfirmTransaction(ctx, hash)
	assert.Equal(t, nil, aptosErr)
	fmt.Printf("transaction confirmed: %v\n", confirmed)
}
