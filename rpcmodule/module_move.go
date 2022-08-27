package rpcmodule

import (
	"encoding/json"
)

type MoveStructField struct {
	Name string `json:"name""`
	Type string `json:"type"`
}

type MoveStructGenericTypeParam struct {
	Constraints []string
}

type MoveStruct struct {
	Name              string                       `json:"name"`
	IsNative          bool                         `json:"is_native"`
	Abilities         []string                     `json:"abilities"`
	GenericTypeParams []MoveStructGenericTypeParam `json:"generic_type_params"`
	Fields            []MoveStructField            `json:"fields"`
}

type MoveFunctionGenericTypeParam struct {
	Constraints []string
}

type MoveFunction struct {
	Name              string                         `json:"name"`
	Visibility        string                         `json:"visibility"`
	IsEntry           bool                           `json:"is_entry"`
	GenericTypeParams []MoveFunctionGenericTypeParam `json:"generic_type_params"`
	Params            []string                       `json:"params"`
	Return            []string                       `json:"return"`
}

type MoveAbi struct {
	Address          string         `json:"address"`
	Name             string         `json:"name"`
	Friends          []string       `json:"friends"`
	ExposedFunctions []MoveFunction `json:"exposed_functions"`
	Structs          []MoveStruct   `json:"structs"`
}

type MoveResources []MoveResource

type MoveResource struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type MoveModules []MoveModule

type MoveModule struct {
	ByteCode string  `json:"bytecode"`
	Abi      *MoveAbi `json:"abi,omitempty"`
}

type MoveCode struct {
	ByteCode string  `json:"bytecode"`
	Abi      *MoveAbi `json:"abi,omitempty"`
}
