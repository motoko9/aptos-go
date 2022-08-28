package rpcmodule

import (
	"encoding/json"
	"fmt"
)

type AccountData struct {
	SequenceNumber    uint64 `json:"sequence_number,string"`
	AuthenticationKey string `json:"authentication_key"`
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
	switch j.Type {
	case "ed25519_signature":
		var accountSignature AccountSignatureEd25519Signature
		if err := json.Unmarshal(data, &accountSignature); err != nil {
			return err
		}
		j.Object = accountSignature
		return nil
	case "multi_ed25519_signature":
		var accountSignature AccountSignatureMultiEd25519Signature
		if err := json.Unmarshal(data, &accountSignature); err != nil {
			return err
		}
		j.Object = accountSignature
		return nil
	default:
		return fmt.Errorf("unsupport account signature type")
	}
}
