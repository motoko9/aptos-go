package move_example

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/faucet"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"io/ioutil"
	"testing"
	"time"
)

func TestMovePublish(t *testing.T) {
	ctx := context.Background()

	// move Module account
	wallet, err := wallet.NewFromKeygenFile("account_move_publish")
	if err != nil {
		panic(err)
	}
	address := wallet.Address()
	fmt.Printf("account move publish address: %s\n", wallet.Address())

	// fund (max: 20000)
	amount := uint64(20000)
	hashes, err := faucet.FundAccount(address, amount)
	if err != nil {
		panic(err)
	}
	fmt.Printf("fund txs: %v\n", hashes)

	time.Sleep(time.Second * 5)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// read move byte code
	content, err := ioutil.ReadFile("./Message.mv")
	if err != nil {
		panic(err)
	}

	// publish message
	tx, err := client.PublishMoveModule(ctx, address, content, wallet)
	if err != nil {
		panic(err)
	}
	//
	fmt.Printf("publish move module transaction hash: %s\n", tx.Hash)

	//
	confirmed, err := client.ConfirmTransaction(ctx, tx.Hash)
	if err != nil {
		panic(err)
	}
	fmt.Printf("publish move module transaction confirmed: %v\n", confirmed)
}
