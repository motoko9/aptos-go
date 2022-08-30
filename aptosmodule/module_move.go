package aptosmodule

type Guid struct {
	Id struct {
		Addr        string `json:"addr"`
		CreationNum uint64 `json:"creation_num,string"`
	} `json:"id"`
}

type Events struct {
	Counter uint64 `json:"counter,string"`
	Guid    Guid   `json:"guid"`
}

type Coin struct {
	Value uint64 `json:"value,string"`
}

type CoinStore struct {
	Coin           Coin   `json:"coin"`
	DepositEvents  Events `json:"deposit_events"`
	WithdrawEvents Events `json:"withdraw_events"`
}

func CoinStoreCreator() interface{} {
	return &CoinStore{}
}

type CoinInfo struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint64 `json:"decimals"`
	// todo
	// supply is not uint64
	// Supply uint64 `json:"supply"`
}

func CoinInfoCreator() interface{} {
	return &CoinInfo{}
}

type Account struct {
	AuthenticationKey  string `json:"authentication_key"`
	CoinRegisterEvents Events `json:"coin_register_events"`
	GuidCreationNum    uint64 `json:"guid_creation_num,string"`
	KeyRotationEvents  Events `json:"key_rotation_events"`
	SequenceNumber     uint64 `json:"sequence_number,string"`
}

func AccountCreator() interface{} {
	return &Account{}
}
