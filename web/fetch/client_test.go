package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func createTestServer(fn func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(fn))
}

func createGetServer(t *testing.T) *httptest.Server {
	var attempt int32
	var sequence int32
	var lastRequest time.Time
	ts := createTestServer(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Method: %v", r.Method)
		t.Logf("Path: %v", r.URL.Path)

		if r.Method == "GET" {
			switch r.URL.Path {
			case "/":
				_, _ = w.Write([]byte("TestGet: text response"))
			case "/no-content":
				_, _ = w.Write([]byte(""))
			case "/json":
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"TestGet": "JSON response"}`))
			case "/json-invalid":
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte("TestGet: Invalid JSON"))
			case "/long-text":
				_, _ = w.Write([]byte("TestGet: text response with size > 30"))
			case "/long-json":
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"TestGet": "JSON response with size > 30"}`))
			case "/mypage":
				w.WriteHeader(http.StatusBadRequest)
			case "/mypage2":
				_, _ = w.Write([]byte("TestGet: text response from mypage2"))
			case "/set-retrycount-test":
				attp := atomic.AddInt32(&attempt, 1)
				if attp <= 4 {
					time.Sleep(time.Second * 6)
				}
				_, _ = w.Write([]byte("TestClientRetry page"))
			case "/set-retrywaittime-test":
				// Returns time.Duration since last request here
				// or 0 for the very first request
				if atomic.LoadInt32(&attempt) == 0 {
					lastRequest = time.Now()
					_, _ = fmt.Fprint(w, "0")
				} else {
					now := time.Now()
					sinceLastRequest := now.Sub(lastRequest)
					lastRequest = now
					_, _ = fmt.Fprintf(w, "%d", uint64(sinceLastRequest))
				}
				atomic.AddInt32(&attempt, 1)

			case "/set-retry-error-recover":
				w.Header().Set(hdrContentTypeKey, "application/json; charset=utf-8")
				if atomic.LoadInt32(&attempt) == 0 {
					w.WriteHeader(http.StatusTooManyRequests)
					_, _ = w.Write([]byte(`{ "message": "too many" }`))
				} else {
					_, _ = w.Write([]byte(`{ "message": "hello" }`))
				}
				atomic.AddInt32(&attempt, 1)
			case "/set-timeout-test-with-sequence":
				seq := atomic.AddInt32(&sequence, 1)
				time.Sleep(time.Second * 2)
				_, _ = fmt.Fprintf(w, "%d", seq)
			case "/set-timeout-test":
				time.Sleep(time.Second * 6)
				_, _ = w.Write([]byte("TestClientTimeout page"))
			case "/get-method-payload-test":
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					t.Errorf("Error: could not read get body: %s", err.Error())
				}
				_, _ = w.Write(body)
			case "/host-header":
				_, _ = w.Write([]byte(r.Host))
			}

			switch {
			case strings.HasPrefix(r.URL.Path, "/v1/users/sample@sample.com/100002"):
				if strings.HasSuffix(r.URL.Path, "details") {
					_, _ = w.Write([]byte("TestGetPathParams: text response: " + r.URL.String()))
				} else {
					_, _ = w.Write([]byte("TestPathParamURLInput: text response: " + r.URL.String()))
				}
			}

		}
	})

	return ts
}
