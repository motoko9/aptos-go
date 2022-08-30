package aptosmodule

// for events
//
type DepositEvent struct {
	Amount uint64 `json:"amount,string"`
}

func DepositEventCreator() interface{} {
	return &DepositEvent{}
}

type WithdrawEvent struct {
	Amount uint64 `json:"amount,string"`
}

func WithdrawEventCreator() interface{} {
	return &WithdrawEvent{}
}
