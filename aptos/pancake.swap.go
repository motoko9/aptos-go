package aptos

import (
    "context"
    "fmt"
    "github.com/motoko9/aptos-go/crypto"
    "github.com/motoko9/aptos-go/rpcmodule"
)

const (
    PancakeSwapAccount = "0x83502b80f2e5792e18f979de2cc68cc15daefa94c88284351f769e0a49771cf9"
)

func (cl *Client) CreatePair(ctx context.Context, sender string, coin1 string, coin2 string, signer crypto.Signer) (string, *rpcmodule.AptosError) {
    payload := &rpcmodule.TransactionPayloadEntryFunctionPayload{
        Type:          rpcmodule.EntryFunctionPayload,
        Function:      fmt.Sprintf("%s::router::create_pair", PancakeSwapAccount),
        TypeArguments: []string{coin1, coin2},
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

func (cl *Client) AddLiquidity(ctx context.Context, sender string, coin1 string, amount1 uint64, coin2 string, amount2 uint64, signer crypto.Signer) (string, *rpcmodule.AptosError) {
    payload := &rpcmodule.TransactionPayloadEntryFunctionPayload{
        Type:          rpcmodule.EntryFunctionPayload,
        Function:      fmt.Sprintf("%s::router::add_liquidity", PancakeSwapAccount),
        TypeArguments: []string{coin1, coin2},
        Arguments:     []interface{}{fmt.Sprintf("%d", amount1), fmt.Sprintf("%d", amount2), "0", "0"},
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

func (cl *Client) SwapExactInput(ctx context.Context, sender string, coin1 string, coin2 string, amount1 uint64, signer crypto.Signer) (string, *rpcmodule.AptosError) {
    payload := &rpcmodule.TransactionPayloadEntryFunctionPayload{
        Type:          rpcmodule.EntryFunctionPayload,
        Function:      fmt.Sprintf("%s::router::swap_exact_input", PancakeSwapAccount),
        TypeArguments: []string{coin1, coin2},
        Arguments:     []interface{}{fmt.Sprintf("%d", amount1), fmt.Sprintf("%d", 0)},
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
