package rpcmodule

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Address [32]byte

func (j Address) ToString() string {
	return "0x" + hex.EncodeToString(j[:])
}

func (j Address) FromString(s string) error {
	d, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	if len(d) != 32 {
		return fmt.Errorf("address format is invalid")
	}
	copy(j[:], d)
	return nil
}

func (j Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.ToString())
}

func (j *Address) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return j.FromString(s)
}
