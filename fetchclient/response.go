package fetchclient

import (
	"net/http"
)

type Response struct {
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

func (r *Response) BodyBytes() []byte {
	return r.bodyBytes
}
