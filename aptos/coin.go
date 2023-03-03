package aptos

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/motoko9/aptos-go/aptosmodule"
	"github.com/motoko9/aptos-go/crypto"
	"github.com/motoko9/aptos-go/rpcmodule"
	"github.com/motoko9/aptos-go/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

type CoinInfo struct {
	Source   string `json:"source"`
	ChainId  int    `json:"chainId"`
	Name     string `json:"name"`
	Decimals int    `json:"decimals"`
	Symbol   string `json:"symbol"`
	T        string `json:"type"`
}

//
func readCoinFromPontemNetwork() ([]*CoinInfo, error) {
	url := "https://raw.githubusercontent.com/pontem-network/coins-registry/main/src/coins.json"
	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	//
	type PontemNetworkCoinInfo struct {
		Source   string `json:"source"`
		ChainId  int    `json:"chainId"`
		Name     string `json:"name"`
		Decimals int    `json:"decimals"`
		Symbol   string `json:"symbol"`
		T        string `json:"type"`
	}

	pontemNetworkCoins := make([]*PontemNetworkCoinInfo, 0)
	err = json.Unmarshal(data, &pontemNetworkCoins)
	if err != nil {
		return nil, err
	}
	coins := make([]*CoinInfo, 0)
	for _, pontemNetworkCoin := range pontemNetworkCoins {
		coins = append(coins, &CoinInfo{
			Source:   pontemNetworkCoin.Source,
			ChainId:  pontemNetworkCoin.ChainId,
			Name:     pontemNetworkCoin.Name,
			Decimals: pontemNetworkCoin.Decimals,
			Symbol:   pontemNetworkCoin.Symbol,
			T:        pontemNetworkCoin.T,
		})
	}
	return coins, nil
}

func readCoinFromPancakeSwap() ([]*CoinInfo, error) {
	url := "https://raw.githubusercontent.com/pancakeswap/token-list/main/src/tokens/pancakeswap-aptos.json"
	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	//
	type PancakeSwapCoinInfo struct {
		ChainId  int    `json:"chainId"`
		Name     string `json:"name"`
		Decimals int    `json:"decimals"`
		Symbol   string `json:"symbol"`
		Address  string `json:"address"`
	}

	pancakeSwapCoins := make([]*PancakeSwapCoinInfo, 0)
	err = json.Unmarshal(data, &pancakeSwapCoins)
	if err != nil {
		return nil, err
	}
	coins := make([]*CoinInfo, 0)
	for _, pancakeSwapCoin := range pancakeSwapCoins {
		coins = append(coins, &CoinInfo{
			Source:   "pancake",
			ChainId:  pancakeSwapCoin.ChainId,
			Name:     pancakeSwapCoin.Name,
			Decimals: pancakeSwapCoin.Decimals,
			Symbol:   pancakeSwapCoin.Symbol,
			T:        pancakeSwapCoin.Address,
		})
	}
	return coins, nil
}

func CoinAlias(symbol string, source string) string {
	return fmt.Sprintf("%s(%s)", symbol, source)
}

func CoinSymbolSource(alias string) (string, string) {
	symbol, source := "", ""
	index1 := strings.IndexByte(alias, '(')
	index2 := strings.IndexByte(alias, ')')
	if index1 == -1 || index2 == -1 {
		return symbol, source
	}
	symbol = alias[0:index1]
	source = alias[index1+1:index2]
	return symbol, source
}

func (cl *Client) FindCoinBySymbolSource(symbol string, source string) *CoinInfo {
	for _, item := range cl.coinType {
		if item.Symbol == symbol && item.Source == source {
			return item
		}
	}
	return nil
}

func (cl *Client) FindCoinByType(t string) *CoinInfo {
	for _, item := range cl.coinType {
		if item.T == t {
			return item
		}
	}
	return nil
}

func (cl *Client) TryParseCoinType(coin string) string {
	if utils.IsCoinType(coin) {
		return coin
	}
	symbol, source := CoinSymbolSource(coin)
	coinInfo := cl.FindCoinBySymbolSource(symbol, source)
	if coinInfo == nil {
		return ""
	}
	return coinInfo.T
}

func (cl *Client) GetCoinNameByType(t string) string {
	if !utils.IsCoinType(t) {
		return ""
	}
	coinInfo := cl.FindCoinByType(t)
	if coinInfo == nil {
		return ""
	}
	return CoinAlias(coinInfo.Symbol, coinInfo.Source)
}

func (cl *Client) CustomizeCoin(source string, name string, decimals int, symbol string, t string) {
	cl.coinType[t] = &CoinInfo{
		Source:   source,
		Name:     name,
		Decimals: decimals,
		Symbol:   symbol,
		T:        t,
	}
}

func (cl *Client) AllCoins() []*CoinInfo {
	coins := make([]*CoinInfo, 0)
	for _, item := range cl.coinType {
		coins = append(coins, item)
	}
	return coins
}

func (cl *Client) CoinInfo(ctx context.Context, coin string, version uint64) (*aptosmodule.CoinInfo, *rpcmodule.AptosError) {
	coinType := cl.TryParseCoinType(coin)
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
	//
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

	coinType := cl.TryParseCoinType(coin)
	if coinType == "" {
		return "", &rpcmodule.AptosError{
			Message:     fmt.Sprintf("coin %s resouce is not supported", coin),
			ErrorCode:   "400",
			VmErrorCode: 0,
		}
	}

	payload, err := RegisterRecipientPayload(coinType)
	if err != nil {
		return "", err
	}
	return cl.SignAndSubmitTransaction(ctx, addr, account.SequenceNumber, payload, signer)
}

func TransferCoinPayload(coin string, amount uint64, receipt string) (*rpcmodule.TransactionPayload, *rpcmodule.AptosError) {
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

	coinType := cl.TryParseCoinType(coin)
	if coinType == "" {
		return "", &rpcmodule.AptosError{
			Message:     fmt.Sprintf("coin %s resouce is not supported", coin),
			ErrorCode:   "400",
			VmErrorCode: 0,
		}
	}

	payload, err := TransferCoinPayload(coin, amount, receipt)
	if err != nil {
		return "", err
	}

	return cl.SignAndSubmitTransaction(ctx, from, accountFrom.SequenceNumber, payload, signer)
}
