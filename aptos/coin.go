package aptos

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptosmodule"
	"github.com/motoko9/aptos-go/crypto"
	"github.com/motoko9/aptos-go/rpcmodule"
	"strings"
)

const (
	AptosCoin = "Aptos"
	BTCCoin   = "BTC"
	USDTCoin  = "USDT"
	MOONCoin  = "MOON"
)

// only for devnet, mainnet is diffierent
// todo
var CoinType = map[string]string{
	"Aptos": "0x1::aptos_coin::AptosCoin",
	"BTC":   "0x43417434fd869edee76cca2a4d2301e528a1551b1d719b75c350c3c97d15b8b9::coins::BTC",
	"USDT":  "0x1685cdc9a83c3da34c59208f34bddb3217f63bfbe0c393f04462d1ba06465d08::usdt::USDT",
	"MOON":  "0xbb04c2079bc5611345689582eabab626732411b909045f8326d2b4980eac9b07::moon_coin::MoonCoin",
}

func AddressFromCoinType(coinType string) string {
	items := strings.Split(coinType, "::")
	if len(items) != 3 {
		return ""
	}
	return items[0]
}

func (cl *Client) CoinInfo(ctx context.Context, coin string, version uint64) (*aptosmodule.CoinInfo, *rpcmodule.AptosError) {
	coinType, ok := CoinType[coin]
	if !ok {
		return nil, &rpcmodule.AptosError{
			Message:     fmt.Sprintf("coin %s resouce is invalid", coin),
			ErrorCode:   "400",
			VmErrorCode: 0,
		}
	}
	//
	coinAddress := AddressFromCoinType(coinType)
	coinInfoResourceType := fmt.Sprintf("0x1::coin::CoinInfo<%s>", coinType)
	accountResource, err := cl.AccountResourceByAddressAndType(ctx, coinAddress, coinInfoResourceType, version)
	if err != nil {
		return nil, err
	}
	coinInfo, ok := accountResource.Object.(*aptosmodule.CoinInfo)
	if !ok {
		return nil, &rpcmodule.AptosError{
			Message:     fmt.Sprintf("coin %s resouce is invalid", coin),
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
			Message:     fmt.Sprintf("coin %s resouce is invalid", coin),
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
