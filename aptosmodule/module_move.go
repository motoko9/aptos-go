package aptosmodule

type Guid struct {
	Id struct {
		Addr        string `json:"addr"`
		CreationNum uint64 `json:"creation_num,string"`
	} `json:"id"`
}
type Coin struct {
	Value uint64 `json:"value,string"`
}

type CoinEvents struct {
	Counter uint64 `json:"counter,string"`
	Guid    Guid   `json:"guid"`
}

type CoinStore struct {
	Coin           Coin       `json:"coin"`
	DepositEvents  CoinEvents `json:"deposit_events"`
	WithdrawEvents CoinEvents `json:"withdraw_events"`
}

type CoinInfo struct {
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Decimals uint64 `json:"decimals"`
	// todo
	// supply is not uint64
	// Supply uint64 `json:"supply"`
}
