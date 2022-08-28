package aptos

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/rpcmodule"
	"github.com/motoko9/aptos-go/utils"
	"time"
)

type Signer interface {
	Sign(data []byte) ([]byte, error)
	PublicKey() utils.PublicKey
}

func (cl *Client) TransactionPending(ctx context.Context, hash string) (bool, *rpcmodule.AptosError) {
	var transaction rpcmodule.Transaction
	err, aptosErr := cl.Get(ctx, "/transactions/by_hash/"+hash, nil, &transaction)
	if err != nil {
		return false, rpcmodule.AptosErrorFromError(err)
	}
	if aptosErr != nil {
		if aptosErr.ErrorCode == rpcmodule.TransactionNotFound {
			return true, nil
		}
	}
	// can get transaction
	if transaction.Type == rpcmodule.PendingTransaction {
		return true, nil
	} else {
		return false, nil
	}
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

func (cl *Client) PublishMoveModuleLegacyReq(addr string, sequenceNumber uint64, content []byte) (*rpcmodule.EncodeSubmissionRequest, error) {
	publishPayload := rpcmodule.TransactionPayloadModuleBundlePayload{
		Type: rpcmodule.ModuleBundlePayload,
		Modules: []rpcmodule.MoveModule{
			{
				ByteCode: "0x" + hex.EncodeToString(content),
			},
		},
	}
	return rpcmodule.EncodeSubmissionReq(addr, sequenceNumber,
		rpcmodule.TransactionPayload{
			Type:   rpcmodule.ModuleBundlePayload,
			Object: publishPayload,
		})
}

func (cl *Client) PublishMoveModuleReq(addr string, sequenceNumber uint64, content []byte, meta []byte) (*rpcmodule.EncodeSubmissionRequest, error) {
	publishPayload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:     rpcmodule.EntryFunctionPayload,
		Function: "0x1::code::publish_package_txn",
		Arguments: []interface{}{
			"0x" + hex.EncodeToString(meta),
			[]interface{}{
				"0x" + hex.EncodeToString(content),
			},
		},
		TypeArguments: []string{},
	}
	return rpcmodule.EncodeSubmissionReq(addr, sequenceNumber,
		rpcmodule.TransactionPayload{
			Type:   rpcmodule.EntryFunctionPayload,
			Object: publishPayload,
		})
}

func (cl *Client) TransferCoinReq(from string, sequenceNumber uint64, coin string, amount uint64, receipt string) (*rpcmodule.EncodeSubmissionRequest, error) {
	// transfer
	coin, ok := CoinType[coin]
	if !ok {
		return nil, fmt.Errorf("coin %s is not supported", coin)
	}
	transferPayload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      "0x1::coin::transfer",
		Arguments:     []interface{}{receipt, fmt.Sprintf("%d", amount)},
		TypeArguments: []string{coin},
	}
	return rpcmodule.EncodeSubmissionReq(from, sequenceNumber,
		rpcmodule.TransactionPayload{
			Type:   rpcmodule.EntryFunctionPayload,
			Object: transferPayload,
		})
}

func (cl *Client) RegisterRecipientReq(from string, sequenceNumber uint64, coin string) (*rpcmodule.EncodeSubmissionRequest, error) {
	// transfer
	coin, ok := CoinType[coin]
	if !ok {
		return nil, fmt.Errorf("coin %s is not supported", coin)
	}
	transferPayload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      "0x1::managed_coin::register",
		Arguments:     []interface{}{},
		TypeArguments: []string{coin},
	}
	return rpcmodule.EncodeSubmissionReq(from, sequenceNumber,
		rpcmodule.TransactionPayload{
			Type:   rpcmodule.EntryFunctionPayload,
			Object: transferPayload,
		})
}

func (cl *Client) TransferCoin(ctx context.Context, from string, coin string, amount uint64, receipt string, signer Signer) (string, *rpcmodule.AptosError) {
	// from account
	accountFrom, aptosErr := cl.Account(ctx, from, 0)
	if aptosErr != nil {
		return "", aptosErr
	}

	encodeSubmissionReq, err := cl.TransferCoinReq(from, accountFrom.SequenceNumber, coin, amount, receipt)
	if err != nil {
		return "", rpcmodule.AptosErrorFromError(err)
	}

	// sign message
	signData, aptosErr := cl.EncodeSubmission(ctx, encodeSubmissionReq)
	if aptosErr != nil {
		return "", aptosErr
	}

	// sign
	signature, err := signer.Sign(signData)
	if err != nil {
		return "", rpcmodule.AptosErrorFromError(err)
	}

	// add signature
	submitReq, err := rpcmodule.SubmitTransactionReq(encodeSubmissionReq, rpcmodule.AccountSignature{
		Type: "ed25519_signature",
		Object: rpcmodule.AccountSignatureEd25519Signature{
			Type:      "ed25519_signature",
			PublicKey: "0x" + signer.PublicKey().String(),
			Signature: "0x" + hex.EncodeToString(signature),
		},
	})
	if err != nil {
		return "", rpcmodule.AptosErrorFromError(err)
	}

	// submit
	txHash, aptosErr := cl.SubmitTransaction(ctx, submitReq)
	if aptosErr != nil {
		return "", aptosErr
	}
	//
	return txHash, nil
}

// PublishMoveModuleLegacy
// can publish move module with batch
// do not working
func (cl *Client) PublishMoveModuleLegacy(ctx context.Context, addr string, content []byte, signer Signer) (string, *rpcmodule.AptosError) {
	// from account
	account, aptosErr := cl.Account(ctx, addr, 0)
	if aptosErr != nil {
		return "", aptosErr
	}

	// publish message
	encodeSubmissionReq, err := cl.PublishMoveModuleLegacyReq(addr, account.SequenceNumber, content)
	if err != nil {
		return "", rpcmodule.AptosErrorFromError(err)
	}

	// sign message
	signData, aptosErr := cl.EncodeSubmission(ctx, encodeSubmissionReq)
	if aptosErr != nil {
		return "", aptosErr
	}

	// sign
	signature, err := signer.Sign(signData)
	if err != nil {
		return "", rpcmodule.AptosErrorFromError(err)
	}

	// add signature
	submitReq, err := rpcmodule.SubmitTransactionReq(encodeSubmissionReq, rpcmodule.AccountSignature{
		Type: "ed25519_signature",
		Object: rpcmodule.AccountSignatureEd25519Signature{
			Type:      "ed25519_signature",
			PublicKey: "0x" + signer.PublicKey().String(),
			Signature: "0x" + hex.EncodeToString(signature),
		},
	})
	if err != nil {
		return "", rpcmodule.AptosErrorFromError(err)
	}

	// submit
	txHash, aptosErr := cl.SubmitTransaction(ctx, submitReq)
	if aptosErr != nil {
		return "", aptosErr
	}
	//
	return txHash, nil
}

// PublishMoveModule
// todo
// do not working
func (cl *Client) PublishMoveModule(ctx context.Context, addr string, content []byte, signer Signer) (string, *rpcmodule.AptosError) {
	// from account
	account, aptosErr := cl.Account(ctx, addr, 0)
	if aptosErr != nil {
		return "", aptosErr
	}

	// publish message
	// todo
	// how to get meta ?
	encodeSubmissionReq, err := cl.PublishMoveModuleReq(addr, account.SequenceNumber, content, content)
	if err != nil {
		return "", rpcmodule.AptosErrorFromError(err)
	}

	// sign message
	signData, aptosErr := cl.EncodeSubmission(ctx, encodeSubmissionReq)
	if aptosErr != nil {
		return "", aptosErr
	}

	// sign
	signature, err := signer.Sign(signData)
	if err != nil {
		return "", rpcmodule.AptosErrorFromError(err)
	}

	// add signature
	submitReq, err := rpcmodule.SubmitTransactionReq(encodeSubmissionReq, rpcmodule.AccountSignature{
		Type: "ed25519_signature",
		Object: rpcmodule.AccountSignatureEd25519Signature{
			Type:      "ed25519_signature",
			PublicKey: "0x" + signer.PublicKey().String(),
			Signature: "0x" + hex.EncodeToString(signature),
		},
	})
	if err != nil {
		return "", rpcmodule.AptosErrorFromError(err)
	}

	// submit
	txHash, aptosErr := cl.SubmitTransaction(ctx, submitReq)
	if aptosErr != nil {
		return "", aptosErr
	}
	//
	return txHash, nil
}

func (cl *Client) RegisterRecipient(ctx context.Context, addr string, coin string, signer Signer) (string, *rpcmodule.AptosError) {
	// recipient account
	account, aptosErr := cl.Account(ctx, addr, 0)
	if aptosErr != nil {
		return "", aptosErr
	}

	encodeSubmissionReq, err := cl.RegisterRecipientReq(addr, account.SequenceNumber, coin)
	if err != nil {
		return "", rpcmodule.AptosErrorFromError(err)
	}

	// sign message
	signData, aptosErr := cl.EncodeSubmission(ctx, encodeSubmissionReq)
	if aptosErr != nil {
		return "", aptosErr
	}

	// sign
	signature, err := signer.Sign(signData)
	if err != nil {
		return "", rpcmodule.AptosErrorFromError(err)
	}

	// add signature
	submitReq, err := rpcmodule.SubmitTransactionReq(encodeSubmissionReq, rpcmodule.AccountSignature{
		Type: "ed25519_signature",
		Object: rpcmodule.AccountSignatureEd25519Signature{
			Type:      "ed25519_signature",
			PublicKey: "0x" + signer.PublicKey().String(),
			Signature: "0x" + hex.EncodeToString(signature),
		},
	})
	if err != nil {
		return "", rpcmodule.AptosErrorFromError(err)
	}

	// submit
	txHash, aptosErr := cl.SubmitTransaction(ctx, submitReq)
	if aptosErr != nil {
		return "", aptosErr
	}
	return txHash, nil
}
