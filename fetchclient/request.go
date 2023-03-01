package fetchclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Request struct {
	URL        string
	Method     string
	QueryParam url.Values
	Header     http.Header

	Body      interface{}
	bodyBytes []byte // buffer for marshalled Body bytes

	rawRequest *http.Request
	ctx        context.Context

	client ClientInterface
}

func NewRequest(method string, client ClientInterface) *Request {
	return &Request{
		QueryParam: make(url.Values),
		Header:     make(http.Header),
		Method:     method,
		client:     client,
	}
}

func (r *Request) SetURL(url string) *Request {
	r.URL = url
	return r
}

func (r *Request) SetJSONBody(body interface{}) *Request {
	r.Body = body
	r.SetHeader(hdrContentTypeKey, "application/json")
	return r
}

func (r *Request) SetBody(body interface{}) *Request {
	r.Body = body
	return r
}

func (r *Request) SetHeader(header, value string) *Request {
	r.Header.Set(header, value)
	return r
}

func (r *Request) SetHeaders(headers map[string]string) *Request {
	for k, v := range headers {
		r.SetHeader(k, v)
	}
	return r
}

func (r *Request) SetQueryParam(param, value string) *Request {
	r.QueryParam.Set(param, value)
	return r
}

func (r *Request) SetQueryParams(params map[string]string) *Request {
	for p, v := range params {
		r.SetQueryParam(p, v)
	}
	return r
}

func (r *Request) WithContext(ctx context.Context) *Request {
	r.ctx = ctx
	return r
}

func (r *Request) Execute(rsp interface{}, general interface{}, err FetchError) {
	//
	type rspHeader struct {
		ChainId             string `json:"chain_id"`
		Epoch               string `json:"epoch"`
		LedgerVersion       string `json:"ledger_version"`
		OldestLedgerVersion string `json:"oldest_ledger_version"`
		BlockHeight         string `json:"block_height"`
		OldestBlockHeight   string `json:"oldest_block_height"`
		LedgerTimestamp     string `json:"ledger_timestamp"`
		Cursor              string `json:"cursor"`
	}
	var header rspHeader
	headerJson, _ := json.Marshal(header)
	json.Unmarshal(headerJson, general)

	// request
	httpRsp, internalErr := r.client.execute(r)
	// handle error
	if internalErr != nil {
		err.SetError("internal_error", internalErr.Error())
		return
	}

	// http error response
	// todo
	// need to handle http server error and aptos server error
	if httpRsp.StatusCode() >= 400 {
		// aptos server error
		if internalErr = json.Unmarshal(httpRsp.bodyBytes, err); internalErr != nil {
			err.SetError("internal_error", internalErr.Error())
			return
		}
		// http server error
		if !err.IsError() {
			err.SetError("web_framework_error", fmt.Sprintf("status code: %d", httpRsp.StatusCode()))
			return
		}
		return
	}

	// response header
	header.BlockHeight = httpRsp.Header().Get("X-APTOS-BLOCK-HEIGHT")
	header.ChainId = httpRsp.Header().Get("X-APTOS-CHAIN-ID")
	header.Epoch = httpRsp.Header().Get("X-APTOS-EPOCH")
	header.OldestLedgerVersion = httpRsp.Header().Get("X-APTOS-LEDGER-OLDEST-VERSION")
	header.LedgerTimestamp = httpRsp.Header().Get("X-APTOS-LEDGER-TIMESTAMPUSEC")
	header.LedgerVersion = httpRsp.Header().Get("X-APTOS-LEDGER-VERSION")
	header.OldestBlockHeight = httpRsp.Header().Get("X-APTOS-OLDEST-BLOCK-HEIGHT")
	header.Cursor = httpRsp.Header().Get("X-APTOS-CURSOR")
	headerJson, _ = json.Marshal(header)
	json.Unmarshal(headerJson, general)

	// http response
	if internalErr = json.Unmarshal(httpRsp.bodyBytes, rsp); internalErr != nil {
		err.SetError("internal_error", internalErr.Error())
	}
}
