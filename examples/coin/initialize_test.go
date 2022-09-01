package coin

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Initialize(t *testing.T) {
	ctx := context.Background()

	privFrom := priv
	pubKFrom := privFrom.PublicKey()
	initializeAddress := pubKFrom.Address()
	fmt.Printf("coin initialize address: %s\n", initializeAddress)

	hash, err := client.InitializeCoin(context.Background(), initializeAddress,
		[]string{fmt.Sprintf("%s::moon_coin::MoonCoin", initializeAddress)},
		[]interface{}{
			hex.EncodeToString([]byte("Moon Coin")),
			hex.EncodeToString([]byte("MOON")),
			6,
			false,
		}, privFrom)
	assert.NoError(t, err, "submit transaction failed")
	fmt.Printf("submit transaction hash: %s\n", hash)

	time.Sleep(5 * time.Second)

	confirmed, err := client.ConfirmTransaction(ctx, hash)
	assert.NoError(t, err, "confirmation transaction failed")
	fmt.Printf("transaction confirmed: %v\n", confirmed)
}
