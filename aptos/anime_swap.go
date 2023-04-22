package aptos

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/crypto"
	"github.com/motoko9/aptos-go/rpcmodule"
)

const (
	AnimeSwapAccount = "0x16fe2df00ea7dde4a63409201f7f4e536bde7bb7335526a35d05111e68aa322c"
)

func (cl *Client) AnimeAddLiquidity(ctx context.Context, sender string, coin1 string, amount1 uint64, coin2 string, amount2 uint64, signer crypto.Signer) (string, *rpcmodule.AptosError) {
	payload := &rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      fmt.Sprintf("%s::AnimeSwapPoolV1f1::add_liquidity_entry", AnimeSwapAccount),
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
