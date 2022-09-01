package rpc

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) Account(ctx context.Context, address string, version uint64) (*rpcmodule.AccountData, error) {
    params := make(map[string]string)
    if version != 0 {
        params["version"] = fmt.Sprintf("%d", version)
    }
    url := fmt.Sprintf("/accounts/%s", address)

    resp, err := cl.fetchClient.Get(url).SetQueryParams(params).Execute()
    if err != nil {
        return nil, err
    }
    var account rpcmodule.AccountData
    if err = json.Unmarshal(resp.BodyBytes(), &account); err != nil {
        return nil, err
    }
    return &account, nil
}

func (cl *Client) AccountResources(ctx context.Context, address string, version uint64) (*rpcmodule.MoveResources, error) {
    params := make(map[string]string)
    if version != 0 {
        params["version"] = fmt.Sprintf("%d", version)
    }
    resp, err := cl.fetchClient.Get("/accounts/" + address + "/resources").
        SetQueryParams(params).Execute()
    if err != nil {
        return nil, err
    }

    var moveResources rpcmodule.MoveResources
    if err = json.Unmarshal(resp.BodyBytes(), &moveResources); err != nil {
        return nil, err
    }
    return &moveResources, nil
}

func (cl *Client) AccountResourceByAddressAndType(ctx context.Context,
    address string, resourceType string, version uint64, resource interface{}) error {
    url := fmt.Sprintf("/accounts/%v/resource/%v", address, resourceType)
    params := make(map[string]string)
    if version != 0 {
        params["version"] = fmt.Sprintf("%d", version)
    }
    resp, err := cl.fetchClient.Get(url).SetQueryParams(params).Execute()
    if err != nil {
        return err
    }

    var moveResource rpcmodule.MoveResource
    if err = json.Unmarshal(resp.BodyBytes(), &moveResource); err != nil {
        return err
    }

    if err = json.Unmarshal(moveResource.Raw, resource); err != nil {
        return err
    }
    return nil
}

func (cl *Client) AccountModules(ctx context.Context, address string, version uint64) (*rpcmodule.MoveModules, error) {
    var params map[string]string
    if version != 0 {
        params = make(map[string]string)
        params["version"] = fmt.Sprintf("%d", version)
    }
    resp, err := cl.fetchClient.Get("/accounts/" + address + "/modules").SetQueryParams(params).Execute()
    if err != nil {
        return nil, err
    }

    var moveModules rpcmodule.MoveModules
    if err = json.Unmarshal(resp.BodyBytes(), &moveModules); err != nil {
        return nil, err
    }

    return &moveModules, nil
}

func (cl *Client) AccountModuleByAddressAndName(ctx context.Context, address string, name string, version uint64) (*rpcmodule.MoveModule, error) {
    var params map[string]string
    if version != 0 {
        params = make(map[string]string)
        params["version"] = fmt.Sprintf("%d", version)
    }

    resp, err := cl.fetchClient.Get("/accounts/" + address + "/module/" + name).SetQueryParams(params).Execute()
    if err != nil {
        return nil, err
    }

    var moveModule rpcmodule.MoveModule
    if err = json.Unmarshal(resp.BodyBytes(), &moveModule); err != nil {
        return nil, err
    }
    return &moveModule, nil
}
