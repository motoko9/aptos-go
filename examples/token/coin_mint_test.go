package move_example

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/rpcmodule"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
	"time"
)

func TestNewMintAccount(t *testing.T) {
	// new account
	mintWallet := wallet.New()
	mintWallet.SaveToKeygenFile("account_mint")
	mintAddress := mintWallet.Address()
	fmt.Printf("mint address: %s\n", mintAddress)

	//
	faultWallet, err := wallet.NewFromKeygenFile("./../account_fault")
	if err != nil {
		panic(err)
	}
	faultAddress := faultWallet.Address()
	fmt.Printf("fault address: %s\n", faultAddress)
	fmt.Printf("fault public key: %s\n", faultWallet.PublicKey().String())
	fmt.Printf("fault private key: %s\n", faultWallet.PrivateKey.String())

	// new rpc
	client := aptos.New(rpc.TestNet_RPC, false)

	// create account on aptos
	txHash, aptosErr := client.CreateAccount(context.Background(), faultAddress, mintAddress, faultWallet)
	if aptosErr != nil {
		panic(err)
	}
	fmt.Printf("create mint account transacion: %s\n", txHash)

	// fund
	txHash, aptosErr = client.TransferCoin(context.Background(), faultAddress, aptos.CoinAlias("APT", "aptos"), 100000000, mintAddress, faultWallet)
	if aptosErr != nil {
		panic(err)
	}
	fmt.Printf("fund mint account transacion: %s\n", txHash)

	//
	time.Sleep(time.Second * 5)

	// latest ledger
	ledger, aptosErr := client.Ledger(context.Background())
	if aptosErr != nil {
		panic(aptosErr)
	}

	// check account
	balance, aptosErr := client.AccountBalance(context.Background(), mintAddress, aptos.CoinAlias("APT", "aptos"), ledger.LedgerVersion)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("account balance: %d\n", balance)
}

func TestReadMintAccount(t *testing.T) {
	// new account
	wallet, err := wallet.NewFromKeygenFile("account_mint")
	if err != nil {
		panic(err)
	}
	address := wallet.Address()
	fmt.Printf("address: %s\n", address)
	fmt.Printf("public key: %s\n", wallet.PublicKey().String())
	fmt.Printf("private key: %s\n", wallet.PrivateKey.String())

	// new rpc
	client := aptos.New(rpc.TestNet_RPC, false)

	// latest ledger
	ledger, aptosErr := client.Ledger(context.Background())
	if aptosErr != nil {
		panic(aptosErr)
	}

	// check account
	balance, aptosErr := client.AccountBalance(context.Background(), address, aptos.CoinAlias("APT", "aptos"), ledger.LedgerVersion)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("account apt balance: %d\n", balance)

	// check account
	balance, aptosErr = client.AccountBalance(context.Background(), address, aptos.CoinAlias("USDT", "wormhole"), ledger.LedgerVersion)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("account usdt balance: %d\n", balance)
}

func TestMint(t *testing.T) {
	// token account
	coinWallet, err := wallet.NewFromKeygenFile("account_usdt")
	if err != nil {
		panic(err)
	}
	coinAddress := coinWallet.Address()
	fmt.Printf("token address: %s\n", coinAddress)

	// recipient account
	userWallet, err := wallet.NewFromKeygenFile("account_mint")
	if err != nil {
		panic(err)
	}
	address := userWallet.Address()
	fmt.Printf("recipient address: %s\n", address)

	// new rpc
	client := aptos.New(rpc.TestNet_RPC, false)

	// token account
	coinAccount, aptosErr := client.Account(context.Background(), coinAddress, 0)
	if aptosErr != nil {
		panic(aptosErr)
	}

	//
	mintAmount := uint64(1000000000)
	payload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      "0x1::managed_coin::mint",
		TypeArguments: []string{fmt.Sprintf("%s::usdt::USDT", coinAddress)},
		Arguments: []interface{}{
			address,
			fmt.Sprintf("%d", mintAmount),
		},
	}

	payload1 := &rpcmodule.TransactionPayload{
		Type:   rpcmodule.EntryFunctionPayload,
		Object: payload,
	}

	txHash, aptosErr := client.SignAndSubmitTransaction(context.Background(), coinAddress, coinAccount.SequenceNumber, payload1, coinWallet)
	if aptosErr != nil {
		panic(aptosErr)
	}
	//
	fmt.Printf("token mint transaction hash: %s\n", txHash)

	//
	confirmed, aptosErr := client.ConfirmTransaction(context.Background(), txHash)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("token mint transaction confirmed: %v\n", confirmed)
}
