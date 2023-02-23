package aptos

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptosmodule"
	"github.com/motoko9/aptos-go/crypto"
	"github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) AccountBalance(ctx context.Context, address string, coin string, version uint64) (uint64, *rpcmodule.AptosError) {
	coin, ok := CoinType[coin]
	if !ok {
		return 0, &rpcmodule.AptosError{
			Message:     fmt.Sprintf("token %s is not supported", coin),
			ErrorCode:   "400",
			VmErrorCode: 0,
		}
	}

	resourceType := fmt.Sprintf("0x1::coin::CoinStore<%s>", coin)
	resource, err := cl.AccountResourceByAddressAndType(ctx, address, resourceType, version)
	if err != nil {
		// resource not found, so balance is zero
		if err.ErrorCode == rpcmodule.ResourceNotFound {
			return 0, nil
		}
		return 0, err
	}
	coinStore, ok := resource.Object.(*aptosmodule.CoinStore)
	if !ok {
		return 0, &rpcmodule.AptosError{
			Message:     fmt.Sprintf("address %s resouce is invalid", address),
			ErrorCode:   "400",
			VmErrorCode: 0,
		}
	}
	return coinStore.Coin.Value, nil
}

func CreateAccountPayload(newAccount string) (*rpcmodule.TransactionPayload, *rpcmodule.AptosError) {
	publishPayload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      "0x1::aptos_account::create_account",
		TypeArguments: []string{},
		Arguments: []interface{}{newAccount},
	}
	return &rpcmodule.TransactionPayload{
		Type:   rpcmodule.EntryFunctionPayload,
		Object: publishPayload,
	}, nil
}

func (cl *Client) CreateAccount(ctx context.Context, addr string, newAccount string, signer crypto.Signer) (string, *rpcmodule.AptosError) {
	// from account
	account, err := cl.Account(ctx, addr, 0)
	if err != nil {
		return "", err
	}

	payload, err := CreateAccountPayload(newAccount)
	if err != nil {
		return "", err
	}

	return cl.SignAndSubmitTransaction(ctx, addr, account.SequenceNumber, payload, signer)
}