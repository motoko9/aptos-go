package rpcmodule

import (
	"encoding/json"
	"fmt"
	"time"
)

type Transactions []Transaction

const (
	BlockMetadataTransaction   = "block_metadata_transaction"
	GenesisTransaction         = "genesis_transaction"
	PendingTransaction         = "pending_transaction"
	StateCheckpointTransaction = "state_checkpoint_transaction"
	UserTransaction            = "user_transaction"
)

func BlockMetadataTransactionCreator() interface{} {
	return &TransactionBlockMetadataTransaction{}
}

func GenesisTransactionCreator() interface{} {
	return &TransactionGenesisTransaction{}
}

func PendingTransactionCreator() interface{} {
	return &TransactionPendingTransaction{}
}

func StateCheckpointTransactionCreator() interface{} {
	return &TransactionStateCheckpointTransaction{}
}

func UserTransactionCreator() interface{} {
	return &TransactionUserTransaction{}
}

type Transaction struct {
	Type   string `json:"type"`
	Raw    json.RawMessage
	Object interface{}
}

type TransactionPendingTransaction struct {
	Type                    string             `json:"type"`
	Hash                    string             `json:"hash"`
	Sender                  string             `json:"sender"`
	SequenceNumber          uint64             `json:"sequence_number,string"`
	MaxGasAmount            uint64             `json:"max_gas_amount,string"`
	GasUnitPrice            uint64             `json:"gas_unit_price,string"`
	ExpirationTimestampSecs uint64             `json:"expiration_timestamp_secs,string"`
	Payload                 TransactionPayload `json:"payload"`
	Signature               Signature          `json:"signature,omitempty"`
}

type TransactionStateCheckpointTransaction struct {
	Type                string           `json:"type"`
	Version             uint64           `json:"version,string"`
	Hash                string           `json:"hash"`
	StateChangeHash     string           `json:"state_change_hash"`
	EventRootHash       string           `json:"event_root_hash"`
	StateCheckpointHash string           `json:"state_checkpoint_hash"`
	GasUsed             uint64           `json:"gas_used,string"`
	Success             bool             `json:"success"`
	VmStatus            string           `json:"vm_status"`
	AccumulatorRootHash string           `json:"accumulator_root_hash"`
	Changes             []WriteSetChange `json:"changes"`
	Timestamp           uint64           `json:"timestamp,string"`
}

type TransactionBlockMetadataTransaction struct {
	Type                   string           `json:"type"`
	Version                uint64           `json:"version,string"`
	Hash                   string           `json:"hash"`
	StateChangeHash        string           `json:"state_change_hash"`
	EventRootHash          string           `json:"event_root_hash"`
	StateCheckpointHash    string           `json:"state_checkpoint_hash"`
	GasUsed                uint64           `json:"gas_used,string"`
	Success                bool             `json:"success"`
	VmStatus               string           `json:"vm_status"`
	AccumulatorRootHash    string           `json:"accumulator_root_hash"`
	Changes                []WriteSetChange `json:"changes"`
	Id                     string           `json:"id"`
	Epoch                  uint64           `json:"epoch,string"`
	Round                  uint64           `json:"round,string"`
	Events                 []Event          `json:"events"`
	PreviousBlockVotesBits []int            `json:"previous_block_votes_bitvec"`
	Proposer               string           `json:'proposer'`
	FailedProposerIndices  []uint64         `json:"failed_proposer_indices"`
	Timestamp              uint64           `json:"timestamp,string"`
}

type TransactionGenesisTransaction struct {
	Type                string           `json:"type"`
	Version             uint64           `json:"version,string"`
	Hash                string           `json:"hash"`
	StateChangeHash     string           `json:"state_change_hash"`
	EventRootHash       string           `json:"event_root_hash"`
	StateCheckpointHash string           `json:"state_checkpoint_hash"`
	GasUsed             uint64           `json:"gas_used,string"`
	Success             bool             `json:"success"`
	VmStatus            string           `json:"vm_status"`
	AccumulatorRootHash string           `json:"accumulator_root_hash"`
	Changes             []WriteSetChange `json:"changes"`
	// todo
	Payload TransactionPayload `json:"payload"`
	Events  []Event            `json:"events"`
}

type TransactionUserTransaction struct {
	Type                    string             `json:"type"`
	Version                 uint64             `json:"version,string"`
	Hash                    string             `json:"hash"`
	StateChangeHash         string             `json:"state_change_hash"`
	EventRootHash           string             `json:"event_root_hash"`
	StateCheckpointHash     string             `json:"state_checkpoint_hash"`
	GasUsed                 uint64             `json:"gas_used,string"`
	Success                 bool               `json:"success"`
	VmStatus                string             `json:"vm_status"`
	AccumulatorRootHash     string             `json:"accumulator_root_hash"`
	Changes                 []WriteSetChange   `json:"changes"`
	Sender                  string             `json:"sender"`
	SequenceNumber          uint64             `json:"sequence_number,string"`
	MaxGasAmount            uint64             `json:"max_gas_amount,string"`
	GasUnitPrice            uint64             `json:"gas_unit_price,string"`
	ExpirationTimestampSecs uint64             `json:"expiration_timestamp_secs,string"`
	Payload                 TransactionPayload `json:"payload"`
	Signature               Signature          `json:"signature,omitempty"`
	Events                  []Event            `json:"events"`
	Timestamp               uint64             `json:"timestamp,string"`
}

func (j Transaction) MarshalJSON() ([]byte, error) {
	raw, err := json.Marshal(j.Object)
	if err != nil {
		return nil, err
	}
	j.Raw = raw
	return raw, nil
}

func (j *Transaction) UnmarshalJSON(data []byte) error {
	type Aux Transaction
	aux := (*Aux)(j)
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	j.Raw = data
	//
	object := createTransactionObject(j.Type)
	if object == nil {
		return fmt.Errorf("unsupport transaction type")
	}
	if err := json.Unmarshal(data, object); err != nil {
		return err
	}
	j.Object = object
	return nil
}

const (
	EntryFunctionPayload = "entry_function_payload"
	ModuleBundlePayload  = "module_bundle_payload"
	ScriptPayload        = "script_payload"
)

func EntryFunctionPayloadCreator() interface{} {
	return &TransactionPayloadEntryFunctionPayload{}
}

func ModuleBundlePayloadCreator() interface{} {
	return &TransactionPayloadModuleBundlePayload{}
}

func ScriptPayloadCreator() interface{} {
	return &TransactionPayloadScriptPayload{}
}

type TransactionPayload struct {
	Type   string `json:"type"`
	Raw    json.RawMessage
	Object interface{}
}

type TransactionPayloadEntryFunctionPayload struct {
	Type          string        `json:"type"`
	Function      string        `json:"function,omitempty"`
	TypeArguments []string      `json:"type_arguments"` //todo maybe need to omitempty, but move function call is needed event if empty
	Arguments     []interface{} `json:"arguments"`      //todo maybe need to omitempty, but move function call is needed event if empty
}

type TransactionPayloadModuleBundlePayload struct {
	Type    string       `json:"type"`
	Modules []MoveModule `json:"modules"`
}

type TransactionPayloadScriptPayload struct {
	Type          string        `json:"type"`
	Code          MoveCode      `json:"code"`
	TypeArguments []string      `json:"type_arguments"` //todo maybe need to omitempty, but move function call is needed event if empty
	Arguments     []interface{} `json:"arguments"`      //todo maybe need to omitempty, but move function call is needed event if empty
}

func (j TransactionPayload) MarshalJSON() ([]byte, error) {
	raw, err := json.Marshal(j.Object)
	if err != nil {
		return nil, err
	}
	j.Raw = raw
	return raw, nil
}

func (j *TransactionPayload) UnmarshalJSON(data []byte) error {
	type Aux TransactionPayload
	aux := (*Aux)(j)
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	j.Raw = data
	//
	object := createTransactionPayloadObject(j.Type)
	if object == nil {
		return fmt.Errorf("unsupport transaction payload type")
	}
	if err := json.Unmarshal(data, object); err != nil {
		return err
	}
	j.Object = object
	return nil
}

type EncodeSubmissionRequest struct {
	Sender                  string              `json:"sender"`
	SequenceNumber          uint64              `json:"sequence_number,string"`
	MaxGasAmount            uint64              `json:"max_gas_amount,string"`
	GasUnitPrice            uint64              `json:"gas_unit_price,string"`
	ExpirationTimestampSecs uint64              `json:"expiration_timestamp_secs,string"`
	Payload                 *TransactionPayload `json:"payload"`
	SecondarySigners        []string            `json:"secondary_signers,omitempty"`
}

func EncodeSubmissionReq(sender string, sequence uint64, payload *TransactionPayload) *EncodeSubmissionRequest {
	req := EncodeSubmissionRequest{
		Sender:                  sender,
		SequenceNumber:          sequence,
		MaxGasAmount:            uint64(2000),
		GasUnitPrice:            uint64(1),
		ExpirationTimestampSecs: uint64(time.Now().Unix() + 600), // now + 10 minutes
		Payload:                 payload,
	}
	return &req
}

func EncodeSubmissionWithSecondarySignersReq(sender string, sequence uint64, payload *TransactionPayload, secondarySigners []string) *EncodeSubmissionRequest {
	return &EncodeSubmissionRequest{
		Sender:                  sender,
		SequenceNumber:          sequence,
		MaxGasAmount:            uint64(2000),
		GasUnitPrice:            uint64(1),
		ExpirationTimestampSecs: uint64(time.Now().Unix() + 600), // now + 10 minutes
		Payload:                 payload,
		SecondarySigners:        secondarySigners,
	}
}

type SubmitTransactionRequest struct {
	Sender                  string              `json:"sender"`
	SequenceNumber          uint64              `json:"sequence_number,string"`
	MaxGasAmount            uint64              `json:"max_gas_amount,string"`
	GasUnitPrice            uint64              `json:"gas_unit_price,string"`
	ExpirationTimestampSecs uint64              `json:"expiration_timestamp_secs,string"`
	Payload                 *TransactionPayload `json:"payload"`
	Signature               Signature           `json:"signature"`
}

func SubmitTransactionReq(encodeSubmissionReq *EncodeSubmissionRequest, signature Signature) *SubmitTransactionRequest {
	return &SubmitTransactionRequest{
		Sender:                  encodeSubmissionReq.Sender,
		SequenceNumber:          encodeSubmissionReq.SequenceNumber,
		MaxGasAmount:            encodeSubmissionReq.MaxGasAmount,
		GasUnitPrice:            encodeSubmissionReq.GasUnitPrice,
		ExpirationTimestampSecs: encodeSubmissionReq.ExpirationTimestampSecs,
		Payload:                 encodeSubmissionReq.Payload,
		Signature:               signature,
	}
}

type GasEstimate struct {
	GasEstimate uint64 `json:"gas_estimate"`
}

type PendingTransactionRsp struct {
	Hash                    string             `json:"hash"`
	Sender                  string             `json:"sender"`
	SequenceNumber          uint64             `json:"sequence_number,string"`
	MaxGasAmount            uint64             `json:"max_gas_amount,string"`
	GasUnitPrice            uint64             `json:"gas_unit_price,string"`
	ExpirationTimestampSecs uint64             `json:"expiration_timestamp_secs,string"`
	Payload                 TransactionPayload `json:"payload"`
	Signature               Signature          `json:"signature,omitempty"`
}

type UserTransactionRsp struct {
	Version                 uint64             `json:"version,string"`
	Hash                    string             `json:"hash"`
	StateChangeHash         string             `json:"state_change_hash"`
	EventRootHash           string             `json:"event_root_hash"`
	StateCheckpointHash     string             `json:"state_checkpoint_hash"`
	GasUsed                 uint64             `json:"gas_used,string"`
	Success                 bool               `json:"success"`
	VmStatus                string             `json:"vm_status"`
	AccumulatorRootHash     string             `json:"accumulator_root_hash"`
	Changes                 []WriteSetChange   `json:"changes"`
	Sender                  string             `json:"sender"`
	SequenceNumber          uint64             `json:"sequence_number,string"`
	MaxGasAmount            uint64             `json:"max_gas_amount,string"`
	GasUnitPrice            uint64             `json:"gas_unit_price,string"`
	ExpirationTimestampSecs uint64             `json:"expiration_timestamp_secs,string"`
	Payload                 TransactionPayload `json:"payload"`
	Signature               Signature          `json:"signature,omitempty"`
	Events                  []Event            `json:"events"`
	Timestamp               uint64             `json:"timestamp,string"`
}
