package rpcmodule

type TransferEvent struct {
	Amount uint64 `json:"amount,string"`
}

func TransferEventCreator() interface{} {
	return &TransferEvent{}
}
