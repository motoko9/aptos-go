package move_example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/faucet"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/rpcmodule"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
	"time"
)

func TestNewToAccount(t *testing.T) {
	ctx := context.Background()

	// new account
	wallet := wallet.New()
	wallet.Save("account_to")
	address := wallet.Address()
	fmt.Printf("address: %s\n", address)

	// fund (max: 20000)
	amount := uint64(20000)
	hashes, err := faucet.FundAccount(address, amount)
	if err != nil {
		panic(err)
	}
	fmt.Printf("fund txs: %v\n", hashes)

	//
	time.Sleep(time.Second * 5)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// latest ledger
	ledger, err := client.Ledger(ctx)
	if err != nil {
		panic(err)
	}

	// check account
	balance, err := client.AccountBalance(ctx, address, aptos.AptosCoin, ledger.LedgerVersion)
	if err != nil {
		panic(err)
	}
	fmt.Printf("account balance: %d\n", balance)
}

func TestReadToAccount(t *testing.T) {
	ctx := context.Background()

	// new account
	wallet, err := wallet.NewFromKeygenFile("account_to")
	if err != nil {
		panic(err)
	}
	address := wallet.Address()
	fmt.Printf("address: %s\n", address)

	// fund (max: 20000)
	amount := uint64(20000)
	hashes, aptosErr := faucet.FundAccount(address, amount)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("fund txs: %v\n", hashes)

	//
	time.Sleep(time.Second * 5)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// latest ledger
	ledger, aptosErr := client.Ledger(ctx)
	if aptosErr != nil {
		panic(aptosErr)
	}

	// check account
	balance, aptosErr := client.AccountBalance(ctx, address, aptos.AptosCoin, ledger.LedgerVersion)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("account balance: %d\n", balance)
}

// need to register to account first !!!!
//
func TestTransfer(t *testing.T) {
	ctx := context.Background()

	// coin account
	coinWallet, err := wallet.NewFromKeygenFile("account_usdc")
	if err != nil {
		panic(err)
	}
	coinAddress := coinWallet.Address()
	fmt.Printf("coin address: %s\n", coinAddress)

	// recipient account
	fromWallet, err := wallet.NewFromKeygenFile("account_recipient")
	if err != nil {
		panic(err)
	}
	fromAddress := fromWallet.Address()
	fmt.Printf("from address: %s\n", fromAddress)

	// to account
	toWallet, err := wallet.NewFromKeygenFile("account_to")
	if err != nil {
		panic(err)
	}
	toAddress := toWallet.Address()
	fmt.Printf("to address: %s\n", toAddress)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// coin account
	fromAccount, aptosErr := client.Account(ctx, fromAddress, 0)
	if aptosErr != nil {
		panic(aptosErr)
	}

	//
	transferAmount := uint64(50000000)
	payload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      fmt.Sprintf("%s::usdc::transfer", coinAddress),
		TypeArguments: []string{},
		Arguments: []interface{}{
			toAddress,
			fmt.Sprintf("%d", transferAmount),
		},
	}
	encodeSubmissionReq := rpcmodule.EncodeSubmissionReq(fromAddress, fromAccount.SequenceNumber, &rpcmodule.TransactionPayload{
		Type:   rpcmodule.EntryFunctionPayload,
		Object: payload,
	})
	if err != nil {
		panic(err)
	}

	// sign message
	signData, aptosErr := client.EncodeSubmission(ctx, encodeSubmissionReq)
	if aptosErr != nil {
		panic(aptosErr)
	}

	// sign
	signature, err := fromWallet.Sign(signData)
	if err != nil {
		panic(err)
	}

	// add signature
	submitReq := rpcmodule.SubmitTransactionReq(encodeSubmissionReq, rpcmodule.Signature{
		Type: rpcmodule.Ed25519Signature,
		Object: rpcmodule.SignatureEd25519Signature{
			Type: rpcmodule.Ed25519Signature,
			//PublicKey: fromAccount.AuthenticationKey,
			PublicKey: "0x" + fromWallet.PublicKey().String(),
			Signature: "0x" + hex.EncodeToString(signature),
		},
	})
	if err != nil {
		panic(err)
	}

	// submit
	txHash, aptosErr := client.SubmitTransaction(ctx, submitReq)
	if aptosErr != nil {
		panic(aptosErr)
	}
	//
	fmt.Printf("transaction hash: %s\n", txHash)

	//
	confirmed, aptosErr := client.ConfirmTransaction(ctx, txHash)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("transaction confirmed: %v\n", confirmed)
}
