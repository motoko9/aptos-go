package tokenswap_example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/faucet"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
	"time"
)

func TestNewUserAccount(t *testing.T) {
	ctx := context.Background()

	// new account
	wallet := wallet.New()
	wallet.Save("account_user")
	address := wallet.Address()
	fmt.Printf("user address: %s\n", address)

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

func TestReadUserAccount(t *testing.T) {
	ctx := context.Background()

	// new account
	wallet, err := wallet.NewFromKeygenFile("account_user")
	if err != nil {
		panic(err)
	}
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

func TestAddLiquidity(t *testing.T) {
	ctx := context.Background()

	// swap account
	swapWallet, err := wallet.NewFromKeygenFile("account_swap")
	if err != nil {
		panic(err)
	}
	swapAddress := swapWallet.Address()
	fmt.Printf("swap module publish address: %s\n", swapAddress)

	userWallet, err := wallet.NewFromKeygenFile("account_user")
	if err != nil {
		panic(err)
	}
	userAddress := userWallet.Address()
	fmt.Printf("user address: %s\n", userAddress)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// from account
	account, err := client.Account(ctx, userAddress, 0)
	if err != nil {
		panic(err)
	}

	// add liquidity
	// todo
	coin1 := aptos.CoinType[aptos.AptosCoin]
	coin2 := aptos.CoinType[aptos.USDTCoin]
	payload := rpc.EntryFunctionPayload{
		T:             "entry_function_payload",
		Function:      fmt.Sprintf("%s::swap::add_liquidity", swapAddress),
		TypeArguments: []string{coin1, coin2},
		Arguments:     []interface{}{0, 0, 0, 0},
	}
	transaction := rpc.Transaction{
		T:                       "",
		Hash:                    "",
		Sender:                  userAddress,
		SequenceNumber:          account.SequenceNumber,
		MaxGasAmount:            uint64(2000),
		GasUnitPrice:            uint64(1),
		GasCurrencyCode:         "",
		ExpirationTimestampSecs: uint64(time.Now().Unix() + 600), // now + 10 minutes
		Payload:                 &payload,
		Signature:               nil,
	}

	// sign message
	signData, err := client.EncodeSubmission(ctx, &transaction)
	if err != nil {
		panic(err)
	}

	// sign
	signature, err := userWallet.Sign(signData)
	if err != nil {
		panic(err)
	}

	// add signature
	transaction.Signature = &rpc.Signature{
		T: "ed25519_signature",
		//PublicKey: fromAccount.AuthenticationKey,
		PublicKey: "0x" + userWallet.PublicKey().String(),
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

func TestRemoveLiquidity(t *testing.T) {
	ctx := context.Background()

	// swap account
	swapWallet, err := wallet.NewFromKeygenFile("account_swap")
	if err != nil {
		panic(err)
	}
	swapAddress := swapWallet.Address()
	fmt.Printf("swap module publish address: %s\n", swapAddress)

	userWallet, err := wallet.NewFromKeygenFile("account_user")
	if err != nil {
		panic(err)
	}
	userAddress := userWallet.Address()
	fmt.Printf("user address: %s\n", userAddress)

	// new rpc
	client := aptos.New(rpc.DevNet_RPC)

	// from account
	account, err := client.Account(ctx, userAddress, 0)
	if err != nil {
		panic(err)
	}

	// remove liquidity
	// todo
	coin1 := aptos.CoinType[aptos.AptosCoin]
	coin2 := aptos.CoinType[aptos.USDTCoin]
	payload := rpc.EntryFunctionPayload{
		T:             "entry_function_payload",
		Function:      fmt.Sprintf("%s::swap::remove_liquidity", swapAddress),
		TypeArguments: []string{coin1, coin2},
		Arguments:     []interface{}{0, 0, 0},
	}
	transaction := rpc.Transaction{
		T:                       "",
		Hash:                    "",
		Sender:                  userAddress,
		SequenceNumber:          account.SequenceNumber,
		MaxGasAmount:            uint64(2000),
		GasUnitPrice:            uint64(1),
		GasCurrencyCode:         "",
		ExpirationTimestampSecs: uint64(time.Now().Unix() + 600), // now + 10 minutes
		Payload:                 &payload,
		Signature:               nil,
	}

	// sign message
	signData, err := client.EncodeSubmission(ctx, &transaction)
	if err != nil {
		panic(err)
	}

	// sign
	signature, err := userWallet.Sign(signData)
	if err != nil {
		panic(err)
	}

	// add signature
	transaction.Signature = &rpc.Signature{
		T: "ed25519_signature",
		//PublicKey: fromAccount.AuthenticationKey,
		PublicKey: "0x" + userWallet.PublicKey().String(),
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