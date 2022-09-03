package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	NotParsableContent = errors.New("body cannot be parsed to JSON")
)

type Response struct {
	Request     *Request
	rawResponse *http.Response
	bodyBytes   []byte
}

// StatusCode method returns the HTTP status code for the executed request.
//	Example: 200
func (r *Response) StatusCode() int {
	if r.rawResponse == nil {
		return 0
	}
	return r.rawResponse.StatusCode
}

// Header method returns the response headers
func (r *Response) Header() http.Header {
	if r.rawResponse == nil {
		return http.Header{}
	}
	return r.rawResponse.Header
}

// IsSuccess method returns true if HTTP status code >= 200 and < 300, otherwise false.
func (r *Response) IsSuccess() bool {
	return r.StatusCode() >= 200 && r.StatusCode() < 300
}

// IsError method returns true if HTTP status code >= 400 otherwise false.
func (r *Response) IsError() bool {
	return r.StatusCode() >= 400
}

func (r *Response) Parse(v interface{}) error {
	if r.bodyBytes == nil || !IsJSONType(r.Header().Get(hdrContentTypeKey)) {
		return NotParsableContent
	}
	return json.Unmarshal(r.bodyBytes, v)
}

func (r *Response) BodyBytes() []byte {
	return r.bodyBytes
}

// String method returns the body of the server response as String.
func (r *Response) String() string {
	if r.bodyBytes == nil {
		return ""
	}
	return strings.TrimSpace(string(r.bodyBytes))
}

func (r *Response) fmtBodyString(sl int64) string {
	if r.bodyBytes != nil {
		if int64(len(r.bodyBytes)) > sl {
			return fmt.Sprintf("***** RESPONSE TOO LARGE (size - %d) *****", len(r.bodyBytes))
		}
		ct := r.Header().Get(hdrContentTypeKey)
		if IsJSONType(ct) {
			out := acquireBuffer()
			defer releaseBuffer(out)
			err := json.Indent(out, r.bodyBytes, "", " ")
			if err != nil {
				return fmt.Sprintf("*** Error: Unable to format response body - \"%s\" ***\n\nLog Body as-is:\n%s", err, r.String())
			}
			return out.String()
		}
		return r.String()
	}

	return "***** NO CONTENT *****"
}
