package utils

import (
	"fmt"
	"github.com/aptos-labs/serde-reflection/serde-generate/runtime/golang/serde"
)

type AccessPath struct {
	Address AccountAddress
	Path    []byte
}

func (obj *AccessPath) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Address.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(obj.Path); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeAccessPath(deserializer serde.Deserializer) (AccessPath, error) {
	var obj AccessPath
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Address = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.Path = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type AccountAddress [32]uint8

func (obj *AccountAddress) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serialize_array32_u8_array((([32]uint8)(*obj)), serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeAccountAddress(deserializer serde.Deserializer) (AccountAddress, error) {
	var obj [32]uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (AccountAddress)(obj), err
	}
	if val, err := deserialize_array32_u8_array(deserializer); err == nil {
		obj = val
	} else {
		return ((AccountAddress)(obj)), err
	}
	deserializer.DecreaseContainerDepth()
	return (AccountAddress)(obj), nil
}

type AccountAuthenticator interface {
	isAccountAuthenticator()
	Serialize(serializer serde.Serializer) error
}

func DeserializeAccountAuthenticator(deserializer serde.Deserializer) (AccountAuthenticator, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_AccountAuthenticator__Ed25519(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_AccountAuthenticator__MultiEd25519(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for AccountAuthenticator: %d", index)
	}
}

type AccountAuthenticator__Ed25519 struct {
	PublicKey Ed25519PublicKey
	Signature Ed25519Signature
}

func (*AccountAuthenticator__Ed25519) isAccountAuthenticator() {}

func (obj *AccountAuthenticator__Ed25519) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := obj.PublicKey.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Signature.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_AccountAuthenticator__Ed25519(deserializer serde.Deserializer) (AccountAuthenticator__Ed25519, error) {
	var obj AccountAuthenticator__Ed25519
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeEd25519PublicKey(deserializer); err == nil {
		obj.PublicKey = val
	} else {
		return obj, err
	}
	if val, err := DeserializeEd25519Signature(deserializer); err == nil {
		obj.Signature = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type AccountAuthenticator__MultiEd25519 struct {
	PublicKey MultiEd25519PublicKey
	Signature MultiEd25519Signature
}

func (*AccountAuthenticator__MultiEd25519) isAccountAuthenticator() {}

func (obj *AccountAuthenticator__MultiEd25519) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	if err := obj.PublicKey.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Signature.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_AccountAuthenticator__MultiEd25519(deserializer serde.Deserializer) (AccountAuthenticator__MultiEd25519, error) {
	var obj AccountAuthenticator__MultiEd25519
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeMultiEd25519PublicKey(deserializer); err == nil {
		obj.PublicKey = val
	} else {
		return obj, err
	}
	if val, err := DeserializeMultiEd25519Signature(deserializer); err == nil {
		obj.Signature = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type BlockMetadata struct {
	Id                       HashValue
	Epoch                    uint64
	Round                    uint64
	Proposer                 AccountAddress
	PreviousBlockVotesBitvec []byte
	FailedProposerIndices    []uint32
	TimestampUsecs           uint64
}

func (obj *BlockMetadata) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Id.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.Epoch); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.Round); err != nil {
		return err
	}
	if err := obj.Proposer.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(obj.PreviousBlockVotesBitvec); err != nil {
		return err
	}
	if err := serialize_vector_u32(obj.FailedProposerIndices, serializer); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.TimestampUsecs); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeBlockMetadata(deserializer serde.Deserializer) (BlockMetadata, error) {
	var obj BlockMetadata
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeHashValue(deserializer); err == nil {
		obj.Id = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.Epoch = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.Round = val
	} else {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Proposer = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.PreviousBlockVotesBitvec = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_u32(deserializer); err == nil {
		obj.FailedProposerIndices = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.TimestampUsecs = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ChainId uint8

func (obj *ChainId) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeU8(((uint8)(*obj))); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeChainId(deserializer serde.Deserializer) (ChainId, error) {
	var obj uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (ChainId)(obj), err
	}
	if val, err := deserializer.DeserializeU8(); err == nil {
		obj = val
	} else {
		return ((ChainId)(obj)), err
	}
	deserializer.DecreaseContainerDepth()
	return (ChainId)(obj), nil
}

type ChangeSet struct {
	WriteSet WriteSet
	Events   []ContractEvent
}

func (obj *ChangeSet) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.WriteSet.Serialize(serializer); err != nil {
		return err
	}
	if err := serialize_vector_ContractEvent(obj.Events, serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeChangeSet(deserializer serde.Deserializer) (ChangeSet, error) {
	var obj ChangeSet
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeWriteSet(deserializer); err == nil {
		obj.WriteSet = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_ContractEvent(deserializer); err == nil {
		obj.Events = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ContractEvent interface {
	isContractEvent()
	Serialize(serializer serde.Serializer) error
}

func DeserializeContractEvent(deserializer serde.Deserializer) (ContractEvent, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_ContractEvent__V0(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for ContractEvent: %d", index)
	}
}

type ContractEvent__V0 struct {
	Value ContractEventV0
}

func (*ContractEvent__V0) isContractEvent() {}

func (obj *ContractEvent__V0) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_ContractEvent__V0(deserializer serde.Deserializer) (ContractEvent__V0, error) {
	var obj ContractEvent__V0
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeContractEventV0(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ContractEventV0 struct {
	Key            EventKey
	SequenceNumber uint64
	TypeTag        TypeTag
	EventData      []byte
}

func (obj *ContractEventV0) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Key.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.SequenceNumber); err != nil {
		return err
	}
	if err := obj.TypeTag.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(obj.EventData); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeContractEventV0(deserializer serde.Deserializer) (ContractEventV0, error) {
	var obj ContractEventV0
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeEventKey(deserializer); err == nil {
		obj.Key = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.SequenceNumber = val
	} else {
		return obj, err
	}
	if val, err := DeserializeTypeTag(deserializer); err == nil {
		obj.TypeTag = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.EventData = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Ed25519PublicKey []byte

func (obj *Ed25519PublicKey) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeEd25519PublicKey(deserializer serde.Deserializer) (Ed25519PublicKey, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (Ed25519PublicKey)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return ((Ed25519PublicKey)(obj)), err
	}
	deserializer.DecreaseContainerDepth()
	return (Ed25519PublicKey)(obj), nil
}

type Ed25519Signature []byte

func (obj *Ed25519Signature) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeEd25519Signature(deserializer serde.Deserializer) (Ed25519Signature, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (Ed25519Signature)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return ((Ed25519Signature)(obj)), err
	}
	deserializer.DecreaseContainerDepth()
	return (Ed25519Signature)(obj), nil
}

type EntryFunction struct {
	Module   ModuleId
	Function Identifier
	TyArgs   []TypeTag
	Args     [][]byte
}

func (obj *EntryFunction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Module.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Function.Serialize(serializer); err != nil {
		return err
	}
	if err := serialize_vector_TypeTag(obj.TyArgs, serializer); err != nil {
		return err
	}
	if err := serialize_vector_bytes(obj.Args, serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeEntryFunction(deserializer serde.Deserializer) (EntryFunction, error) {
	var obj EntryFunction
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeModuleId(deserializer); err == nil {
		obj.Module = val
	} else {
		return obj, err
	}
	if val, err := DeserializeIdentifier(deserializer); err == nil {
		obj.Function = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_TypeTag(deserializer); err == nil {
		obj.TyArgs = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_bytes(deserializer); err == nil {
		obj.Args = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type EventKey struct {
	CreationNumber uint64
	AccountAddress AccountAddress
}

func (obj *EventKey) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.CreationNumber); err != nil {
		return err
	}
	if err := obj.AccountAddress.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeEventKey(deserializer serde.Deserializer) (EventKey, error) {
	var obj EventKey
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.CreationNumber = val
	} else {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.AccountAddress = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type HashValue struct {
	Hash [32]uint8
}

func (obj *HashValue) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serialize_array32_u8_array(obj.Hash, serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeHashValue(deserializer serde.Deserializer) (HashValue, error) {
	var obj HashValue
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserialize_array32_u8_array(deserializer); err == nil {
		obj.Hash = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Identifier string

func (obj *Identifier) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeStr(((string)(*obj))); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeIdentifier(deserializer serde.Deserializer) (Identifier, error) {
	var obj string
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (Identifier)(obj), err
	}
	if val, err := deserializer.DeserializeStr(); err == nil {
		obj = val
	} else {
		return ((Identifier)(obj)), err
	}
	deserializer.DecreaseContainerDepth()
	return (Identifier)(obj), nil
}

type Module struct {
	Code []byte
}

func (obj *Module) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(obj.Code); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeModule(deserializer serde.Deserializer) (Module, error) {
	var obj Module
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.Code = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ModuleBundle struct {
	Codes []Module
}

func (obj *ModuleBundle) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serialize_vector_Module(obj.Codes, serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeModuleBundle(deserializer serde.Deserializer) (ModuleBundle, error) {
	var obj ModuleBundle
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserialize_vector_Module(deserializer); err == nil {
		obj.Codes = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ModuleId struct {
	Address AccountAddress
	Name    Identifier
}

func (obj *ModuleId) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Address.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Name.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeModuleId(deserializer serde.Deserializer) (ModuleId, error) {
	var obj ModuleId
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Address = val
	} else {
		return obj, err
	}
	if val, err := DeserializeIdentifier(deserializer); err == nil {
		obj.Name = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type MultiEd25519PublicKey []byte

func (obj *MultiEd25519PublicKey) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeMultiEd25519PublicKey(deserializer serde.Deserializer) (MultiEd25519PublicKey, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (MultiEd25519PublicKey)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return ((MultiEd25519PublicKey)(obj)), err
	}
	deserializer.DecreaseContainerDepth()
	return (MultiEd25519PublicKey)(obj), nil
}

type MultiEd25519Signature []byte

func (obj *MultiEd25519Signature) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeMultiEd25519Signature(deserializer serde.Deserializer) (MultiEd25519Signature, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (MultiEd25519Signature)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return ((MultiEd25519Signature)(obj)), err
	}
	deserializer.DecreaseContainerDepth()
	return (MultiEd25519Signature)(obj), nil
}

type RawTransaction struct {
	Sender                  AccountAddress
	SequenceNumber          uint64
	Payload                 TransactionPayload
	MaxGasAmount            uint64
	GasUnitPrice            uint64
	ExpirationTimestampSecs uint64
	ChainId                 ChainId
}

func (obj *RawTransaction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Sender.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.SequenceNumber); err != nil {
		return err
	}
	if err := obj.Payload.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.MaxGasAmount); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.GasUnitPrice); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.ExpirationTimestampSecs); err != nil {
		return err
	}
	if err := obj.ChainId.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeRawTransaction(deserializer serde.Deserializer) (RawTransaction, error) {
	var obj RawTransaction
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Sender = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.SequenceNumber = val
	} else {
		return obj, err
	}
	if val, err := DeserializeTransactionPayload(deserializer); err == nil {
		obj.Payload = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.MaxGasAmount = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.GasUnitPrice = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.ExpirationTimestampSecs = val
	} else {
		return obj, err
	}
	if val, err := DeserializeChainId(deserializer); err == nil {
		obj.ChainId = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Script struct {
	Code   []byte
	TyArgs []TypeTag
	Args   []TransactionArgument
}

func (obj *Script) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(obj.Code); err != nil {
		return err
	}
	if err := serialize_vector_TypeTag(obj.TyArgs, serializer); err != nil {
		return err
	}
	if err := serialize_vector_TransactionArgument(obj.Args, serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeScript(deserializer serde.Deserializer) (Script, error) {
	var obj Script
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.Code = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_TypeTag(deserializer); err == nil {
		obj.TyArgs = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_TransactionArgument(deserializer); err == nil {
		obj.Args = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type SignedTransaction struct {
	RawTxn        RawTransaction
	Authenticator TransactionAuthenticator
}

func (obj *SignedTransaction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.RawTxn.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Authenticator.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeSignedTransaction(deserializer serde.Deserializer) (SignedTransaction, error) {
	var obj SignedTransaction
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeRawTransaction(deserializer); err == nil {
		obj.RawTxn = val
	} else {
		return obj, err
	}
	if val, err := DeserializeTransactionAuthenticator(deserializer); err == nil {
		obj.Authenticator = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type StateKey interface {
	isStateKey()
	Serialize(serializer serde.Serializer) error
}

func DeserializeStateKey(deserializer serde.Deserializer) (StateKey, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_StateKey__AccessPath(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_StateKey__TableItem(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_StateKey__Raw(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for StateKey: %d", index)
	}
}

type StateKey__AccessPath struct {
	Value AccessPath
}

func (*StateKey__AccessPath) isStateKey() {}

func (obj *StateKey__AccessPath) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_StateKey__AccessPath(deserializer serde.Deserializer) (StateKey__AccessPath, error) {
	var obj StateKey__AccessPath
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccessPath(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type StateKey__TableItem struct {
	Handle TableHandle
	Key    []byte
}

func (*StateKey__TableItem) isStateKey() {}

func (obj *StateKey__TableItem) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	if err := obj.Handle.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(obj.Key); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_StateKey__TableItem(deserializer serde.Deserializer) (StateKey__TableItem, error) {
	var obj StateKey__TableItem
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeTableHandle(deserializer); err == nil {
		obj.Handle = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.Key = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type StateKey__Raw []byte

func (*StateKey__Raw) isStateKey() {}

func (obj *StateKey__Raw) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(2)
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_StateKey__Raw(deserializer serde.Deserializer) (StateKey__Raw, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (StateKey__Raw)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return ((StateKey__Raw)(obj)), err
	}
	deserializer.DecreaseContainerDepth()
	return (StateKey__Raw)(obj), nil
}

type StructTag struct {
	Address  AccountAddress
	Module   Identifier
	Name     Identifier
	TypeArgs []TypeTag
}

func (obj *StructTag) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Address.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Module.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Name.Serialize(serializer); err != nil {
		return err
	}
	if err := serialize_vector_TypeTag(obj.TypeArgs, serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeStructTag(deserializer serde.Deserializer) (StructTag, error) {
	var obj StructTag
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Address = val
	} else {
		return obj, err
	}
	if val, err := DeserializeIdentifier(deserializer); err == nil {
		obj.Module = val
	} else {
		return obj, err
	}
	if val, err := DeserializeIdentifier(deserializer); err == nil {
		obj.Name = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_TypeTag(deserializer); err == nil {
		obj.TypeArgs = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TableHandle struct {
	Value AccountAddress
}

func (obj *TableHandle) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeTableHandle(deserializer serde.Deserializer) (TableHandle, error) {
	var obj TableHandle
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Transaction interface {
	isTransaction()
	Serialize(serializer serde.Serializer) error
}

func DeserializeTransaction(deserializer serde.Deserializer) (Transaction, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_Transaction__UserTransaction(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_Transaction__GenesisTransaction(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_Transaction__BlockMetadata(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_Transaction__StateCheckpoint(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for Transaction: %d", index)
	}
}

type Transaction__UserTransaction struct {
	Value SignedTransaction
}

func (*Transaction__UserTransaction) isTransaction() {}

func (obj *Transaction__UserTransaction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_Transaction__UserTransaction(deserializer serde.Deserializer) (Transaction__UserTransaction, error) {
	var obj Transaction__UserTransaction
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeSignedTransaction(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Transaction__GenesisTransaction struct {
	Value WriteSetPayload
}

func (*Transaction__GenesisTransaction) isTransaction() {}

func (obj *Transaction__GenesisTransaction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_Transaction__GenesisTransaction(deserializer serde.Deserializer) (Transaction__GenesisTransaction, error) {
	var obj Transaction__GenesisTransaction
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeWriteSetPayload(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Transaction__BlockMetadata struct {
	Value BlockMetadata
}

func (*Transaction__BlockMetadata) isTransaction() {}

func (obj *Transaction__BlockMetadata) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(2)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_Transaction__BlockMetadata(deserializer serde.Deserializer) (Transaction__BlockMetadata, error) {
	var obj Transaction__BlockMetadata
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeBlockMetadata(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Transaction__StateCheckpoint struct {
	Value HashValue
}

func (*Transaction__StateCheckpoint) isTransaction() {}

func (obj *Transaction__StateCheckpoint) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(3)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_Transaction__StateCheckpoint(deserializer serde.Deserializer) (Transaction__StateCheckpoint, error) {
	var obj Transaction__StateCheckpoint
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeHashValue(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionArgument interface {
	isTransactionArgument()
	Serialize(serializer serde.Serializer) error
}

func DeserializeTransactionArgument(deserializer serde.Deserializer) (TransactionArgument, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_TransactionArgument__U8(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TransactionArgument__U64(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_TransactionArgument__U128(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_TransactionArgument__Address(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 4:
		if val, err := load_TransactionArgument__U8Vector(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 5:
		if val, err := load_TransactionArgument__Bool(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TransactionArgument: %d", index)
	}
}

type TransactionArgument__U8 uint8

func (*TransactionArgument__U8) isTransactionArgument() {}

func (obj *TransactionArgument__U8) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := serializer.SerializeU8(((uint8)(*obj))); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TransactionArgument__U8(deserializer serde.Deserializer) (TransactionArgument__U8, error) {
	var obj uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (TransactionArgument__U8)(obj), err
	}
	if val, err := deserializer.DeserializeU8(); err == nil {
		obj = val
	} else {
		return ((TransactionArgument__U8)(obj)), err
	}
	deserializer.DecreaseContainerDepth()
	return (TransactionArgument__U8)(obj), nil
}

type TransactionArgument__U64 uint64

func (*TransactionArgument__U64) isTransactionArgument() {}

func (obj *TransactionArgument__U64) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	if err := serializer.SerializeU64(((uint64)(*obj))); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TransactionArgument__U64(deserializer serde.Deserializer) (TransactionArgument__U64, error) {
	var obj uint64
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (TransactionArgument__U64)(obj), err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj = val
	} else {
		return ((TransactionArgument__U64)(obj)), err
	}
	deserializer.DecreaseContainerDepth()
	return (TransactionArgument__U64)(obj), nil
}

type TransactionArgument__U128 serde.Uint128

func (*TransactionArgument__U128) isTransactionArgument() {}

func (obj *TransactionArgument__U128) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(2)
	if err := serializer.SerializeU128(((serde.Uint128)(*obj))); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TransactionArgument__U128(deserializer serde.Deserializer) (TransactionArgument__U128, error) {
	var obj serde.Uint128
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (TransactionArgument__U128)(obj), err
	}
	if val, err := deserializer.DeserializeU128(); err == nil {
		obj = val
	} else {
		return ((TransactionArgument__U128)(obj)), err
	}
	deserializer.DecreaseContainerDepth()
	return (TransactionArgument__U128)(obj), nil
}

type TransactionArgument__Address struct {
	Value AccountAddress
}

func (*TransactionArgument__Address) isTransactionArgument() {}

func (obj *TransactionArgument__Address) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(3)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TransactionArgument__Address(deserializer serde.Deserializer) (TransactionArgument__Address, error) {
	var obj TransactionArgument__Address
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionArgument__U8Vector []byte

func (*TransactionArgument__U8Vector) isTransactionArgument() {}

func (obj *TransactionArgument__U8Vector) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(4)
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TransactionArgument__U8Vector(deserializer serde.Deserializer) (TransactionArgument__U8Vector, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (TransactionArgument__U8Vector)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return ((TransactionArgument__U8Vector)(obj)), err
	}
	deserializer.DecreaseContainerDepth()
	return (TransactionArgument__U8Vector)(obj), nil
}

type TransactionArgument__Bool bool

func (*TransactionArgument__Bool) isTransactionArgument() {}

func (obj *TransactionArgument__Bool) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(5)
	if err := serializer.SerializeBool(((bool)(*obj))); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TransactionArgument__Bool(deserializer serde.Deserializer) (TransactionArgument__Bool, error) {
	var obj bool
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (TransactionArgument__Bool)(obj), err
	}
	if val, err := deserializer.DeserializeBool(); err == nil {
		obj = val
	} else {
		return ((TransactionArgument__Bool)(obj)), err
	}
	deserializer.DecreaseContainerDepth()
	return (TransactionArgument__Bool)(obj), nil
}

type TransactionAuthenticator interface {
	isTransactionAuthenticator()
	Serialize(serializer serde.Serializer) error
}

func DeserializeTransactionAuthenticator(deserializer serde.Deserializer) (TransactionAuthenticator, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_TransactionAuthenticator__Ed25519(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TransactionAuthenticator__MultiEd25519(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_TransactionAuthenticator__MultiAgent(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TransactionAuthenticator: %d", index)
	}
}

type TransactionAuthenticator__Ed25519 struct {
	PublicKey Ed25519PublicKey
	Signature Ed25519Signature
}

func (*TransactionAuthenticator__Ed25519) isTransactionAuthenticator() {}

func (obj *TransactionAuthenticator__Ed25519) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := obj.PublicKey.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Signature.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TransactionAuthenticator__Ed25519(deserializer serde.Deserializer) (TransactionAuthenticator__Ed25519, error) {
	var obj TransactionAuthenticator__Ed25519
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeEd25519PublicKey(deserializer); err == nil {
		obj.PublicKey = val
	} else {
		return obj, err
	}
	if val, err := DeserializeEd25519Signature(deserializer); err == nil {
		obj.Signature = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionAuthenticator__MultiEd25519 struct {
	PublicKey MultiEd25519PublicKey
	Signature MultiEd25519Signature
}

func (*TransactionAuthenticator__MultiEd25519) isTransactionAuthenticator() {}

func (obj *TransactionAuthenticator__MultiEd25519) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	if err := obj.PublicKey.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Signature.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TransactionAuthenticator__MultiEd25519(deserializer serde.Deserializer) (TransactionAuthenticator__MultiEd25519, error) {
	var obj TransactionAuthenticator__MultiEd25519
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeMultiEd25519PublicKey(deserializer); err == nil {
		obj.PublicKey = val
	} else {
		return obj, err
	}
	if val, err := DeserializeMultiEd25519Signature(deserializer); err == nil {
		obj.Signature = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionAuthenticator__MultiAgent struct {
	Sender                   AccountAuthenticator
	SecondarySignerAddresses []AccountAddress
	SecondarySigners         []AccountAuthenticator
}

func (*TransactionAuthenticator__MultiAgent) isTransactionAuthenticator() {}

func (obj *TransactionAuthenticator__MultiAgent) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(2)
	if err := obj.Sender.Serialize(serializer); err != nil {
		return err
	}
	if err := serialize_vector_AccountAddress(obj.SecondarySignerAddresses, serializer); err != nil {
		return err
	}
	if err := serialize_vector_AccountAuthenticator(obj.SecondarySigners, serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TransactionAuthenticator__MultiAgent(deserializer serde.Deserializer) (TransactionAuthenticator__MultiAgent, error) {
	var obj TransactionAuthenticator__MultiAgent
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAuthenticator(deserializer); err == nil {
		obj.Sender = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_AccountAddress(deserializer); err == nil {
		obj.SecondarySignerAddresses = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_AccountAuthenticator(deserializer); err == nil {
		obj.SecondarySigners = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionPayload interface {
	isTransactionPayload()
	Serialize(serializer serde.Serializer) error
}

func DeserializeTransactionPayload(deserializer serde.Deserializer) (TransactionPayload, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_TransactionPayload__Script(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TransactionPayload__ModuleBundle(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_TransactionPayload__EntryFunction(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TransactionPayload: %d", index)
	}
}

type TransactionPayload__Script struct {
	Value Script
}

func (*TransactionPayload__Script) isTransactionPayload() {}

func (obj *TransactionPayload__Script) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TransactionPayload__Script(deserializer serde.Deserializer) (TransactionPayload__Script, error) {
	var obj TransactionPayload__Script
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeScript(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionPayload__ModuleBundle struct {
	Value ModuleBundle
}

func (*TransactionPayload__ModuleBundle) isTransactionPayload() {}

func (obj *TransactionPayload__ModuleBundle) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TransactionPayload__ModuleBundle(deserializer serde.Deserializer) (TransactionPayload__ModuleBundle, error) {
	var obj TransactionPayload__ModuleBundle
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeModuleBundle(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionPayload__EntryFunction struct {
	Value EntryFunction
}

func (*TransactionPayload__EntryFunction) isTransactionPayload() {}

func (obj *TransactionPayload__EntryFunction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(2)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TransactionPayload__EntryFunction(deserializer serde.Deserializer) (TransactionPayload__EntryFunction, error) {
	var obj TransactionPayload__EntryFunction
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeEntryFunction(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag interface {
	isTypeTag()
	Serialize(serializer serde.Serializer) error
}

func DeserializeTypeTag(deserializer serde.Deserializer) (TypeTag, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_TypeTag__Bool(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TypeTag__U8(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_TypeTag__U64(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_TypeTag__U128(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 4:
		if val, err := load_TypeTag__Address(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 5:
		if val, err := load_TypeTag__Signer(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 6:
		if val, err := load_TypeTag__Vector(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 7:
		if val, err := load_TypeTag__Struct(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TypeTag: %d", index)
	}
}

type TypeTag__Bool struct {
}

func (*TypeTag__Bool) isTypeTag() {}

func (obj *TypeTag__Bool) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TypeTag__Bool(deserializer serde.Deserializer) (TypeTag__Bool, error) {
	var obj TypeTag__Bool
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__U8 struct {
}

func (*TypeTag__U8) isTypeTag() {}

func (obj *TypeTag__U8) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TypeTag__U8(deserializer serde.Deserializer) (TypeTag__U8, error) {
	var obj TypeTag__U8
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__U64 struct {
}

func (*TypeTag__U64) isTypeTag() {}

func (obj *TypeTag__U64) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(2)
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TypeTag__U64(deserializer serde.Deserializer) (TypeTag__U64, error) {
	var obj TypeTag__U64
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__U128 struct {
}

func (*TypeTag__U128) isTypeTag() {}

func (obj *TypeTag__U128) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(3)
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TypeTag__U128(deserializer serde.Deserializer) (TypeTag__U128, error) {
	var obj TypeTag__U128
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__Address struct {
}

func (*TypeTag__Address) isTypeTag() {}

func (obj *TypeTag__Address) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(4)
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TypeTag__Address(deserializer serde.Deserializer) (TypeTag__Address, error) {
	var obj TypeTag__Address
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__Signer struct {
}

func (*TypeTag__Signer) isTypeTag() {}

func (obj *TypeTag__Signer) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(5)
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TypeTag__Signer(deserializer serde.Deserializer) (TypeTag__Signer, error) {
	var obj TypeTag__Signer
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__Vector struct {
	Value TypeTag
}

func (*TypeTag__Vector) isTypeTag() {}

func (obj *TypeTag__Vector) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(6)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TypeTag__Vector(deserializer serde.Deserializer) (TypeTag__Vector, error) {
	var obj TypeTag__Vector
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeTypeTag(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__Struct struct {
	Value StructTag
}

func (*TypeTag__Struct) isTypeTag() {}

func (obj *TypeTag__Struct) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(7)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_TypeTag__Struct(deserializer serde.Deserializer) (TypeTag__Struct, error) {
	var obj TypeTag__Struct
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeStructTag(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type WriteOp interface {
	isWriteOp()
	Serialize(serializer serde.Serializer) error
}

func DeserializeWriteOp(deserializer serde.Deserializer) (WriteOp, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_WriteOp__Creation(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_WriteOp__Modification(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_WriteOp__Deletion(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for WriteOp: %d", index)
	}
}

type WriteOp__Creation []byte

func (*WriteOp__Creation) isWriteOp() {}

func (obj *WriteOp__Creation) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_WriteOp__Creation(deserializer serde.Deserializer) (WriteOp__Creation, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (WriteOp__Creation)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return ((WriteOp__Creation)(obj)), err
	}
	deserializer.DecreaseContainerDepth()
	return (WriteOp__Creation)(obj), nil
}

type WriteOp__Modification []byte

func (*WriteOp__Modification) isWriteOp() {}

func (obj *WriteOp__Modification) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_WriteOp__Modification(deserializer serde.Deserializer) (WriteOp__Modification, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (WriteOp__Modification)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return ((WriteOp__Modification)(obj)), err
	}
	deserializer.DecreaseContainerDepth()
	return (WriteOp__Modification)(obj), nil
}

type WriteOp__Deletion struct {
}

func (*WriteOp__Deletion) isWriteOp() {}

func (obj *WriteOp__Deletion) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(2)
	serializer.DecreaseContainerDepth()
	return nil
}

func load_WriteOp__Deletion(deserializer serde.Deserializer) (WriteOp__Deletion, error) {
	var obj WriteOp__Deletion
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type WriteSet interface {
	isWriteSet()
	Serialize(serializer serde.Serializer) error
}

func DeserializeWriteSet(deserializer serde.Deserializer) (WriteSet, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_WriteSet__V0(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for WriteSet: %d", index)
	}
}

type WriteSet__V0 struct {
	Value WriteSetV0
}

func (*WriteSet__V0) isWriteSet() {}

func (obj *WriteSet__V0) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_WriteSet__V0(deserializer serde.Deserializer) (WriteSet__V0, error) {
	var obj WriteSet__V0
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeWriteSetV0(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type WriteSetMut struct {
	WriteSet map[StateKey]WriteOp
}

func (obj *WriteSetMut) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serialize_map_StateKey_to_WriteOp(obj.WriteSet, serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeWriteSetMut(deserializer serde.Deserializer) (WriteSetMut, error) {
	var obj WriteSetMut
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserialize_map_StateKey_to_WriteOp(deserializer); err == nil {
		obj.WriteSet = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type WriteSetPayload interface {
	isWriteSetPayload()
	Serialize(serializer serde.Serializer) error
}

func DeserializeWriteSetPayload(deserializer serde.Deserializer) (WriteSetPayload, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_WriteSetPayload__Direct(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_WriteSetPayload__Script(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for WriteSetPayload: %d", index)
	}
}

type WriteSetPayload__Direct struct {
	Value ChangeSet
}

func (*WriteSetPayload__Direct) isWriteSetPayload() {}

func (obj *WriteSetPayload__Direct) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_WriteSetPayload__Direct(deserializer serde.Deserializer) (WriteSetPayload__Direct, error) {
	var obj WriteSetPayload__Direct
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeChangeSet(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type WriteSetPayload__Script struct {
	ExecuteAs AccountAddress
	Script    Script
}

func (*WriteSetPayload__Script) isWriteSetPayload() {}

func (obj *WriteSetPayload__Script) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	if err := obj.ExecuteAs.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Script.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func load_WriteSetPayload__Script(deserializer serde.Deserializer) (WriteSetPayload__Script, error) {
	var obj WriteSetPayload__Script
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.ExecuteAs = val
	} else {
		return obj, err
	}
	if val, err := DeserializeScript(deserializer); err == nil {
		obj.Script = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type WriteSetV0 struct {
	Value WriteSetMut
}

func (obj *WriteSetV0) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func DeserializeWriteSetV0(deserializer serde.Deserializer) (WriteSetV0, error) {
	var obj WriteSetV0
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeWriteSetMut(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}
func serialize_array32_u8_array(value [32]uint8, serializer serde.Serializer) error {
	for _, item := range value {
		if err := serializer.SerializeU8(item); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_array32_u8_array(deserializer serde.Deserializer) ([32]uint8, error) {
	var obj [32]uint8
	for i := range obj {
		if val, err := deserializer.DeserializeU8(); err == nil {
			obj[i] = val
		} else {
			return obj, err
		}
	}
	return obj, nil
}

func serialize_map_StateKey_to_WriteOp(value map[StateKey]WriteOp, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	offsets := make([]uint64, len(value))
	count := 0
	for k, v := range value {
		offsets[count] = serializer.GetBufferOffset()
		count += 1
		if err := k.Serialize(serializer); err != nil {
			return err
		}
		if err := v.Serialize(serializer); err != nil {
			return err
		}
	}
	serializer.SortMapEntries(offsets)
	return nil
}

func deserialize_map_StateKey_to_WriteOp(deserializer serde.Deserializer) (map[StateKey]WriteOp, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make(map[StateKey]WriteOp)
	previous_slice := serde.Slice{0, 0}
	for i := 0; i < int(length); i++ {
		var slice serde.Slice
		slice.Start = deserializer.GetBufferOffset()
		var key StateKey
		if val, err := DeserializeStateKey(deserializer); err == nil {
			key = val
		} else {
			return nil, err
		}
		slice.End = deserializer.GetBufferOffset()
		if i > 0 {
			err := deserializer.CheckThatKeySlicesAreIncreasing(previous_slice, slice)
			if err != nil {
				return nil, err
			}
		}
		previous_slice = slice
		if val, err := DeserializeWriteOp(deserializer); err == nil {
			obj[key] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func serialize_vector_AccountAddress(value []AccountAddress, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	for _, item := range value {
		if err := item.Serialize(serializer); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_vector_AccountAddress(deserializer serde.Deserializer) ([]AccountAddress, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make([]AccountAddress, length)
	for i := range obj {
		if val, err := DeserializeAccountAddress(deserializer); err == nil {
			obj[i] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func serialize_vector_AccountAuthenticator(value []AccountAuthenticator, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	for _, item := range value {
		if err := item.Serialize(serializer); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_vector_AccountAuthenticator(deserializer serde.Deserializer) ([]AccountAuthenticator, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make([]AccountAuthenticator, length)
	for i := range obj {
		if val, err := DeserializeAccountAuthenticator(deserializer); err == nil {
			obj[i] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func serialize_vector_ContractEvent(value []ContractEvent, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	for _, item := range value {
		if err := item.Serialize(serializer); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_vector_ContractEvent(deserializer serde.Deserializer) ([]ContractEvent, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make([]ContractEvent, length)
	for i := range obj {
		if val, err := DeserializeContractEvent(deserializer); err == nil {
			obj[i] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func serialize_vector_Module(value []Module, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	for _, item := range value {
		if err := item.Serialize(serializer); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_vector_Module(deserializer serde.Deserializer) ([]Module, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make([]Module, length)
	for i := range obj {
		if val, err := DeserializeModule(deserializer); err == nil {
			obj[i] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func serialize_vector_TransactionArgument(value []TransactionArgument, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	for _, item := range value {
		if err := item.Serialize(serializer); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_vector_TransactionArgument(deserializer serde.Deserializer) ([]TransactionArgument, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make([]TransactionArgument, length)
	for i := range obj {
		if val, err := DeserializeTransactionArgument(deserializer); err == nil {
			obj[i] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func serialize_vector_TypeTag(value []TypeTag, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	for _, item := range value {
		if err := item.Serialize(serializer); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_vector_TypeTag(deserializer serde.Deserializer) ([]TypeTag, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make([]TypeTag, length)
	for i := range obj {
		if val, err := DeserializeTypeTag(deserializer); err == nil {
			obj[i] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func serialize_vector_bytes(value [][]byte, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	for _, item := range value {
		if err := serializer.SerializeBytes(item); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_vector_bytes(deserializer serde.Deserializer) ([][]byte, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make([][]byte, length)
	for i := range obj {
		if val, err := deserializer.DeserializeBytes(); err == nil {
			obj[i] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func serialize_vector_u32(value []uint32, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	for _, item := range value {
		if err := serializer.SerializeU32(item); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_vector_u32(deserializer serde.Deserializer) ([]uint32, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make([]uint32, length)
	for i := range obj {
		if val, err := deserializer.DeserializeU32(); err == nil {
			obj[i] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}
