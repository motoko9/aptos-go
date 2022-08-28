package rpcmodule

import (
	"encoding/json"
	"fmt"
	"time"
)

type Transactions []Transaction

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
	Signature               AccountSignature   `json:"signature,omitempty"`
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
	Signature               AccountSignature   `json:"signature,omitempty"`
	Events                  []Event            `json:"events"`
	Timestamp               uint64             `json:"timestamp,string"`
}

func (j Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Object)
}

func (j *Transaction) UnmarshalJSON(data []byte) error {
	type Aux Transaction
	aux := (*Aux)(j)
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	j.Raw = data
	switch j.Type {
	case "block_metadata_transaction":
		var transaction TransactionBlockMetadataTransaction
		if err := json.Unmarshal(data, &transaction); err != nil {
			return err
		}
		j.Object = transaction
		return nil
	case "genesis_transaction":
		var transaction TransactionGenesisTransaction
		if err := json.Unmarshal(data, &transaction); err != nil {
			return err
		}
		j.Object = transaction
		return nil
	case "pending_transaction":
		var transaction TransactionPendingTransaction
		if err := json.Unmarshal(data, &transaction); err != nil {
			return err
		}
		j.Object = transaction
		return nil
	case "state_checkpoint_transaction":
		var transaction TransactionStateCheckpointTransaction
		if err := json.Unmarshal(data, &transaction); err != nil {
			return err
		}
		j.Object = transaction
		return nil
	case "user_transaction":
		var transaction TransactionUserTransaction
		if err := json.Unmarshal(data, &transaction); err != nil {
			return err
		}
		j.Object = transaction
		return nil
	default:
		return fmt.Errorf("unsupport transaction type")
	}
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
	return json.Marshal(j.Object)
}

func (j *TransactionPayload) UnmarshalJSON(data []byte) error {
	type Aux TransactionPayload
	aux := (*Aux)(j)
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	j.Raw = data
	switch j.Type {
	case "entry_function_payload":
		var transactionPayload TransactionPayloadEntryFunctionPayload
		if err := json.Unmarshal(data, &transactionPayload); err != nil {
			return err
		}
		j.Object = transactionPayload
		return nil
	case "module_bundle_payload":
		var transactionPayload TransactionPayloadModuleBundlePayload
		if err := json.Unmarshal(data, &transactionPayload); err != nil {
			return err
		}
		j.Object = transactionPayload
		return nil
	case "script_payload":
		var transactionPayload TransactionPayloadScriptPayload
		if err := json.Unmarshal(data, &transactionPayload); err != nil {
			return err
		}
		j.Object = transactionPayload
		return nil
	default:
		return fmt.Errorf("unsupport transaction payload type")
	}
}

type EncodeSubmissionRequest struct {
	Sender                  string             `json:"sender"`
	SequenceNumber          uint64             `json:"sequence_number,string"`
	MaxGasAmount            uint64             `json:"max_gas_amount,string"`
	GasUnitPrice            uint64             `json:"gas_unit_price,string"`
	ExpirationTimestampSecs uint64             `json:"expiration_timestamp_secs,string"`
	Payload                 TransactionPayload `json:"payload"`
	SecondarySigners        []string           `json:"secondary_signers,omitempty"`
}

func EncodeSubmissionReq(sender string, sequence uint64, payload TransactionPayload) (*EncodeSubmissionRequest, error) {
	req := EncodeSubmissionRequest{
		Sender:                  sender,
		SequenceNumber:          sequence,
		MaxGasAmount:            uint64(2000),
		GasUnitPrice:            uint64(1),
		ExpirationTimestampSecs: uint64(time.Now().Unix() + 600), // now + 10 minutes
		Payload:                 payload,
	}
	return &req, nil
}

type SubmitTransactionRequest struct {
	Sender                  string             `json:"sender"`
	SequenceNumber          uint64             `json:"sequence_number,string"`
	MaxGasAmount            uint64             `json:"max_gas_amount,string"`
	GasUnitPrice            uint64             `json:"gas_unit_price,string"`
	ExpirationTimestampSecs uint64             `json:"expiration_timestamp_secs,string"`
	Payload                 TransactionPayload `json:"payload"`
	Signature               AccountSignature   `json:"signature"`
}

func SubmitTransactionReq(encodeSubmissionReq *EncodeSubmissionRequest, signature AccountSignature) (*SubmitTransactionRequest, error) {
	req := SubmitTransactionRequest{
		Sender:                  encodeSubmissionReq.Sender,
		SequenceNumber:          encodeSubmissionReq.SequenceNumber,
		MaxGasAmount:            encodeSubmissionReq.MaxGasAmount,
		GasUnitPrice:            encodeSubmissionReq.GasUnitPrice,
		ExpirationTimestampSecs: encodeSubmissionReq.ExpirationTimestampSecs,
		Payload:                 encodeSubmissionReq.Payload,
		Signature:               signature,
	}
	return &req, nil
}

type GasEstimate struct {
	GasEstimate uint64 `json:"gas_estimate"`
}
