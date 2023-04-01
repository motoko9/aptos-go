package aptos

import (
    "context"
    "fmt"
    "github.com/motoko9/aptos-go/crypto"
    "github.com/motoko9/aptos-go/rpcmodule"
)

const (
    AuxAccount = "0x8b7311d78d47e37d09435b8dc37c14afd977c5cfa74f974d45f0258d986eef53"
)

func (cl *Client) AuxCreatePair(ctx context.Context, sender string, coin1 string, coin2 string, fee uint64, signer crypto.Signer) (string, *rpcmodule.AptosError) {
    payload := &rpcmodule.TransactionPayloadEntryFunctionPayload{
        Type:          rpcmodule.EntryFunctionPayload,
        Function:      fmt.Sprintf("%s::amm::create_pool", AuxAccount),
        TypeArguments: []string{coin1, coin2},
        Arguments:     []interface{}{fmt.Sprintf("%d", fee)},
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

func (cl *Client) AuxAmmLiquidity(ctx context.Context, sender string, coin1 string, amount1 uint64, coin2 string, amount2 uint64, slippage uint64, signer crypto.Signer) (string, *rpcmodule.AptosError) {
    payload := &rpcmodule.TransactionPayloadEntryFunctionPayload{
        Type:          rpcmodule.EntryFunctionPayload,
        Function:      fmt.Sprintf("%s::amm::add_liquidity", AuxAccount),
        TypeArguments: []string{coin1, coin2},
        Arguments:     []interface{}{fmt.Sprintf("%d", amount1), fmt.Sprintf("%d", amount2), fmt.Sprintf("%d", slippage)},
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
