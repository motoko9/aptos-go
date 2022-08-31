package rpcmodule

import (
	"encoding/json"
	"fmt"
)

type AccountData struct {
	SequenceNumber    uint64 `json:"sequence_number,string"`
	AuthenticationKey string `json:"authentication_key"`
}

const (
	Ed25519Signature      = "ed25519_signature"
	MultiEd25519Signature = "multi_ed25519_signature"
	MultiAgentSignature   = "multi_agent_signature"
)

func Ed25519SignatureCreator() interface{} {
	return &SignatureEd25519Signature{}
}

func MultiEd25519SignatureCreator() interface{} {
	return &SignatureMultiEd25519Signature{}
}

func MultiAgentSignatureCreator() interface{} {
	return &SignatureMultiAgentSignature{}
}

type Signature struct {
	Type   string `json:"type"`
	Raw    json.RawMessage
	Object interface{}
}

type SignatureEd25519Signature struct {
	Type      string `json:"type"`
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
}

type SignatureMultiEd25519Signature struct {
	Type       string   `json:"type"`
	PublicKeys []string `json:"public_keys"`
	Signatures []string `json:"signatures"`
	Threshold  uint64   `json:"threshold"`
	Bitmap     string   `json:"bitmap"`
}

type SignatureMultiAgentSignature struct {
	Type                     string      `json:"type"`
	Signature                Signature   `json:"sender"`
	SecondarySignerAddresses []string    `json:"secondary_signer_addresses"`
	SecondarySigners         []Signature `json:"secondary_signers"`
}

func (j Signature) MarshalJSON() ([]byte, error) {
	raw, err := json.Marshal(j.Object)
	if err != nil {
		return nil, err
	}
	j.Raw = raw
	return raw, nil
}

func (j *Signature) UnmarshalJSON(data []byte) error {
	type Aux Signature
	aux := (*Aux)(j)
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	j.Raw = data
	//
	object := createSignatureObject(j.Type)
	if object == nil {
		return fmt.Errorf("unsupport signature type")
	}
	if err := json.Unmarshal(data, object); err != nil {
		return err
	}
	j.Object = object
	return nil
}
