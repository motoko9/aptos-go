package fetchclient

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	hdrUserAgentKey       = http.CanonicalHeaderKey("User-Agent")
	hdrAcceptKey          = http.CanonicalHeaderKey("Accept")
	hdrContentTypeKey     = http.CanonicalHeaderKey("Content-Type")
	hdrContentLengthKey   = http.CanonicalHeaderKey("Content-Length")
	hdrContentEncodingKey = http.CanonicalHeaderKey("Content-Encoding")
	hdrLocationKey        = http.CanonicalHeaderKey("Location")
	applicationJSON       = "application/json"
)

type ClientInterface interface {
	execute(r *Request) (*Response, error)
	Get(url string) *Request
	Post(url string) *Request
	Do(method, url string) *Request
}

type Client struct {
	Header             http.Header
	host               string
	httpClient         *http.Client
	traceBodySizeLimit int64
}

func NewClient() *Client {
	return NewClientWithEndpoint("")
}

// NewClientWithEndpoint create a new fetch client with the predefined host url
func NewClientWithEndpoint(endpoint string) *Client {
	return NewClientWithCustomHttpClient(&http.Client{}, endpoint)
}

func NewClientWithCustomHttpClient(httpClient *http.Client, endpoint string) *Client {
	if t, ok := httpClient.Transport.(*http.Transport); ok {
		t.MaxIdleConnsPerHost = 10
	}
	return &Client{
		host:               endpoint,
		httpClient:         httpClient,
		Header:             http.Header{},
		traceBodySizeLimit: 1024,
	}
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

	err = c.createHttpRawRequest(r)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.httpClient.Do(r.rawRequest)
	if err != nil {
		return nil, err
	}

	// before here, http request handle error
	// after here, http request handle ok, http server may give error response
	//
	defer rawResponse.Body.Close()
	bodyBytes, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		rawResponse: rawResponse,
		bodyBytes:   bodyBytes,
	}, nil
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
			if r.Header.Get(hdrContentTypeKey) == applicationJSON {
				r.bodyBytes, err = json.Marshal(v)
				if err != nil {
					return
				}
			}
		}
	}

	return nil
}
