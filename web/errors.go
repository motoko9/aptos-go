package web

import "encoding/json"

var (
	ErrResourceNotFound = ClientError{
		code:    404,
		message: "Resource not found",
	}
	ErrResourceUnauthorized = ClientError{
		code:    401,
		message: "Resource unauthorized",
	}
	ErrInvalidParam = ClientError{
		code:    400,
		message: "Invalid parameter",
	}
	ErrInternal = ServerError{
		code:    500,
		message: "Internal server error",
	}
	ErrDuplicateRequest = ClientError{
		code:    400,
		message: "Duplicate request",
	}
)

type Error interface {
	// Satisfy the generic error interface.
	error

	// Code returns the short phrase depicting the classification of the error.
	Code() int

	// Message returns the error details message.
	Message() string
}

// ServerError is a model for the server error display.
type ServerError struct {
	code    int
	message string
}

func ServerErrorCtor(code int, message string) ServerError {
	return ServerError{
		code:    code,
		message: message,
	}
}

func (e ServerError) Error() string {
	return e.message
}

func (e ServerError) Code() int {
	return e.code
}

func (e ServerError) Message() string {
	return e.message
}

func (e ServerError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    e.code,
		Message: e.message,
	})
}

// ClientError is a model for the client error display.
type ClientError struct {
	code    int
	message string
}

func ClientErrorCtor(code int, message string) ClientError {
	return ClientError{
		code:    code,
		message: message,
	}
}

// MarshalJSON marshals ClientError to []byte
func (e ClientError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    e.code,
		Message: e.message,
	})
}

func (e ClientError) Error() string {
	return e.message
}

func (e ClientError) Code() int {
	return e.code
}

func (e ClientError) Message() string {
	return e.message
}
