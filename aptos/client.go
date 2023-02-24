package aptos

import (
	"github.com/motoko9/aptos-go/rpc"
)

type Client struct {
	*rpc.Client
}

func New(endpoint string, mainNet bool) *Client {
	client := rpc.New(endpoint)
	if mainNet {
		CoinType[AptosCoin] = "0x1::aptos_coin::AptosCoin"
		CoinType[USDTCoin] = "0xf22bede237a07e121b56d91a491eb7bcdfd1f5907926a9e58338f964a01b17fa::asset::USDT"
		CoinType[BTCCoin] = "0xae478ff7d83ed072dbc5e264250e67ef58f57c99d89b447efd8a0a2e8b2be76e::coin::T"
		CoinType[ETHCoin] = "0xcc8a89c8dce9693d354449f1f73e60e14e347417854f029db5bc8e7454008abb::coin::T"
		CoinType[USDCCoin] = "0xf22bede237a07e121b56d91a491eb7bcdfd1f5907926a9e58338f964a01b17fa::asset::USDC"
	} else { // for testnet
		CoinType[AptosCoin] = "0x1::aptos_coin::AptosCoin"
		CoinType[USDTCoin] = "0xbeca0b2fd5f778302e405182e5c250e1f6648492d53e48f5b29446f61dbcc848::usdt::USDT"
	}
	return &Client{
		client,
	}
}
