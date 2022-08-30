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
)

func Ed25519SignatureCreator() interface{} {
	return &AccountSignatureEd25519Signature{}
}

func MultiEd25519SignatureCreator() interface{} {
	return &AccountSignatureMultiEd25519Signature{}
}

type AccountSignature struct {
	Type   string `json:"type"`
	Raw    json.RawMessage
	Object interface{}
}

type AccountSignatureEd25519Signature struct {
	Type      string `json:"type"`
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
}

type AccountSignatureMultiEd25519Signature struct {
	Type       string   `json:"type"`
	PublicKeys []string `json:"public_keys"`
	Signatures []string `json:"signatures"`
	Threshold  uint64   `json:"threshold"`
	Bitmap     string   `json:"bitmap"`
}

func (j AccountSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Object)
}

func (j *AccountSignature) UnmarshalJSON(data []byte) error {
	type Aux AccountSignature
	aux := (*Aux)(j)
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	j.Raw = data
	//
	object := createAccountSignatureObject(j.Type)
	if object == nil {
		return fmt.Errorf("unsupport account signature type")
	}
	if err := json.Unmarshal(data, object); err != nil {
		return err
	}
	j.Object = object
	return nil
}
