package httpclient

import (
	"bytes"
	"fmt"
	"github.com/motoko9/aptos-go/common/stringutil"
	"net/http"
	"reflect"
	"sort"
	"strings"
)

// RequestLog struct is used to collected information from resty request
// instance for debug logging. It sent to request log callback before resty
// actually logs the information.
type RequestLog struct {
	Header http.Header
	Body   string
}

// ResponseLog struct is used to collected information from resty response
// instance for debug logging. It sent to response log callback before resty
// actually logs the information.
type ResponseLog struct {
	Header http.Header
	Body   string
}

func copyHeaders(hdrs http.Header) http.Header {
	nh := http.Header{}
	for k, v := range hdrs {
		nh[k] = v
	}
	return nh
}

func sortHeaderKeys(hdrs http.Header) []string {
	keys := make([]string, 0, len(hdrs))
	for key := range hdrs {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func composeHeaders(hdrs http.Header) string {
	str := make([]string, 0, len(hdrs))
	for _, k := range sortHeaderKeys(hdrs) {
		var v string
		if stringutil.ContainsIgnoreCase(k, []string{"token", "authorization"}) {
			v = strings.TrimSpace(fmt.Sprintf("%25s: %s", k, "****"))
		} else if k == "Cookie" {
			cv := strings.TrimSpace(strings.Join(hdrs[k], ", "))
			v = strings.TrimSpace(fmt.Sprintf("%25s: %s", k, cv))
		} else {
			v = strings.TrimSpace(fmt.Sprintf("%25s: %s", k, strings.Join(hdrs[k], ", ")))
		}
		if v != "" {
			str = append(str, "\t"+v)
		}
	}
	return strings.Join(str, "\n")
}

func acquireBuffer() *bytes.Buffer {
	return bufPool.Get().(*bytes.Buffer)
}

func releaseBuffer(buf *bytes.Buffer) {
	if buf != nil {
		buf.Reset()
		bufPool.Put(buf)
	}
}

func isPayloadSupported(m string) bool {
	return !(m == http.MethodHead || m == http.MethodOptions)
}

func canJSONMarshal(contentType string, kind reflect.Kind) bool {
	return IsJSONType(contentType) && (kind == reflect.Struct || kind == reflect.Map || kind == reflect.Slice)
}