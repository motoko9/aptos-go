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

func (cl *Client) ExeSimulateTransaction(ctx context.Context, sender string, sequence uint64, payload *rpcmodule.TransactionPayload) (rpcmodule.SimulateTransactionRsp, *rpcmodule.AptosError) {
	tx := &rpcmodule.SubmitTransactionRequest{
		Sender:                  sender,
		SequenceNumber:          sequence,
		MaxGasAmount:            uint64(80000),
		GasUnitPrice:            uint64(100),
		ExpirationTimestampSecs: uint64(time.Now().Unix() + 600),
		Payload:                 payload,
		Signature: rpcmodule.Signature{
			Type: rpcmodule.Ed25519Signature,
			Object: rpcmodule.SignatureEd25519Signature{
				Type:      rpcmodule.Ed25519Signature,
				PublicKey: "0xd36df53c46ca6c046648a85083482149a6423b06a3e35cd2f91b01d656c06e73",
				Signature: "0x945073c4b0d389271b1e6959e5238a1d6f6f82aa6dc09ca6cd31eb6952bde2781b6338d2470b6d703d844af58da952ef0635e29b89da4af1d4a98ef690e8990d",
			},
		},
	}
	return cl.SimulateTransaction(ctx, tx)
}
