package rpcmodule

import "strings"

type Creator func() interface{}

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

// SignatureCreators
// for Signature objects
//
var SignatureCreators = map[string]Creator{}

func RegisterSignatureCreator(t string, creator Creator) {
	SignatureCreators[t] = creator
}

func createSignatureObject(t string) interface{} {
	creator, ok := SignatureCreators[t]
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
	// register Transaction creator
	RegisterTransactionCreator(BlockMetadataTransaction, BlockMetadataTransactionCreator)
	RegisterTransactionCreator(GenesisTransaction, GenesisTransactionCreator)
	RegisterTransactionCreator(PendingTransaction, PendingTransactionCreator)
	RegisterTransactionCreator(StateCheckpointTransaction, StateCheckpointTransactionCreator)
	RegisterTransactionCreator(UserTransaction, UserTransactionCreator)

	// register TransactionPayload creator
	RegisterTransactionPayloadCreator(EntryFunctionPayload, EntryFunctionPayloadCreator)
	RegisterTransactionPayloadCreator(ModuleBundlePayload, ModuleBundlePayloadCreator)
	RegisterTransactionPayloadCreator(ScriptPayload, ScriptPayloadCreator)

	// register Signature creator
	RegisterSignatureCreator(Ed25519Signature, Ed25519SignatureCreator)
	RegisterSignatureCreator(MultiEd25519Signature, MultiEd25519SignatureCreator)
	RegisterSignatureCreator(MultiAgentSignature, MultiAgentSignatureCreator)
}
