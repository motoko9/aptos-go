package move_example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/faucet"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"io/ioutil"
	"testing"
	"time"
)

func TestCoinPublish(t *testing.T) {
	ctx := context.Background()

	// coin account
	wallet, err := wallet.NewFromKeygenFile("account_example")
	if err != nil {
		panic(err)
	}
	address := wallet.Address()
	fmt.Printf("coin publish address: %s\n", wallet.Address())

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

	// from account
	account, err := client.Account(ctx, address)
	if err != nil {
		panic(err)
	}

	// read move byte code
	content, err := ioutil.ReadFile("./usdt.mv")
	if err != nil {
		panic(err)
	}

	// publish message
	transaction, err := client.PublishMoveModule(ctx, address, account.SequenceNumber, content)
	if err != nil {
		panic(err)
	}

	// sign message
	signData, err := client.SignMessage(ctx, transaction)
	if err != nil {
		panic(err)
	}

	// sign
	signature, err := wallet.Sign(signData)
	if err != nil {
		panic(err)
	}

	// add signature
	transaction.Signature = &rpc.Signature{
		T: "ed25519_signature",
		//PublicKey: fromAccount.AuthenticationKey,
		PublicKey: "0x" + wallet.PublicKey().String(),
		Signature: "0x" + hex.EncodeToString(signature),
	}

	// submit
	tx, err := client.SubmitTransaction(ctx, transaction)
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
