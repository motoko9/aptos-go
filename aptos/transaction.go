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

func (cl *Client) TransactionPending(ctx context.Context, hash string) (bool, error) {
	var transaction rpcmodule.Transaction
	code, err := cl.Get(ctx, "/transactions/by_hash/"+hash, nil, &transaction)
	if code == -1 {
		return false, err
	}
	if code == 404 {
		// resource not found, maybe transaction is not on chain
		return true, nil
	}
	if code == 200 {
		if transaction.Type == "pending_transaction" {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, err
}

func (cl *Client) ConfirmTransaction(ctx context.Context, hash string) (bool, error) {
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
		Type: "module_bundle_payload",
		Modules: []rpcmodule.MoveModule{
			{
				ByteCode: "0x" + hex.EncodeToString(content),
			},
		},
	}
	return rpcmodule.EncodeSubmissionReq(addr, sequenceNumber,
		rpcmodule.TransactionPayload{
			Type:   "module_bundle_payload",
			Object: publishPayload,
		})
}

func (cl *Client) PublishMoveModuleReq(addr string, sequenceNumber uint64, content []byte, meta []byte) (*rpcmodule.EncodeSubmissionRequest, error) {
	publishPayload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Function: "0x1::code::publish_package_txn",
		Arguments: []interface{}{
			"0x" + hex.EncodeToString(meta),
			[]interface{}{
				"0x" + hex.EncodeToString(content),
			},
		},
		Type:          "entry_function_payload",
		TypeArguments: []string{},
	}
	return rpcmodule.EncodeSubmissionReq(addr, sequenceNumber,
		rpcmodule.TransactionPayload{
			Type:   "entry_function_payload",
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
		Function:      "0x1::coin::transfer",
		Arguments:     []interface{}{receipt, fmt.Sprintf("%d", amount)},
		Type:          "entry_function_payload",
		TypeArguments: []string{coin},
	}
	return rpcmodule.EncodeSubmissionReq(from, sequenceNumber,
		rpcmodule.TransactionPayload{
			Type:   "entry_function_payload",
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
		Function:      "0x1::managed_coin::register",
		Arguments:     []interface{}{},
		Type:          "entry_function_payload",
		TypeArguments: []string{coin},
	}
	return rpcmodule.EncodeSubmissionReq(from, sequenceNumber,
		rpcmodule.TransactionPayload{
			Type:   "entry_function_payload",
			Object: transferPayload,
		})
}

func (cl *Client) TransferCoin(ctx context.Context, from string, coin string, amount uint64, receipt string, signer Signer) (string, error) {
	// from account
	accountFrom, err := cl.Account(ctx, from, 0)
	if err != nil {
		return "", err
	}

	encodeSubmissionReq, err := cl.TransferCoinReq(from, accountFrom.SequenceNumber, coin, amount, receipt)
	if err != nil {
		return "", err
	}

	// sign message
	signData, err := cl.EncodeSubmission(ctx, encodeSubmissionReq)
	if err != nil {
		return "", err
	}

	// sign
	signature, err := signer.Sign(signData)
	if err != nil {
		return "", err
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
		return "", err
	}

	// submit
	txHash, err := cl.SubmitTransaction(ctx, submitReq)
	if err != nil {
		return "", err
	}
	//
	return txHash, nil
}

// PublishMoveModuleLegacy
// can publish move module with batch
// do not working
func (cl *Client) PublishMoveModuleLegacy(ctx context.Context, addr string, content []byte, signer Signer) (string, error) {
	// from account
	account, err := cl.Account(ctx, addr, 0)
	if err != nil {
		return "", err
	}

	// publish message
	encodeSubmissionReq, err := cl.PublishMoveModuleLegacyReq(addr, account.SequenceNumber, content)
	if err != nil {
		return "", err
	}

	// sign message
	signData, err := cl.EncodeSubmission(ctx, encodeSubmissionReq)
	if err != nil {
		return "", err
	}

	// sign
	signature, err := signer.Sign(signData)
	if err != nil {
		return "", err
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
		return "", err
	}

	// submit
	txHash, err := cl.SubmitTransaction(ctx, submitReq)
	if err != nil {
		return "", err
	}
	//
	return txHash, nil
}

// PublishMoveModule
// todo
// do not working
func (cl *Client) PublishMoveModule(ctx context.Context, addr string, content []byte, signer Signer) (string, error) {
	// from account
	account, err := cl.Account(ctx, addr, 0)
	if err != nil {
		return "", err
	}

	// publish message
	// todo
	// how to get meta ?
	encodeSubmissionReq, err := cl.PublishMoveModuleReq(addr, account.SequenceNumber, content, content)
	if err != nil {
		return "", err
	}

	// sign message
	signData, err := cl.EncodeSubmission(ctx, encodeSubmissionReq)
	if err != nil {
		return "", err
	}

	// sign
	signature, err := signer.Sign(signData)
	if err != nil {
		return "", err
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
		return "", err
	}

	// submit
	txHash, err := cl.SubmitTransaction(ctx, submitReq)
	if err != nil {
		return "", err
	}
	//
	return txHash, nil
}

func (cl *Client) RegisterRecipient(ctx context.Context, addr string, coin string, signer Signer) (string, error) {
	// recipient account
	account, err := cl.Account(ctx, addr, 0)
	if err != nil {
		return "", err
	}

	encodeSubmissionReq, err := cl.RegisterRecipientReq(addr, account.SequenceNumber, coin)
	if err != nil {
		return "", err
	}

	// sign message
	signData, err := cl.EncodeSubmission(ctx, encodeSubmissionReq)
	if err != nil {
		return "", err
	}

	// sign
	signature, err := signer.Sign(signData)
	if err != nil {
		return "", err
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
		return "", err
	}

	// submit
	txHash, err := cl.SubmitTransaction(ctx, submitReq)
	if err != nil {
		return "", err
	}
	return txHash, nil
}