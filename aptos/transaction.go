package aptos

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/crypto"
	"github.com/motoko9/aptos-go/rpcmodule"
	"time"
)

func (cl *Client) TransactionPending(ctx context.Context, hash string) (bool, *rpcmodule.AptosError) {
	tx, err := cl.TransactionByHash(ctx, hash)
	if err != nil {
		if err.ErrorCode == rpcmodule.TransactionNotFound {
			return true, nil
		}
		return false, err
	}

	return tx.Type == rpcmodule.PendingTransaction, nil
}

func (cl *Client) ConfirmTransaction(ctx context.Context, hash string) (bool, *rpcmodule.AptosError) {
	counter := 0
	for counter < 100 {
		pending, err := cl.TransactionPending(ctx, hash)
		if err != nil {
			return false, err
		}
		if !pending {
			return true, nil
		}
		counter++
		time.Sleep(time.Second * 1)
	}
	return false, nil
}

func TransferCoinPayload(coin string, amount uint64, receipt string) (*rpcmodule.TransactionPayload, *rpcmodule.AptosError) {
	// transfer
	coin, ok := CoinType[coin]
	if !ok {
		return nil, &rpcmodule.AptosError{
			Message:     fmt.Sprintf("coin %s resouce is invalid", coin),
			ErrorCode:   "400",
			VmErrorCode: 0,
		}
	}
	transferPayload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      "0x1::coin::transfer",
		Arguments:     []interface{}{receipt, fmt.Sprintf("%d", amount)},
		TypeArguments: []string{coin},
	}
	return &rpcmodule.TransactionPayload{
		Type:   rpcmodule.EntryFunctionPayload,
		Object: transferPayload,
	}, nil
}

func (cl *Client) TransferCoin(ctx context.Context, from string, coin string, amount uint64, receipt string, signer crypto.Signer) (string, *rpcmodule.AptosError) {
	accountFrom, err := cl.Account(ctx, from, 0)
	if err != nil {
		return "", err
	}

	payload, err := TransferCoinPayload(coin, amount, receipt)
	if err != nil {
		return "", err
	}

	return cl.SignAndSubmitTransaction(ctx, from, accountFrom.SequenceNumber, payload, signer)
}

func PublishMoveModuleLegacyPayload(content []byte) (*rpcmodule.TransactionPayload, *rpcmodule.AptosError) {
	publishPayload := rpcmodule.TransactionPayloadModuleBundlePayload{
		Type: rpcmodule.ModuleBundlePayload,
		Modules: []rpcmodule.MoveModule{
			{
				ByteCode: "0x" + hex.EncodeToString(content),
			},
		},
	}
	return &rpcmodule.TransactionPayload{
		Type:   rpcmodule.ModuleBundlePayload,
		Object: publishPayload,
	}, nil
}

// PublishMoveModuleLegacy
// can publish move module with batch
// do not working
// return with hash, but can not find tx in explorer
func (cl *Client) PublishMoveModuleLegacy(ctx context.Context, addr string, content []byte, signer crypto.Signer) (string, *rpcmodule.AptosError) {
	// from account
	account, err := cl.Account(ctx, addr, 0)
	if err != nil {
		return "", err
	}

	// publish message
	payload, err := PublishMoveModuleLegacyPayload(content)
	if err != nil {
		return "", err
	}

	return cl.SignAndSubmitTransaction(ctx, addr, account.SequenceNumber, payload, signer)
}

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

func (cl *Client) PublishMoveModule(ctx context.Context, addr string, content []byte, signer crypto.Signer) (string, *rpcmodule.AptosError) {
	// from account
	account, err := cl.Account(ctx, addr, 0)
	if err != nil {
		return "", err
	}

	// says from  https://github.com/aptos-labs/aptos-core/blob/06b946df79889a1ac19f13aa336f8c069603345b/ecosystem/typescript/sdk/src/aptos_client.ts#L619
	// * Publishes a move package. `packageMetadata` and `modules` can be generated with command
	// * `aptos move compile --save-metadata [ --included-artifacts=<...> ]`.
	// did not work
	// publish message
	// todo
	// how to get meta ?
	payload, err := PublishMoveModulePayload(content, content)
	if err != nil {
		return "", err
	}

	return cl.SignAndSubmitTransaction(ctx, addr, account.SequenceNumber, payload, signer)
}

func (cl *Client) SignAndSubmitTransaction(ctx context.Context, sender string, sequence uint64, payload *rpcmodule.TransactionPayload, signer crypto.Signer) (string, *rpcmodule.AptosError) {
	encodeSubmissionReq := rpcmodule.EncodeSubmissionReq(sender, sequence, payload)

	// sign message
	signData, err := cl.EncodeSubmission(ctx, encodeSubmissionReq)
	if err != nil {
		return "", err
	}

	// sign
	signature, err1 := signer.Sign(signData)
	if err1 != nil {
		return "", rpcmodule.AptosErrorFromError(err1)
	}

	// add signature
	submitReq := rpcmodule.SubmitTransactionReq(encodeSubmissionReq, rpcmodule.Signature{
		Type: rpcmodule.Ed25519Signature,
		Object: rpcmodule.SignatureEd25519Signature{
			Type:      rpcmodule.Ed25519Signature,
			PublicKey: "0x" + signer.PublicKey().String(),
			Signature: "0x" + hex.EncodeToString(signature),
		},
	})

	// submit
	return cl.SubmitTransaction(ctx, submitReq)
}
