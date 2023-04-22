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

func TestAnimeSwap_AddLiquidity(t *testing.T) {
	userWallet, err := wallet.NewFromKey("f476ba25a9df047f8d4c024896a171c60f32eb31b89bccbbbf1462b46e0475e3")
	if err != nil {
		panic(err)
	}
	fmt.Printf("user address: %s\n", userWallet.Address())

	client := New(rpc.TestNet_RPC, false)
	txHash, aptosErr := client.AnimeAddLiquidity(context.Background(), userWallet.Address(),
		"0x2f88a12a17f01228f4ba72ec6214127abb930512dcb3d6205909ca510aca7b29::asset::WETH",
		1000000000,
		"0x2f88a12a17f01228f4ba72ec6214127abb930512dcb3d6205909ca510aca7b29::asset::USDT",
		1200000000000,
		userWallet)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("add liquidity tx hash: %s\n", txHash)
}

func TestAnimeSwap_View(t *testing.T) {
	//
	userWallet, err := wallet.NewFromKey("f476ba25a9df047f8d4c024896a171c60f32eb31b89bccbbbf1462b46e0475e3")
	if err != nil {
		panic(err)
	}
	fmt.Printf("user address: %s\n", userWallet.Address())
	//
	client := New(rpc.TestNet_RPC, false)
	//
	payload := &rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      fmt.Sprintf("%s::AnimeSwapPoolV1f1::add_liquidity_entry", AnimeSwapAccount),
		TypeArguments: []string{"0x2f88a12a17f01228f4ba72ec6214127abb930512dcb3d6205909ca510aca7b29::asset::WETH", "0x2f88a12a17f01228f4ba72ec6214127abb930512dcb3d6205909ca510aca7b29::asset::USDC"},
		Arguments:     []interface{}{fmt.Sprintf("%d", 1000000000), fmt.Sprintf("%d", 1000000000000), "0", "0"},
	}
	payload1 := &rpcmodule.TransactionPayload{
		Type:   rpcmodule.EntryFunctionPayload,
		Object: payload,
	}
	//
	sender := userWallet.Address()
	account, aptosErr := client.Account(context.Background(), sender, 0)
	if aptosErr != nil {
		panic(aptosErr)
	}

	tx := &rpcmodule.SubmitTransactionRequest{
		Sender:                  sender,
		SequenceNumber:          account.SequenceNumber,
		MaxGasAmount:            uint64(80000),
		GasUnitPrice:            uint64(100),
		ExpirationTimestampSecs: uint64(time.Now().Unix() + 600),
		Payload:                 payload1,
		Signature: rpcmodule.Signature{
			Type: rpcmodule.Ed25519Signature,
			Object: rpcmodule.SignatureEd25519Signature{
				Type:      rpcmodule.Ed25519Signature,
				PublicKey: "0xd36df53c46ca6c046648a85083482149a6423b06a3e35cd2f91b01d656c06e73",
				Signature: "0x945073c4b0d389271b1e6959e5238a1d6f6f82aa6dc09ca6cd31eb6952bde2781b6338d2470b6d703d844af58da952ef0635e29b89da4af1d4a98ef690e8990d",
			},
		},
	}

	simulateRsp, aptosErr := client.SimulateTransaction(context.Background(), tx)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("result: %d\n", len(simulateRsp))
}
