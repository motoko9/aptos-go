package rpc

import (
	"context"
	"encoding/json"
	"fmt"
)

type Account struct {
	SequenceNumber    uint64 `json:"sequence_number,string"`
	AuthenticationKey string `json:"authentication_key"`
}

func (cl *Client) Account(ctx context.Context, address string, version uint64) (*Account, error) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var account Account
	code, err := cl.Get(ctx, "/accounts/"+address, params, &account)
	if err != nil || code != 200 {
		return nil, err
	}
	return &account, nil
}

type AccountResource struct {
	T    string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type AccountResources []AccountResource

func (cl *Client) AccountResources(ctx context.Context, address string, version uint64) (*AccountResources, error) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var accountResources AccountResources
	code, err := cl.Get(ctx, "/accounts/"+address+"/resources", params, &accountResources)
	if err != nil || code != 200 {
		return nil, err
	}
	return &accountResources, nil
}

func (cl *Client) AccountResourceByAddressAndType(ctx context.Context, address string, t string, version uint64) (*AccountResource, error) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var accountResource AccountResource
	code, err := cl.Get(ctx, "/accounts/"+address+"/resource/"+t, params, &accountResource)
	if err != nil || code != 200 {
		return nil, err
	}
	return &accountResource, nil
}

type Param struct {
	Constraints []string `json:"constraints"`
}

type Function struct {
	Name              string   `json:"name"`
	Visibility        string   `json:"visibility"`
	IsEntry           bool     `json:"is_entry"`
	GenericTypeParams []*Param `json:"generic_type_params"`
	Params            []string `json:"params"`
	Return            []string `json:"return"`
}

type Struct struct {
	Name              string              `json:"name"`
	IsNative          bool                `json:"is_native"`
	Abilities         []string            `json:"abilities"`
	GenericTypeParams []*Param            `json:"generic_type_params"`
	Fields            []map[string]string `json:"fields"`
}

type Abi struct {
	Address          string        `json:"address"`
	Name             string        `json:"name"`
	Friends          []interface{} `json:"friends"`
	ExposedFunctions []*Function   `json:"exposed_functions"`
	Structs          []*Struct     `json:"structs"`
}

type AccountModule struct {
	ByteCode string `json:"bytecode"`
	Abi      Abi    `json:"abi"`
}

type AccountModules []AccountModule

func (cl *Client) AccountModules(ctx context.Context, address string, version uint64) (*AccountModules, error) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var accountModules AccountModules
	code, err := cl.Get(ctx, "/accounts/"+address+"/modules", params, &accountModules)
	if err != nil || code != 200 {
		return nil, err
	}
	return &accountModules, nil
}

func (cl *Client) AccountModuleByAddressAndName(ctx context.Context, address string, name string, version uint64) (*AccountModule, error) {
	var params map[string]string
	if version != 0 {
		params = make(map[string]string)
		params["version"] = fmt.Sprintf("%d", version)
	}
	var accountModule AccountModule
	code, err := cl.Get(ctx, "/accounts/"+address+"/module/"+name, params, &accountModule)
	if err != nil || code != 200 {
		return nil, err
	}
	return &accountModule, nil
}
