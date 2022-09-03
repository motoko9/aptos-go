package rpcmodule

import (
	"encoding/json"
)

type Events []Event

type Event struct {
	Key            string          `json:"key"`
	SequenceNumber uint64          `json:"sequence_number,string"`
	Type           string          `json:"type"`
	Raw            json.RawMessage `json:"data"`
	Object         interface{} `json:",omitempty"`
}

func (j Event) MarshalJSON() ([]byte, error) {
	if len(j.Raw) == 0 {
		raw, _ := json.Marshal(j.Object)
		j.Raw = raw
	}
	type Aux Event
	aux := Aux(j)
	return json.Marshal(aux)
}

func (j *Event) UnmarshalJSON(data []byte) error {
	type Aux Event
	aux := (*Aux)(j)
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	//
	object := createEventObject(j.Type)
	if object == nil {
		return nil
	}
	if err := json.Unmarshal(j.Raw, object); err != nil {
		return err
	}
	j.Object = object
	return nil
}
