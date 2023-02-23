package rpcmodule

import (
	"fmt"
	"strings"
)

func ExtractAddressFromType(t string) (string, *AptosError) {
	items := strings.Split(t, "::")
	if len(items) != 3 {
		return "", &AptosError{
			Message:     fmt.Sprintf("type %s is invalid", t),
			ErrorCode:   "400",
			VmErrorCode: 0,
		}
	}
	return items[0], nil
}

func ExtractFromResource(r string) (string, string, *AptosError) {
	indexStart := strings.IndexByte(r, '<')
	if indexStart == -1 {
		return "", "", &AptosError{
			Message:     fmt.Sprintf("resource type %s is invalid", r),
			ErrorCode:   "400",
			VmErrorCode: 0,
		}
	}
	indexEnd := strings.IndexByte(r, '>')
	if indexEnd == -1 {
		return "", "", &AptosError{
			Message:     fmt.Sprintf("resource type %s is invalid", r),
			ErrorCode:   "400",
			VmErrorCode: 0,
		}
	}
	if indexEnd <= indexStart || indexEnd != len(r)-1 {
		return "", "", &AptosError{
			Message:     fmt.Sprintf("resource type %s is invalid", r),
			ErrorCode:   "400",
			VmErrorCode: 0,
		}
	}
	m := r[0:indexStart]
	t := r[indexStart+1 : indexEnd]
	return m, t, nil
}
