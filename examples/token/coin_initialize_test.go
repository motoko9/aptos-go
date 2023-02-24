package move_example

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/aptos"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/rpcmodule"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
)

func TestCoinInitialize(t *testing.T) {
	ctx := context.Background()
	// token account
	usdtWallet, err := wallet.NewFromKeygenFile("account_usdt")
	if err != nil {
		panic(err)
	}
	usdtAddress := usdtWallet.Address()
	fmt.Printf("usdt address: %s\n", usdtAddress)

	// new rpc
	client := aptos.New(rpc.TestNet_RPC, false)

	// from account
	usdtAccount, aptosErr := client.Account(ctx, usdtAddress, 0)
	if aptosErr != nil {
		panic(aptosErr)
	}

	//
	payload := &rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      "0x1::managed_coin::initialize",
		TypeArguments: []string{fmt.Sprintf("%s::usdt::USDT", usdtAddress)},
		Arguments: []interface{}{
			hex.EncodeToString([]byte("usdt")),
			hex.EncodeToString([]byte("USDT")),
			6,
			true,
		},
	}

	payload1 := &rpcmodule.TransactionPayload{
		Type:   rpcmodule.EntryFunctionPayload,
		Object: payload,
	}

	txHash, aptosErr := client.SignAndSubmitTransaction(ctx, usdtAddress, usdtAccount.SequenceNumber, payload1, usdtWallet)
	if aptosErr != nil {
		panic(aptosErr)
	}
	//
	fmt.Printf("token initialize transaction hash: %s\n", txHash)

	//
	confirmed, aptosErr := client.ConfirmTransaction(ctx, txHash)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("token initialize transaction confirmed: %v\n", confirmed)
}
