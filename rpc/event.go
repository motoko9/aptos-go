package rpc

import (
    "context"
    "encoding/json"
    "github.com/motoko9/aptos-go/rpcmodule"
)

func (cl *Client) EventsByKey(ctx context.Context, key string) (*rpcmodule.Events, error) {
    resp, err := cl.fetchClient.Get("/events/" + key).Execute()
    if err != nil {
        return nil, err
    }

    var events rpcmodule.Events
    if err = json.Unmarshal(resp.BodyBytes(), &events); err != nil {
        return nil, err
    }
    return &events, nil
}

func (cl *Client) EventsByHandle(ctx context.Context, address string, handle string, field string) (*rpcmodule.Events, error) {
    resp, err := cl.fetchClient.Get("/accounts/" + address + "/events/" + handle + "/" + field).Execute()
    if err != nil {
        return nil, err
    }

    var events rpcmodule.Events
    if err = json.Unmarshal(resp.BodyBytes(), &events); err != nil {
        return nil, err
    }
    return &events, nil
}
