package fetch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/motoko9/aptos-go/web"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"sync"
)

var (
	hdrUserAgentKey       = http.CanonicalHeaderKey("User-Agent")
	hdrAcceptKey          = http.CanonicalHeaderKey("Accept")
	hdrContentTypeKey     = http.CanonicalHeaderKey("Content-Type")
	hdrContentLengthKey   = http.CanonicalHeaderKey("Content-Length")
	hdrContentEncodingKey = http.CanonicalHeaderKey("Content-Encoding")
	hdrLocationKey        = http.CanonicalHeaderKey("Location")

	jsonCheck = regexp.MustCompile(`(?i:(application|text)/(json|.*\+json|json\-.*)(;|$))`)

	bufPool = &sync.Pool{New: func() interface{} { return &bytes.Buffer{} }}
)

const traceRequestLogKey = "__fetchTraceRequestLog"

type ClientInterface interface {
	execute(r *Request) (*Response, error)
	log() hclog.Logger

	Get(url string) *Request

	Post(url string) *Request

	Do(method, url string) *Request
}

type Client struct {
	Header             http.Header
	host               string
	httpClient         *http.Client
	logger             hclog.Logger
	traceBodySizeLimit int64

	beforeRequest  []RequestMiddleware
	beforeResponse []ResponseMiddleware
}

func NewClient(logger hclog.Logger) *Client {
	return NewClientWithEndpoint("", logger)
}

// NewClientWithEndpoint create a new fetch client with the predefined host url
func NewClientWithEndpoint(endpoint string, logger hclog.Logger) *Client {
	return NewClientWithCustomHttpClient(&http.Client{}, endpoint, logger)
}

func NewClientWithCustomHttpClient(httpClient *http.Client, endpoint string, logger hclog.Logger) *Client {
	l := logger
	if l == nil {
		l = hclog.Default()
	}
	if t, ok := httpClient.Transport.(*http.Transport); ok {
		t.MaxIdleConnsPerHost = 10
	}
	return &Client{
		host:               endpoint,
		httpClient:         httpClient,
		logger:             l,
		Header:             http.Header{},
		traceBodySizeLimit: 1024,
		beforeRequest:      make([]RequestMiddleware, 0, 0),
		beforeResponse:     make([]ResponseMiddleware, 0, 0),
	}
}

// OnBeforeRequest method appends request middleware into the before request chain.
func (c *Client) OnBeforeRequest(m RequestMiddleware) *Client {
	c.beforeRequest = append(c.beforeRequest, m)
	return c
}

func (c *Client) OnBeforeResponse(m ResponseMiddleware) *Client {
	c.beforeResponse = append(c.beforeResponse, m)
	return c
}

func (c *Client) AddHeaders(header map[string]string) *Client {
	for k, v := range header {
		c.Header.Add(k, v)
	}
	return c
}

func (c *Client) NewRequest() *Request {
	return &Request{
		client: c,
	}
}

func (c *Client) Get(url string) *Request {
	return NewRequest("GET", c).SetURL(url)
}

func (c *Client) Post(url string) *Request {
	return NewRequest("POST", c).SetURL(url)
}

func (c *Client) Do(method, url string) *Request {
	return NewRequest(method, c).SetURL(url)
}

func (c *Client) createHttpRawRequest(r *Request) error {
	var body io.Reader
	if r.bodyBytes != nil {
		body = bytes.NewBuffer(r.bodyBytes)
	}

	var rawRequest *http.Request
	var err error
	if r.ctx != nil {
		rawRequest, err = http.NewRequestWithContext(r.ctx, r.Method, r.URL, body)
	} else {
		rawRequest, err = http.NewRequest(r.Method, r.URL, body)
	}

	if err != nil {
		return err
	}
	rawRequest.Header = r.Header
	r.rawRequest = rawRequest

	return nil
}

func (c *Client) execute(r *Request) (*Response, error) {

	for _, m := range c.beforeRequest {
		if err := m(c, r); err != nil {
			return nil, err
		}
	}

	err := parseRequestURL(c, r)
	if err != nil {
		return nil, err
	}
	err = parseRequestHeader(c, r)
	if err != nil {
		return nil, err
	}
	err = parseRequestBody(c, r)
	if err != nil {
		return nil, err
	}

	_ = traceRequest(c, r)

	err = c.createHttpRawRequest(r)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.httpClient.Do(r.rawRequest)

	response := &Response{
		Request:     r,
		rawResponse: rawResponse,
	}

	if err != nil {
		return response, err
	}

	defer rawResponse.Body.Close()
	respBody := rawResponse.Body
	if response.bodyBytes, err = ioutil.ReadAll(respBody); err != nil {
		return response, err
	}

	_ = traceResponse(c, response)
	err = handleError(c, response)
	if err != nil {
		return response, err
	}
	err = parseResponseBody(c, response)

	return response, err
}

func (c *Client) log() hclog.Logger {
	return c.logger
}

func parseRequestURL(c *Client, r *Request) error {
	reqURL, err := url.Parse(r.URL)
	if err != nil {
		return err
	}
	if !reqURL.IsAbs() {
		r.URL = reqURL.String()
		if len(r.URL) > 0 && r.URL[0] != '/' {
			r.URL = "/" + r.URL
		}
		reqURL, err = url.Parse(c.host + r.URL)
		if err != nil {
			return err
		}
	}
	// Adding query params
	if len(r.QueryParam) > 0 {
		if reqURL.RawQuery != "" {
			reqURL.RawQuery = reqURL.RawQuery + "&" + r.QueryParam.Encode()
		} else {
			reqURL.RawQuery = r.QueryParam.Encode()
		}
	}
	r.URL = reqURL.String()
	return nil
}

func parseRequestHeader(c *Client, r *Request) error {
	hdr := make(http.Header)
	for k := range c.Header {
		hdr[k] = append(hdr[k], c.Header[k]...)
	}

	for k := range r.Header {
		hdr.Del(k)
		hdr[k] = append(hdr[k], r.Header[k]...)
	}

	r.Header = hdr

	return nil
}

func parseRequestBody(c *Client, r *Request) (err error) {
	if r.Body != nil {
		switch v := r.Body.(type) {
		case io.Reader:
			r.bodyBytes, err = ioutil.ReadAll(v)
			if err != nil {
				return
			}
		case []byte:
			r.bodyBytes = v
		case string:
			r.bodyBytes = []byte(v)
		default:
			if IsJSONType(r.Header.Get(hdrContentTypeKey)) {
				r.bodyBytes, err = json.Marshal(v)
				if err != nil {
					return
				}
			}
		}
	}

	return nil
}

func parseResponseBody(c *Client, r *Response) (err error) {
	if r.StatusCode() == http.StatusNoContent {
		return
	}
	return
}

func traceRequest(c *Client, r *Request) error {
	if c.logger != nil && c.logger.IsTrace() {
		rl := &RequestLog{Header: copyHeaders(r.Header), Body: r.fmtBodyString(c.traceBodySizeLimit)}

		reqLog := "\n==============================================================================\n" +
			"~~~ REQUEST ~~~\n" +
			fmt.Sprintf("%s  %s\n", r.Method, r.URL) +
			fmt.Sprintf("HEADERS:\n%s\n", composeHeaders(rl.Header)) +
			fmt.Sprintf("BODY   :\n%v\n", rl.Body) +
			"------------------------------------------------------------------------------\n"

		r.initValuesMap()
		r.values[traceRequestLogKey] = reqLog
	}

	return nil
}

func traceResponse(c *Client, res *Response) error {
	if c.logger.IsTrace() {
		rl := &ResponseLog{Header: copyHeaders(res.Header()), Body: res.fmtBodyString(c.traceBodySizeLimit)}

		debugLog := res.Request.values[traceRequestLogKey].(string)
		debugLog += "~~~ RESPONSE ~~~\n" +
			fmt.Sprintf("STATUS       : %d\n", res.StatusCode()) +
			"HEADERS      :\n" +
			composeHeaders(rl.Header) + "\n"
		debugLog += fmt.Sprintf("BODY         :\n%v\n", rl.Body)
		debugLog += "==============================================================================\n"

		c.logger.Trace(debugLog)
	}

	return nil
}

func handleError(c *Client, r *Response) (err error) {
	if r.IsError() {
		c.logger.Warn("got error response", "statusCode", r.StatusCode())
		return web.ServerErrorCtor(r.StatusCode(), r.fmtBodyString(c.traceBodySizeLimit))
	}
	return nil
}

func IsJSONType(contentType string) bool {
	return jsonCheck.Match([]byte(contentType))
}
