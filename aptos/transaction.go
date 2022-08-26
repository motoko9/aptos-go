package aptos

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/utils"
	"time"
)

type Signer interface {
	Sign(data []byte) ([]byte, error)
	PublicKey() utils.PublicKey
}

func (cl *Client) TransactionPending(ctx context.Context, hash string) (bool, error) {
	var transaction rpc.Transaction
	code, err := cl.Get(ctx, "/transactions/by_hash/"+hash, nil, &transaction)
	if code == -1 {
		return false, err
	}
	if code == 404 {
		// resource not found, maybe transaction is not on chain
		return true, nil
	}
	if code == 200 {
		if transaction.T == "pending_transaction" {
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

func (cl *Client) PublishMoveModuleMsg(addr string, sequenceNumber uint64, module []byte) (*rpc.Transaction, error) {
	publishPayload := rpc.ModuleBundlePayload{
		T: "module_bundle_payload",
		Modules: []rpc.Module{
			{
				ByteCode: "0x" + hex.EncodeToString(module),
			},
		},
	}
	publish := rpc.Transaction{
		T:                       "",
		Hash:                    "",
		Sender:                  addr,
		SequenceNumber:          sequenceNumber,
		MaxGasAmount:            uint64(2000),
		GasUnitPrice:            uint64(2),
		GasCurrencyCode:         "",
		ExpirationTimestampSecs: uint64(time.Now().Unix() + 600), // now + 10 minutes
		Payload:                 &publishPayload,
		Signature:               nil,
	}
	return &publish, nil
}

func (cl *Client) TransferCoinMsg(from string, sequenceNumber uint64, coin string, amount uint64, receipt string) (*rpc.Transaction, error) {
	// transfer
	coin, ok := CoinType[coin]
	if !ok {
		return nil, fmt.Errorf("coin %s is not supported", coin)
	}
	transferPayload := rpc.EntryFunctionPayload{
		Function:      "0x1::coin::transfer",
		Arguments:     []interface{}{receipt, fmt.Sprintf("%d", amount)},
		T:             "entry_function_payload",
		TypeArguments: []string{coin},
	}
	transaction := rpc.Transaction{
		T:                       "",
		Hash:                    "",
		Sender:                  from,
		SequenceNumber:          sequenceNumber,
		MaxGasAmount:            uint64(2000),
		GasUnitPrice:            uint64(1),
		GasCurrencyCode:         "",
		ExpirationTimestampSecs: uint64(time.Now().Unix() + 600), // now + 10 minutes
		Payload:                 &transferPayload,
		Signature:               nil,
	}
	return &transaction, nil
}

func (cl *Client) RegisterRecipientMsg(from string, sequenceNumber uint64, coin string) (*rpc.Transaction, error) {
	// transfer
	coin, ok := CoinType[coin]
	if !ok {
		return nil, fmt.Errorf("coin %s is not supported", coin)
	}
	transferPayload := rpc.EntryFunctionPayload{
		Function:      "0x1::coins::register",
		Arguments:     []interface{}{},
		T:             "entry_function_payload",
		TypeArguments: []string{coin},
	}
	transaction := rpc.Transaction{
		T:                       "",
		Hash:                    "",
		Sender:                  from,
		SequenceNumber:          sequenceNumber,
		MaxGasAmount:            uint64(2000),
		GasUnitPrice:            uint64(1),
		GasCurrencyCode:         "",
		ExpirationTimestampSecs: uint64(time.Now().Unix() + 600), // now + 10 minutes
		Payload:                 &transferPayload,
		Signature:               nil,
	}
	return &transaction, nil
}

func (cl *Client) TransferCoin(ctx context.Context, from string, coin string, amount uint64, receipt string, signer Signer) (*rpc.Transaction, error) {
	// from account
	accountFrom, err := cl.Account(ctx, from, 0)
	if err != nil {
		return nil, err
	}

	transaction, err := cl.TransferCoinMsg(from, accountFrom.SequenceNumber, coin, amount, receipt)
	if err != nil {
		return nil, err
	}

	// sign message
	signData, err := cl.EncodeSubmission(ctx, transaction)
	if err != nil {
		return nil, err
	}

	// sign
	signature, err := signer.Sign(signData)
	if err != nil {
		return nil, err
	}

	// add signature
	transaction.Signature = &rpc.Signature{
		T: "ed25519_signature",
		//PublicKey: fromAccount.AuthenticationKey,
		PublicKey: "0x" + signer.PublicKey().String(),
		Signature: "0x" + hex.EncodeToString(signature),
	}

	// submit
	tx, err := cl.SubmitTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}
	//
	return tx, nil
}

func (cl *Client) PublishMoveModule(ctx context.Context, addr string, module []byte, signer Signer) (*rpc.Transaction, error) {
	// from account
	account, err := cl.Account(ctx, addr, 0)
	if err != nil {
		return nil, err
	}

	// publish message
	transaction, err := cl.PublishMoveModuleMsg(addr, account.SequenceNumber, module)
	if err != nil {
		return nil, err
	}

	// sign message
	signData, err := cl.EncodeSubmission(ctx, transaction)
	if err != nil {
		return nil, err
	}

	// sign
	signature, err := signer.Sign(signData)
	if err != nil {
		return nil, err
	}

	// add signature
	transaction.Signature = &rpc.Signature{
		T: "ed25519_signature",
		//PublicKey: fromAccount.AuthenticationKey,
		PublicKey: "0x" + signer.PublicKey().String(),
		Signature: "0x" + hex.EncodeToString(signature),
	}

	// submit
	tx, err := cl.SubmitTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}
	//
	return tx, nil
}
