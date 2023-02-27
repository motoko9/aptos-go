package aptos

import (
	"context"
	"fmt"
	"github.com/motoko9/aptos-go/aptosmodule"
	"github.com/motoko9/aptos-go/crypto"
	"github.com/motoko9/aptos-go/rpcmodule"
	"github.com/motoko9/aptos-go/utils"
)

const (
	AptosCoin = "APT"
	WBTCCoin   = "WBTC"
	USDTCoin  = "USDT"
	WUSDTCoin  = "WUSDT"
	WETHCoin   = "WETHCoin"
	USDCCoin  = "USDC"
	WUSDCCoin  = "WUSDC"
	WSOLCoin  = "WSOL"
	MOONCoin  = "MOON"
)

var CoinType = map[string]string{}

func TryParseCoinType(coin string) string {
	if utils.IsCoinType(coin) {
		return coin
	}
	coinType, ok := CoinType[coin]
	if !ok {
		return ""
	}
	return coinType
}

func (cl *Client) GetCoinName(t string) (string, bool) {
	for k, v := range CoinType {
		if v == t {
			return k, true
		}
	}
	return "", false
}

func (cl *Client) CustomizeCoin(name string, t string) {
	CoinType[name] = t
}

func (cl *Client) CoinInfo(ctx context.Context, coin string, version uint64) (*aptosmodule.CoinInfo, *rpcmodule.AptosError) {
	coinType := TryParseCoinType(coin)
	if coinType == "" {
		return nil, &rpcmodule.AptosError{
			Message:     fmt.Sprintf("coin %s resouce is not supported", coin),
			ErrorCode:   "400",
			VmErrorCode: 0,
		}
	}
	//
	coinAddress, err := utils.ExtractAddressFromType(coinType)
	if err != nil {
		return nil, rpcmodule.AptosErrorFromError(err)
	}
	coinInfoResourceType := fmt.Sprintf("0x1::coin::CoinInfo<%s>", coinType)
	accountResource, aptosErr := cl.AccountResourceByAddressAndType(ctx, coinAddress, coinInfoResourceType, version)
	if aptosErr != nil {
		return nil, aptosErr
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
	coinType := TryParseCoinType(coin)
	if coinType == "" {
		return nil, &rpcmodule.AptosError{
			Message:     fmt.Sprintf("coin %s resouce is not supported", coin),
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
	coinType := TryParseCoinType(coin)
	if coinType == "" {
		return nil, &rpcmodule.AptosError{
			Message:     fmt.Sprintf("coin %s resouce is not supported", coin),
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
