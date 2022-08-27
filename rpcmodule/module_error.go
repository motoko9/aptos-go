package rpcmodule

type AptosError struct {
	Message     string `json:"message"`
	ErrorCode   string `json:"error_code"`
	VmErrorCode int64  `json:"vm_error_code"`
}
