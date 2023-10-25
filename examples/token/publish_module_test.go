package move_example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/wallet"
	"io/ioutil"
	"testing"
	"time"
)

func TestNewUsdtAccount(t *testing.T) {
	// new account
	usdtWallet := wallet.New()
	usdtWallet.SaveToKeygenFile("account_usdt")
	usdtAddress := usdtWallet.Address()
	fmt.Printf("new usdt address: %s\n", usdtAddress)

	// new rpc
	client := aptos.New(rpc.TestNet_RPC, false)

	//
	faultWallet, err := wallet.NewFromKeygenFile("./../account_fault")
	if err != nil {
		panic(err)
	}
	faultAddress := faultWallet.Address()
	fmt.Printf("fault address: %s\n", faultAddress)
	fmt.Printf("fault public key: %s\n", faultWallet.PublicKey().String())
	fmt.Printf("fault private key: %s\n", faultWallet.PrivateKey.String())

	// create account on aptos
	txHash, aptosErr := client.CreateAccount(context.Background(), faultAddress, usdtAddress, faultWallet)
	if aptosErr != nil {
		panic(err)
	}
	fmt.Printf("create ussdt account transacion: %s\n", txHash)

	time.Sleep(time.Second * 5)

	// fund
	txHash, aptosErr = client.TransferCoin(context.Background(), faultAddress, "0x1::aptos_coin::AptosCoin", 100000000, usdtAddress, faultWallet)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("fund usdt account transacion: %s\n", txHash)

	//
	time.Sleep(time.Second * 5)

	// latest ledger
	ledger, aptosErr := client.Ledger(context.Background())
	if aptosErr != nil {
		panic(aptosErr)
	}

	// check account
	balance, aptosErr := client.AccountBalance(context.Background(), usdtAddress, aptos.CoinAlias("APT", "aptos"), ledger.LedgerVersion)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("account balance: %d\n", balance)
}

func TestReadUsdtAccount(t *testing.T) {
	// new account
	wallet, err := wallet.NewFromKeygenFile("account_usdt")
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
	fmt.Printf("account balance: %d\n", balance)
}

func TestModulePublish(t *testing.T) {
	// token account
	wallet, err := wallet.NewFromKeygenFile("account_usdt")
	if err != nil {
		panic(err)
	}
	address := wallet.Address()
	fmt.Printf("usdt token publish address: %s\n", wallet.Address())

	// new rpc
	client := aptos.New(rpc.TestNet_RPC, false)

	// read move byte code
	content, err := ioutil.ReadFile("./usdc.mv")
	if err != nil {
		panic(err)
	}

	meta, err := ioutil.ReadFile("./package-metadata.bcs")
	if err != nil {
		panic(err)
	}

	// publish message
	txHash, aptosErr := client.PublishMoveModule(context.Background(), address, content, meta, wallet)
	if aptosErr != nil {
		panic(aptosErr)
	}
	//
	fmt.Printf("publish move module transaction hash: %s\n", txHash)

	//
	confirmed, aptosErr := client.ConfirmTransaction(context.Background(), txHash)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("publish move module transaction confirmed: %v\n", confirmed)
}

func TestModuleCode(t *testing.T) {
	content, err := ioutil.ReadFile("./usdc.mv")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", hex.EncodeToString(content))
}
