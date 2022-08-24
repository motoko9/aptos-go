package aptos

const (
	AptosCoin = "Aptos"
	BTCCoin   = "BTC"
	USDTCoin  = "USDT"
)

// only for devnet, mainnet is diffierent
// todo
var CoinType = map[string]string{
	"Aptos": "0x1::aptos_coin::AptosCoin",
	"BTC":   "0x43417434fd869edee76cca2a4d2301e528a1551b1d719b75c350c3c97d15b8b9::coins::BTC",
	"USDT":  "0x43417434fd869edee76cca2a4d2301e528a1551b1d719b75c350c3c97d15b8b9::coins::USDT",
}
