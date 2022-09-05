package rpcmodule

import "encoding/json"

const (
	AccountNotFound          = "account_not_found"
	ResourceNotFound         = "resource_not_found"
	ModuleNotFound           = "module_not_found"
	StructFieldNotFound      = "struct_field_not_found"
	VersionNotFound          = "version_not_found"
	TransactionNotFound      = "transaction_not_found"
	TableItemNotFound        = "table_item_not_found"
	BlockNotFound            = "block_not_found"
	VersionPruned            = "version_pruned"
	BlockPruned              = "block_pruned"
	InvalidInput             = "invalid_input"
	InvalidTransactionUpdate = "invalid_transaction_update"
	SequenceNumberTooOld     = "sequence_number_too_old"
	VmError                  = "vm_error"
	HealthCheckFailed        = "health_check_failed"
	MemPoolIsFull            = "mempool_is_full"
	InternalError            = "internal_error"
	WebFrameworkError        = "web_framework_error"
	BcsNotSupported          = "bcs_not_supported"
	ApiDisabled              = "api_disabled"
)

type Error interface {
	// Satisfy the generic error interface.
	error

	// Code returns the short phrase depicting the classification of the error.
	Code() int

	// Message returns the error details message.
	Message() string
}

type ClientError struct {
	code    int    `json:"code"`
	message string `json:"message"`
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

func (ce ClientError) Error() string {
	return ce.message
}

func (ce ClientError) Code() int {
	return ce.code
}

func (ce ClientError) Message() string {
	return ce.message
}

type AptosError struct {
	Message     string `json:"message"`
	ErrorCode   string `json:"error_code"`
	VmErrorCode int64  `json:"vm_error_code"`
}

func (ae AptosError) Error() string {
	return ae.Message
}

func (ae *AptosError) SetError(code string, message string) {
	ae.ErrorCode = code
	ae.Message = message
}

func (ae *AptosError) String() string {
	err, _ := json.Marshal(ae)
	return string(err)
}

func (ae *AptosError) IsError() bool {
	return ae.ErrorCode != ""
}

func AptosErrorFromError(err error) *AptosError {
	aptosErr := &AptosError{
		Message:     err.Error(),
		ErrorCode:   "client_error",
		VmErrorCode: 0,
	}
	return aptosErr
}
