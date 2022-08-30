package rpcmodule

type Creator func() interface{}

// TransactionCreators
// for event objects
//
var TransactionCreators = map[string]Creator{}

func RegisterTransactionCreator(t string, creator Creator) {
	TransactionCreators[t] = creator
}

func createTransactionObject(t string) interface{} {
	creator, ok := TransactionCreators[t]
	if !ok {
		return nil
	}
	return creator()
}

// EventObjectCreators
// for event objects
//
var EventObjectCreators = map[string]Creator{}

func RegisterEventObjectCreator(t string, creator Creator) {
	EventObjectCreators[t] = creator
}

func createEventObject(t string) interface{} {
	creator, ok := EventObjectCreators[t]
	if !ok {
		return nil
	}
	return creator()
}

// ResourceObjectCreators
// for resource objects
//
var ResourceObjectCreators = map[string]Creator{}

func RegisterResourceObjectCreator(t string, creator Creator) {
	ResourceObjectCreators[t] = creator
}

func createResourceObject(t string) interface{} {
	creator, ok := ResourceObjectCreators[t]
	if !ok {
		return nil
	}
	return creator()
}

func init() {
	// register transaction creator
	RegisterTransactionCreator(BlockMetadataTransaction, BlockMetadataTransactionCreator)
	RegisterTransactionCreator(GenesisTransaction, GenesisTransactionCreator)
	RegisterTransactionCreator(PendingTransaction, PendingTransactionCreator)
	RegisterTransactionCreator(StateCheckpointTransaction, StateCheckpointTransactionCreator)
	RegisterTransactionCreator(UserTransaction, UserTransactionCreator)

	RegisterEventObjectCreator("0x1::coin::WithdrawEvent", TransferEventCreator)
	RegisterEventObjectCreator("0x1::coin::DepositEvent", TransferEventCreator)
}
