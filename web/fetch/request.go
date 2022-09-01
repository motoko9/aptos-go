package fetch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/motoko9/aptos-go/common/reflectutil"
	"github.com/motoko9/aptos-go/web"
	"io"
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

	// used for request tracing
	traceId string
	values  map[string]interface{}
	// used for `REQUEST` shortcut
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

func (r *Request) SetTrace(traceId string) *Request {
	r.traceId = traceId
	return r
}

func (r *Request) SetURL(url string) *Request {
	r.URL = url
	return r
}

func (r *Request) SetJSONBody(body interface{}) *Request {
	r.Body = body
	r.SetHeader(web.HdrContentType, web.ApplicationJSON)
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

func (r *Request) Execute() (*Response, error) {
	return r.client.execute(r)
}

func (r *Request) fmtBodyString(sl int64) (body string) {
	body = "***** NO CONTENT *****"
	if !isPayloadSupported(r.Method) {
		return
	}

	if _, ok := r.Body.(io.Reader); ok {
		body = "***** BODY IS io.Reader *****"
		return
	}

	// request body data
	if r.Body == nil {
		return
	}
	var prtBodyBytes []byte
	var err error

	contentType := r.Header.Get(hdrContentTypeKey)
	kind := reflectutil.KindOf(r.Body)
	if canJSONMarshal(contentType, kind) {
		prtBodyBytes, err = json.MarshalIndent(&r.Body, "", " ")
	} else if b, ok := r.Body.(string); ok {
		if IsJSONType(contentType) {
			bodyBytes := []byte(b)
			out := acquireBuffer()
			defer releaseBuffer(out)
			if err = json.Indent(out, bodyBytes, "", "   "); err == nil {
				prtBodyBytes = out.Bytes()
			}
		} else {
			body = b
		}
	} else if b, ok := r.Body.([]byte); ok {
		body = fmt.Sprintf("***** BODY IS byte(s) (size - %d) *****", len(b))
		return
	}

	if prtBodyBytes != nil && err == nil {
		body = string(prtBodyBytes)
	}

	if len(body) > 0 {
		bodySize := int64(len([]byte(body)))
		if bodySize > sl {
			body = fmt.Sprintf("***** REQUEST TOO LARGE (size - %d) *****", bodySize)
		}
	}

	return
}

func (r *Request) initValuesMap() {
	if r.values == nil {
		r.values = make(map[string]interface{})
	}
}
