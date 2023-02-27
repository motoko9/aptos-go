package move_example

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/rpcmodule"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
	"time"
)

func TestNewUserAccount(t *testing.T) {
	// new account
	userWallet := wallet.New()
	userWallet.SaveToKeygenFile("account_user")
	userAddress := userWallet.Address()
	fmt.Printf("user address: %s\n", userAddress)

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
	txHash, aptosErr := client.CreateAccount(context.Background(), faultAddress, userAddress, faultWallet)
	if aptosErr != nil {
		panic(err)
	}
	fmt.Printf("create user account transacion: %s\n", txHash)

	//
	time.Sleep(time.Second * 10)

	// fund
	txHash, aptosErr = client.TransferCoin(context.Background(), faultAddress, aptos.CoinAlias("APT", "aptos"), 100000000, userAddress, faultWallet)
	if aptosErr != nil {
		panic(err)
	}
	fmt.Printf("fund user account transacion: %s\n", txHash)

	//
	time.Sleep(time.Second * 10)

	// latest ledger
	ledger, aptosErr := client.Ledger(context.Background())
	if aptosErr != nil {
		panic(aptosErr)
	}

	// check account
	balance, aptosErr := client.AccountBalance(context.Background(), userAddress, aptos.CoinAlias("APT", "aptos"), ledger.LedgerVersion)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("account balance: %d\n", balance)
}

func TestReadUserAccount(t *testing.T) {
	// new account
	wallet, err := wallet.NewFromKeygenFile("account_user")
	if err != nil {
		panic(err)
	}
	address := wallet.Address()
	fmt.Printf("address: %s\n", address)
	fmt.Printf("public key: %s\n", wallet.PublicKey().String())
	fmt.Printf("private key: %s\n", wallet.PrivateKey.String())

	//
	time.Sleep(time.Second * 5)

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

func TestMoveRead(t *testing.T) {
	// move Module account
	moveModule, err := wallet.NewFromKeygenFile("account_move")
	if err != nil {
		panic(err)
	}
	moduleAddress := moveModule.Address()
	fmt.Printf("move module address: %s\n", moduleAddress)

	// user account
	userWallet, err := wallet.NewFromKeygenFile("account_user")
	if err != nil {
		panic(err)
	}
	address := userWallet.Address()
	fmt.Printf("user address: %s\n", address)

	// new rpc
	client := aptos.New(rpc.TestNet_RPC, false)

	// todo,
	// can not ready resource type Message::MessageHolder
	// only support CoinStore type
	// need update AccountResourceByAddressAndType
	//
	resourceType := fmt.Sprintf("%s::helloworld::MessageHolder", moduleAddress)
	accountResource, aptosErr := client.AccountResourceByAddressAndType(context.Background(), address, resourceType, 0)
	if aptosErr != nil {
		panic(aptosErr)
	}
	accountResourceJson, _ := json.MarshalIndent(accountResource, "", "    ")
	fmt.Printf("account resource: %s\n", string(accountResourceJson))
}

func TestMoveView(t *testing.T) {
	// move Module account
	moveModule, err := wallet.NewFromKeygenFile("account_move")
	if err != nil {
		panic(err)
	}
	moduleAddress := moveModule.Address()
	fmt.Printf("move module address: %s\n", moduleAddress)

	// user account
	userWallet, err := wallet.NewFromKeygenFile("account_user")
	if err != nil {
		panic(err)
	}
	userAddress := userWallet.Address()
	fmt.Printf("user address: %s\n", userAddress)

	// new rpc
	client := aptos.New(rpc.TestNet_RPC, false)

	function := fmt.Sprintf("%s::helloworld::get_message", moduleAddress)
	raw, aptosErr := client.View(context.Background(), &rpcmodule.ViewRequest{
		Function:      function,
		TypeArguments: []string{},
		Arguments:     []interface{}{userAddress},
	})
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("result: %s\n", raw)
}
