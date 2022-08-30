package rpcmodule

import "strings"

type Creator func() interface{}

// AccountSignatureCreators
// for Transaction objects
//
var AccountSignatureCreators = map[string]Creator{}

func RegisterAccountSignatureCreator(t string, creator Creator) {
	AccountSignatureCreators[t] = creator
}

func createAccountSignatureObject(t string) interface{} {
	creator, ok := AccountSignatureCreators[t]
	if !ok {
		return nil
	}
	return creator()
}

// TransactionCreators
// for Transaction objects
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

// TransactionPayloadCreators
// for TransactionPayload objects
//
var TransactionPayloadCreators = map[string]Creator{}

func RegisterTransactionPayloadCreator(t string, creator Creator) {
	TransactionPayloadCreators[t] = creator
}

func createTransactionPayloadObject(t string) interface{} {
	creator, ok := TransactionPayloadCreators[t]
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
	// remove type
	index := strings.IndexByte(t, '<')
	if index != -1 {
		t = t[0:index]
	}
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

	// register transactionpayload creator
	RegisterTransactionPayloadCreator(EntryFunctionPayload, EntryFunctionPayloadCreator)
	RegisterTransactionPayloadCreator(ModuleBundlePayload, ModuleBundlePayloadCreator)
	RegisterTransactionPayloadCreator(ScriptPayload, ScriptPayloadCreator)

	// register accountsignature creator
	RegisterAccountSignatureCreator(Ed25519Signature, Ed25519SignatureCreator)
	RegisterAccountSignatureCreator(MultiEd25519Signature, MultiEd25519SignatureCreator)
}
