package aptos

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/rpcmodule"
	"github.com/motoko9/aptos-go/wallet"
	"testing"
	"time"
)

func TestCoinTransfer(t *testing.T) {
	userWallet, err := wallet.NewFromKey("PrivateKey")
	if err != nil {
		panic(err)
	}
	sender := userWallet.Address()
	fmt.Printf("user address: %s\n", userWallet.Address())

	addresses := "0xb043fd1ea38da0779c2f463725f7998c5b7f641efd350c9f91e118b95b956094"
	amounts := fmt.Sprintf("%d", 2000*100000000)
	client := New(rpc.MainNet_RPC, true)
	for i := 0; i < 1; i++ {
		payload := &rpcmodule.TransactionPayloadEntryFunctionPayload{
			Type:          rpcmodule.EntryFunctionPayload,
			Function:      "0x1::coin::transfer",
			TypeArguments: []string{"0x786bee5d40d1e057e749cf8ca5a86bf4b84cd751d0b9c9e596a0ad68caa542b3::asset::USDC"},
			//TypeArguments: []string{"0x1::aptos_coin::AptosCoin"},
			Arguments: []interface{}{addresses, amounts},
		}
		payload1 := &rpcmodule.TransactionPayload{
			Type:   rpcmodule.EntryFunctionPayload,
			Object: payload,
		}
		account, aptosErr := client.Account(context.Background(), sender, 0)
		if aptosErr != nil {
			panic(aptosErr)
		}
		txHash, aptosErr := client.SignAndSubmitTransaction(context.Background(), sender, account.SequenceNumber, payload1, userWallet)
		if aptosErr != nil {
			panic(aptosErr)
		}
		fmt.Printf("tx hash: %s\n", txHash)
		time.Sleep(time.Second * 2)
	}
}

func TestAptosTransfer(t *testing.T) {
	userWallet, err := wallet.NewFromKey("PrivateKey")
	if err != nil {
		panic(err)
	}
	sender := userWallet.Address()
	fmt.Printf("user address: %s\n", userWallet.Address())

	addresses := []string{"0xb043fd1ea38da0779c2f463725f7998c5b7f641efd350c9f91e118b95b956094"}
	amounts := []string{fmt.Sprintf("%d", 2000*100000000)}
	client := New(rpc.MainNet_RPC, true)
	for i, _ := range addresses {
		for j := 0; j < 1; j++ {
			payload := &rpcmodule.TransactionPayloadEntryFunctionPayload{
				Type:          rpcmodule.EntryFunctionPayload,
				Function:      "0x1::aptos_account::transfer_coins",
				TypeArguments: []string{"0x786bee5d40d1e057e749cf8ca5a86bf4b84cd751d0b9c9e596a0ad68caa542b3::asset::USDC"},
				Arguments:     []interface{}{addresses[i], amounts[i]},
			}
			payload1 := &rpcmodule.TransactionPayload{
				Type:   rpcmodule.EntryFunctionPayload,
				Object: payload,
			}
			account, aptosErr := client.Account(context.Background(), sender, 0)
			if aptosErr != nil {
				panic(aptosErr)
			}
			txHash, aptosErr := client.SignAndSubmitTransaction(context.Background(), sender, account.SequenceNumber, payload1, userWallet)
			if aptosErr != nil {
				panic(aptosErr)
			}
			fmt.Printf("tx hash: %s\n", txHash)

			time.Sleep(time.Second * 5)
		}
	}
}

func TestBatchAptosTransfer(t *testing.T) {
	userWallet, err := wallet.NewFromKey("PrivateKey")
	if err != nil {
		panic(err)
	}
	sender := userWallet.Address()
	fmt.Printf("user address: %s\n", userWallet.Address())

	addresses := []string{"0xb043fd1ea38da0779c2f463725f7998c5b7f641efd350c9f91e118b95b956094"}
	amounts := []string{fmt.Sprintf("%d", 2000*100000000)}
	client := New(rpc.MainNet_RPC, true)
	payload := &rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      "0x1::aptos_account::batch_transfer_coins",
		TypeArguments: []string{"0x786bee5d40d1e057e749cf8ca5a86bf4b84cd751d0b9c9e596a0ad68caa542b3::asset::USDC"},
		Arguments:     []interface{}{addresses, amounts},
	}
	payload1 := &rpcmodule.TransactionPayload{
		Type:   rpcmodule.EntryFunctionPayload,
		Object: payload,
	}
	account, aptosErr := client.Account(context.Background(), sender, 0)
	if aptosErr != nil {
		panic(aptosErr)
	}
	txHash, aptosErr := client.SignAndSubmitTransaction(context.Background(), sender, account.SequenceNumber, payload1, userWallet)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("tx hash: %s\n", txHash)
}
