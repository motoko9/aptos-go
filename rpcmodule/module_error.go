package rpcmodule

const (
	ClientError              = "client_error"
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

type AptosError struct {
	Message     string `json:"message"`
	ErrorCode   string `json:"error_code"`
	VmErrorCode int64  `json:"vm_error_code"`
}

func AptosErrorFromError(err error) *AptosError {
	aptosErr := &AptosError{
		Message:     err.Error(),
		ErrorCode:   "client_error",
		VmErrorCode: 0,
	}
	return aptosErr
}
