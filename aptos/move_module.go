package aptos

import (
	"context"
	"encoding/hex"
	"github.com/motoko9/aptos-go/crypto"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func PublishMoveModulePayload(content []byte, meta []byte) (*rpcmodule.TransactionPayload, *rpcmodule.AptosError) {
	publishPayload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      "0x1::code::publish_package_txn",
		TypeArguments: []string{},
		Arguments: []interface{}{
			hex.EncodeToString(meta),
			[]interface{}{
				hex.EncodeToString(content),
			},
		},
	}
	return &rpcmodule.TransactionPayload{
		Type:   rpcmodule.EntryFunctionPayload,
		Object: publishPayload,
	}, nil
}

func (cl *Client) PublishMoveModule(ctx context.Context, addr string, content []byte, meta []byte, signer crypto.Signer) (string, *rpcmodule.AptosError) {
	// from account
	account, err := cl.Account(ctx, addr, 0)
	if err != nil {
		return "", err
	}

	// says from  https://github.com/aptos-labs/aptos-core/blob/06b946df79889a1ac19f13aa336f8c069603345b/ecosystem/typescript/sdk/src/aptos_client.ts#L619
	// * Publishes a move package. `packageMetadata` and `modules` can be generated with command
	// * `aptos move compile --save-metadata [ --included-artifacts=<...> ]`.
	// publish message
	//
	payload, err := PublishMoveModulePayload(content, meta)
	if err != nil {
		return "", err
	}

	return cl.SignAndSubmitTransaction(ctx, addr, account.SequenceNumber, payload, signer)
}
