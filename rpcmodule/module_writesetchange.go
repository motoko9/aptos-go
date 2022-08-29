package rpcmodule

import (
	"encoding/json"
	"fmt"
)

const (
	DeleteModule    = "delete_module"
	DeleteResource  = "delete_resource"
	DeleteTableItem = "delete_table_item"
	WriteModule     = "write_module"
	WriteResource   = "write_resource"
	WriteTableItem  = "write_table_item"
)

type WriteSetChange struct {
	Type   string `json:"type"`
	Raw    json.RawMessage
	Object interface{}
}

type WriteSetChangeDeleteModule struct {
	Type         string `json:"type"`
	Address      string `json:"address"`
	StateKeyHash string `json:"state_key_hash"`
	Module       string `json:"module"`
}

type WriteSetChangeDeleteResource struct {
	Type         string `json:"type"`
	Address      string `json:"address"`
	StateKeyHash string `json:"state_key_hash"`
	Resource     string `json:"resource"`
}

type WriteSetChangeDeleteTableItem struct {
	Type         string `json:"type"`
	StateKeyHash string `json:"state_key_hash"`
	Handle       string `json:"handle"`
	Key          string `json:"key"`
	// todo
}

type WriteSetChangeWriteModule struct {
	Type         string `json:"type"`
	Address      string `json:"address"`
	StateKeyHash string `json:"state_key_hash"`
	// todo
}

type WriteSetChangeWriteResource struct {
	Type         string `json:"type"`
	Address      string `json:"address"`
	StateKeyHash string `json:"state_key_hash"`
	// todo
}

type WriteSetChangeWriteTableItem struct {
	Type         string `json:"type"`
	StateKeyHash string `json:"state_key_hash"`
	Handle       string `json:"handle"`
	Key          string `json:"key"`
	Value        string `json:"value"`
	// todo
}

func (j WriteSetChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Object)
}

func (j *WriteSetChange) UnmarshalJSON(data []byte) error {
	type Aux WriteSetChange
	aux := (*Aux)(j)
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	switch j.Type {
	case DeleteModule:
		var change WriteSetChangeDeleteModule
		if err := json.Unmarshal(data, &change); err != nil {
			return err
		}
		j.Object = change
		return nil
	case DeleteResource:
		var change WriteSetChangeDeleteResource
		if err := json.Unmarshal(data, &change); err != nil {
			return err
		}
		j.Object = change
		return nil
	case DeleteTableItem:
		var change WriteSetChangeDeleteTableItem
		if err := json.Unmarshal(data, &change); err != nil {
			return err
		}
		j.Object = change
		return nil
	case WriteModule:
		var change WriteSetChangeWriteModule
		if err := json.Unmarshal(data, &change); err != nil {
			return err
		}
		j.Object = change
		return nil
	case WriteResource:
		var change WriteSetChangeWriteResource
		if err := json.Unmarshal(data, &change); err != nil {
			return err
		}
		j.Object = change
		return nil
	case WriteTableItem:
		var change WriteSetChangeWriteTableItem
		if err := json.Unmarshal(data, &change); err != nil {
			return err
		}
		j.Object = change
		return nil
	default:
		return fmt.Errorf("unsupport wirte set change type")
	}
}
