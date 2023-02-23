package aptos

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptosmodule"
	"github.com/motoko9/aptos-go/crypto"
	"github.com/motoko9/aptos-go/rpcmodule"
)

const (
	AptosCoin = "APT"
	BTCCoin   = "BTC"
	USDTCoin  = "USDT"
	MOONCoin  = "MOON"
)

// mainnet is diffierent
// todo
var CoinType = map[string]string{
	"APT":  "0x1::aptos_coin::AptosCoin",
	"BTC":  "0x43417434fd869edee76cca2a4d2301e528a1551b1d719b75c350c3c97d15b8b9::coins::BTC",
	"USDT": "0xbeca0b2fd5f778302e405182e5c250e1f6648492d53e48f5b29446f61dbcc848::usdt::USDT",
	"MOON": "0xbb04c2079bc5611345689582eabab626732411b909045f8326d2b4980eac9b07::moon_coin::MoonCoin",
}

func (cl *Client) CoinInfo(ctx context.Context, coin string, version uint64) (*aptosmodule.CoinInfo, *rpcmodule.AptosError) {
	coinType, ok := CoinType[coin]
	if !ok {
		return nil, &rpcmodule.AptosError{
			Message:     fmt.Sprintf("token %s resouce is not supported", coin),
			ErrorCode:   "400",
			VmErrorCode: 0,
		}
	}
	//
	coinAddress, err := rpcmodule.ExtractAddressFromType(coinType)
	if err != nil {
		return nil, err
	}
	coinInfoResourceType := fmt.Sprintf("0x1::coin::CoinInfo<%s>", coinType)
	accountResource, err := cl.AccountResourceByAddressAndType(ctx, coinAddress, coinInfoResourceType, version)
	if err != nil {
		return nil, err
	}
	coinInfo, ok := accountResource.Object.(*aptosmodule.CoinInfo)
	if !ok {
		return nil, &rpcmodule.AptosError{
			Message:     fmt.Sprintf("token %s resouce is invalid", coin),
			ErrorCode:   "400",
			VmErrorCode: 0,
		}
	}
	return coinInfo, nil
}

func CoinInitializePayload(coinType string, name string, symbol string, decimal byte) *rpcmodule.TransactionPayload {
	payload := &rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      "0x1::managed_coin::initialize",
		TypeArguments: []string{coinType},
		Arguments:     []interface{}{name, symbol, decimal, true},
	}
	return &rpcmodule.TransactionPayload{
		Type:   rpcmodule.EntryFunctionPayload,
		Object: payload,
	}
}

func (cl *Client) InitializeCoin(ctx context.Context, addr string, coinType string, name string, symbol string, decimal byte, signer crypto.Signer) (string, *rpcmodule.AptosError) {
	accountFrom, err := cl.Account(ctx, addr, 0)
	if err != nil {
		return "", err
	}

	payload := CoinInitializePayload(coinType, name, symbol, decimal)
	return cl.SignAndSubmitTransaction(ctx, addr, accountFrom.SequenceNumber, payload, signer)
}

func MintCoinPayload(coinType string, recipient string, amount uint64) *rpcmodule.TransactionPayload {
	payload := &rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      "0x1::managed_coin::mint",
		TypeArguments: []string{coinType},
		Arguments:     []interface{}{recipient, fmt.Sprintf("%d", amount)},
	}
	return &rpcmodule.TransactionPayload{
		Type:   rpcmodule.EntryFunctionPayload,
		Object: payload,
	}
}

func (cl *Client) MintCoin(ctx context.Context, addr string, coinType string, recipient string, amount uint64, signer crypto.Signer) (string, *rpcmodule.AptosError) {
	accountFrom, err := cl.Account(ctx, addr, 0)
	if err != nil {
		return "", err
	}

	payload := MintCoinPayload(coinType, recipient, amount)
	return cl.SignAndSubmitTransaction(ctx, addr, accountFrom.SequenceNumber, payload, signer)
}

func RegisterRecipientPayload(coin string) (*rpcmodule.TransactionPayload, *rpcmodule.AptosError) {
	// transfer
	coin, ok := CoinType[coin]
	if !ok {
		return nil, &rpcmodule.AptosError{
			Message:     fmt.Sprintf("token %s resouce is not supported", coin),
			ErrorCode:   "400",
			VmErrorCode: 0,
		}
	}
	transferPayload := rpcmodule.TransactionPayloadEntryFunctionPayload{
		Type:          rpcmodule.EntryFunctionPayload,
		Function:      "0x1::managed_coin::register",
		TypeArguments: []string{coin},
		Arguments:     []interface{}{},
	}
	return &rpcmodule.TransactionPayload{
		Type:   rpcmodule.EntryFunctionPayload,
		Object: transferPayload,
	}, nil
}

func (cl *Client) RegisterRecipient(ctx context.Context, addr string, coin string, signer crypto.Signer) (string, *rpcmodule.AptosError) {
	// recipient account
	account, err := cl.Account(ctx, addr, 0)
	if err != nil {
		return "", err
	}

	payload, err := RegisterRecipientPayload(coin)
	if err != nil {
		return "", err
	}
	return cl.SignAndSubmitTransaction(ctx, addr, account.SequenceNumber, payload, signer)
}

func TransferCoinPayload(coin string, amount uint64, receipt string) (*rpcmodule.TransactionPayload, *rpcmodule.AptosError) {
	// transfer
	coin, ok := CoinType[coin]
	if !ok {
		return nil, &rpcmodule.AptosError{
			Message:     fmt.Sprintf("token %s resouce is invalid", coin),
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