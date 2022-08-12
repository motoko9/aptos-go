package move_example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/faucet"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
	"time"
)

func TestRegisterRecipient(t *testing.T) {
	ctx := context.Background()

	// coin account
	coinWallet, err := wallet.NewFromKeygenFile("account_example")
	if err != nil {
		panic(err)
	}
	coinAddress := coinWallet.Address()
	fmt.Printf("coin address: %s\n", coinAddress)

	// new account
	wallet := wallet.New()
	wallet.Save("account_recipient")
	address := wallet.Address()
	fmt.Printf("recipient address: %s\n", address)

	// fund (max: 20000)
	amount := uint64(20000)
	hashes, err := faucet.FundAccount(address, amount)
	if err != nil {
		panic(err)
	}
	fmt.Printf("fund txs: %v\n", hashes)

	time.Sleep(time.Second * 5)

	// new rpc
	client := rpc.New(rpc.DevNet_RPC)

	// recipient account
	account, err := client.Account(ctx, address)
	if err != nil {
		panic(err)
	}

	//
	payload := rpc.Payload{
		T:             "script_function_payload",
		Function:      "0x1::coins::register",
		TypeArguments: []string{fmt.Sprintf("%s::usdt::USDTCoin", coinAddress)},
		Arguments:     []interface{}{},
	}
	transaction := rpc.Transaction{
		T:                       "",
		Hash:                    "",
		Sender:                  address,
		SequenceNumber:          account.SequenceNumber,
		MaxGasAmount:            uint64(2000),
		GasUnitPrice:            uint64(1),
		GasCurrencyCode:         "",
		ExpirationTimestampSecs: uint64(time.Now().Unix() + 600), // now + 10 minutes
		Payload:                 &payload,
		Signature:               nil,
	}

	// sign message
	signData, err := client.SignMessage(ctx, &transaction)
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
	tx, err := client.SubmitTransaction(ctx, &transaction)
	if err != nil {
		panic(err)
	}
	//
	fmt.Printf("transaction hash: %s\n", tx.Hash)

	//
	confirmed, err := client.ConfirmTransaction(ctx, tx.Hash)
	if err != nil {
		panic(err)
	}
	fmt.Printf("transaction confirmed: %v\n", confirmed)
}
