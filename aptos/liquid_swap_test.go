package aptos

import (
    "context"
    "fmt"
    "github.com/motoko9/aptos-go/rpc"
    "github.com/motoko9/aptos-go/wallet"
    "testing"
)

func TestPancakeSwap_LiquidCreatePair(t *testing.T) {
    userWallet, err := wallet.NewFromKey("f476ba25a9df047f8d4c024896a171c60f32eb31b89bccbbbf1462b46e0475e3")
    if err != nil {
        panic(err)
    }
    fmt.Printf("user address: %s\n", userWallet.Address())

    client := New(rpc.TestNet_RPC, false)
    txHash, aptosErr := client.LiquidCreatePair(context.Background(), userWallet.Address(),
        "0x2f88a12a17f01228f4ba72ec6214127abb930512dcb3d6205909ca510aca7b29::asset::USDC",
        "0x2f88a12a17f01228f4ba72ec6214127abb930512dcb3d6205909ca510aca7b29::asset::USDT",
        true,
        userWallet)
    if aptosErr != nil {
        panic(aptosErr)
    }
    fmt.Printf("create pair tx hash: %s\n", txHash)
}

func TestPancakeSwap_LiquidAddLiquidity(t *testing.T) {
    userWallet, err := wallet.NewFromKey("f476ba25a9df047f8d4c024896a171c60f32eb31b89bccbbbf1462b46e0475e3")
    if err != nil {
        panic(err)
    }
    fmt.Printf("user address: %s\n", userWallet.Address())

    client := New(rpc.TestNet_RPC, false)
    txHash, aptosErr := client.LiquidAddLiquidity(context.Background(), userWallet.Address(),
        "0x2f88a12a17f01228f4ba72ec6214127abb930512dcb3d6205909ca510aca7b29::asset::USDC",
        1000000000000,
        "0x2f88a12a17f01228f4ba72ec6214127abb930512dcb3d6205909ca510aca7b29::asset::USDT",
        1000000000000,
        true,
        userWallet)
    if aptosErr != nil {
        panic(aptosErr)
    }
    fmt.Printf("add liquidity tx hash: %s\n", txHash)
}
