package aptos

import (
	"context"
	"encoding/hex"
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
