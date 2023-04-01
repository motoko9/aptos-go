package aptos

import (
    "context"
    "fmt"
    "github.com/motoko9/aptos-go/crypto"
    "github.com/motoko9/aptos-go/rpcmodule"
)

const (
    LiquidSwapAccount = "0x190d44266241744264b964a37b8f09863167a12d3e70cda39376cfb4e3561e12"
)

func (cl *Client) LiquidCreatePair(ctx context.Context, sender string, coin1 string, coin2 string, stable bool, signer crypto.Signer) (string, *rpcmodule.AptosError) {
    curve := "0x190d44266241744264b964a37b8f09863167a12d3e70cda39376cfb4e3561e12::curves::Uncorrelated"
    if stable {
        curve = "0x190d44266241744264b964a37b8f09863167a12d3e70cda39376cfb4e3561e12::curves::Stable"
    }
    payload := &rpcmodule.TransactionPayloadEntryFunctionPayload{
        Type:          rpcmodule.EntryFunctionPayload,
        Function:      fmt.Sprintf("%s::scripts::register_pool", LiquidSwapAccount),
        TypeArguments: []string{coin1, coin2, curve},
        Arguments:     []interface{}{},
    }
    payload1 := &rpcmodule.TransactionPayload{
        Type:   rpcmodule.EntryFunctionPayload,
        Object: payload,
    }
    account, aptosErr := cl.Account(ctx, sender, 0)
    if aptosErr != nil {
        return "", aptosErr
    }
    return cl.SignAndSubmitTransaction(ctx, sender, account.SequenceNumber, payload1, signer)
}

func (cl *Client) LiquidAddLiquidity(ctx context.Context, sender string, coin1 string, amount1 uint64, coin2 string, amount2 uint64, stable bool, signer crypto.Signer) (string, *rpcmodule.AptosError) {
    curve := "0x190d44266241744264b964a37b8f09863167a12d3e70cda39376cfb4e3561e12::curves::Uncorrelated"
    if stable {
        curve = "0x190d44266241744264b964a37b8f09863167a12d3e70cda39376cfb4e3561e12::curves::Stable"
    }
    payload := &rpcmodule.TransactionPayloadEntryFunctionPayload{
        Type:          rpcmodule.EntryFunctionPayload,
        Function:      fmt.Sprintf("%s::scripts::add_liquidity", LiquidSwapAccount),
        TypeArguments: []string{coin1, coin2, curve},
        Arguments:     []interface{}{fmt.Sprintf("%d", amount1), "0", fmt.Sprintf("%d", amount2), "0"},
    }
    payload1 := &rpcmodule.TransactionPayload{
        Type:   rpcmodule.EntryFunctionPayload,
        Object: payload,
    }
    account, aptosErr := cl.Account(ctx, sender, 0)
    if aptosErr != nil {
        return "", aptosErr
    }
    return cl.SignAndSubmitTransaction(ctx, sender, account.SequenceNumber, payload1, signer)
}
