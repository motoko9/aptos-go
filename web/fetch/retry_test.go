package fetch

import (
	"context"
	"errors"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Test_exponentialBackoffSuccess(t *testing.T) {
	attempts := 3
	externalCounter := 0
	retryErr := exponentialBackoff(func() (*Response, error) {
		externalCounter++
		if externalCounter < attempts {
			return nil, errors.New("not yet got the number we're after")
		}

		return nil, nil
	}, Retries(3))

	assert.NoError(t, retryErr)
	assert.EqualValues(t, externalCounter, attempts)
}

func Test_exponentialBackoffNoWaitForLastRetry(t *testing.T) {
	attempts := uint(1)
	externalCounter := uint(0)
	numRetries := uint(1)

	canceledCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp := &Response{
		Request: &Request{
			ctx: canceledCtx,
			client: &RetryableClient{
				retryAfter: func(ClientInterface, *Response) (time.Duration, error) {
					return 6, nil
				},
			},
		},
	}

	retryErr := exponentialBackoff(func() (*Response, error) {
		externalCounter++
		return resp, nil
	}, RetryConditions([]RetryConditionFunc{func(response *Response, err error) bool {
		if externalCounter == attempts+numRetries {
			cancel()
		}
		return true
	}}), Retries(numRetries))

	assert.NoError(t, retryErr)
}

func Test_exponentialBackoffTenAttemptsSuccess(t *testing.T) {
	attempts := uint(10)
	externalCounter := uint(0)
	retryErr := exponentialBackoff(func() (*Response, error) {
		externalCounter++
		if externalCounter < attempts {
			return nil, errors.New("not yet got the number we're after")
		}
		return nil, nil
	}, Retries(attempts), WaitTime(4*time.Millisecond), MaxWaitTime(200*time.Millisecond))

	assert.NoError(t, retryErr)
	assert.EqualValues(t, externalCounter, attempts)
}

// Check to make sure the conditional of the retry condition is being used
func Test_conditionalBackoffCondition(t *testing.T) {
	attempts := uint(3)
	counter := uint(0)
	check := RetryConditionFunc(func(*Response, error) bool {
		return attempts != counter
	})
	retryErr := exponentialBackoff(func() (*Response, error) {
		counter++
		return nil, nil
	}, Retries(attempts), RetryConditions([]RetryConditionFunc{check}))

	assert.NoError(t, retryErr)
	assert.EqualValues(t, counter, attempts)
}

// Check to make sure that if the conditional is false we don't retry
func Test_conditionalBackoffConditionNonExecution(t *testing.T) {
	attempts := uint(3)
	counter := uint(0)

	retryErr := exponentialBackoff(func() (*Response, error) {
		counter++
		return nil, nil
	}, Retries(attempts), RetryConditions([]RetryConditionFunc{func(response *Response, err error) bool {
		return false
	}}))

	assert.NoError(t, retryErr)
	assert.NotEqualValues(t, counter, attempts)
}

// Check to make sure the functions added to add conditionals work
func TestConditionalGet(t *testing.T) {
	ts := createGetServer(t)
	defer ts.Close()
	attemptCount := 1
	externalCounter := 0

	// This check should pass on first run, and let the response through
	check := RetryConditionFunc(func(*Response, error) bool {
		externalCounter++
		return attemptCount != externalCounter
	})

	resp, err := NewRetryableClient(NewClient(hclog.L())).
		WithRetryCount(1).
		AddRetryCondition(check).
		Get(ts.URL).
		SetQueryParam("request_no", strconv.FormatInt(time.Now().Unix(), 10)).
		Execute()

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
	assert.EqualValues(t, "200 OK", resp.rawResponse.Status)
	assert.EqualValues(t, "TestGet: text response", resp.String())
	assert.EqualValues(t, attemptCount, externalCounter)
}

func Test_ClientRetryGet(t *testing.T) {
	ts := createGetServer(t)
	defer ts.Close()

	c := NewRetryableClient(NewClient(hclog.L())).
		WithRetryCount(3)
	c.client.httpClient.Timeout = 500 * time.Millisecond

	resp, err := c.Get(ts.URL + "/set-retrycount-test").Execute()
	assert.Equal(t, 0, resp.StatusCode())
	assert.Equal(t, 0, len(resp.Header()))

	assert.True(t, strings.HasPrefix(err.Error(), "Get "+ts.URL+"/set-retrycount-test") ||
		strings.HasPrefix(err.Error(), "Get \""+ts.URL+"/set-retrycount-test\""))
}

func Test_ClientRetryWaitMaxInfinite(t *testing.T) {
	ts := createGetServer(t)
	defer ts.Close()

	attempt := uint(0)

	retryCount := uint(5)
	retryIntervals := make([]uint64, retryCount)

	// Set retry wait times that do not intersect with default ones
	retryWaitTime := 100 * time.Millisecond
	retryMaxWaitTime := time.Duration(-1.0) // negative value

	c := NewRetryableClient(NewClient(hclog.L())).
		WithRetryCount(retryCount).
		WithRetryWaitTime(retryWaitTime).
		WithRetryMaxWaitTime(retryMaxWaitTime).
		AddRetryCondition(
			func(r *Response, _ error) bool {
				timeSlept, _ := strconv.ParseUint(string(r.BodyBytes()), 10, 64)
				retryIntervals[attempt] = timeSlept
				attempt++
				return true
			},
		)
	_, _ = c.Get(ts.URL + "/set-retrywaittime-test").Execute()

	assert.EqualValues(t, attempt, 5)

	// Initial attempt has 0 time slept since last request
	assert.EqualValues(t, retryIntervals[0], uint64(0))

	for i := 1; i < len(retryIntervals); i++ {
		slept := time.Duration(retryIntervals[i])
		// Ensure that client has slept some duration between
		// waitTime and maxWaitTime for consequent requests
		if slept < retryWaitTime {
			t.Errorf("Client has slept %f seconds before retry %d", slept.Seconds(), i)
		}
	}
}

func Test_ClientRetryAfterCallbackError(t *testing.T) {
	ts := createGetServer(t)
	defer ts.Close()

	attempt := uint(0)

	retryCount := uint(5)
	retryIntervals := make([]uint64, retryCount+1)

	// Set retry wait times that do not intersect with default ones
	retryWaitTime := 100 * time.Millisecond
	retryMaxWaitTime := 9 * time.Second

	retryAfter := func(client ClientInterface, resp *Response) (time.Duration, error) {
		return 0, errors.New("quota exceeded")
	}

	c := NewRetryableClient(NewClient(hclog.L())).
		WithRetryCount(retryCount).
		WithRetryWaitTime(retryWaitTime).
		WithRetryMaxWaitTime(retryMaxWaitTime).
		WithRetryAfter(retryAfter).
		AddRetryCondition(
			func(r *Response, _ error) bool {
				timeSlept, _ := strconv.ParseUint(string(r.BodyBytes()), 10, 64)
				retryIntervals[attempt] = timeSlept
				attempt++
				return true
			},
		)

	_, err := c.Get(ts.URL + "/set-retrywaittime-test").Execute()

	assert.Error(t, err)
	// 1 attempts were made
	assert.EqualValues(t, attempt, 1)
}

func Test_ClientRetryCancel(t *testing.T) {
	ts := createGetServer(t)
	defer ts.Close()

	attempt := uint(0)

	retryCount := uint(5)
	retryIntervals := make([]uint64, retryCount+1)

	// Set retry wait times that do not intersect with default ones
	retryWaitTime := time.Duration(10) * time.Second
	retryMaxWaitTime := time.Duration(20) * time.Second

	c := NewRetryableClient(NewClient(hclog.L())).
		WithRetryCount(retryCount).
		WithRetryWaitTime(retryWaitTime).
		WithRetryMaxWaitTime(retryMaxWaitTime).
		AddRetryCondition(
			func(r *Response, _ error) bool {
				timeSlept, _ := strconv.ParseUint(string(r.BodyBytes()), 10, 64)
				retryIntervals[attempt] = timeSlept
				attempt++
				return true
			},
		)

	timeout := 2 * time.Second

	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	_, _ = c.Get(ts.URL + "/set-retrywaittime-test").WithContext(ctx).Execute()

	// 1 attempts were made
	assert.EqualValues(t, attempt, 1)

	// Initial attempt has 0 time slept since last request
	assert.EqualValues(t, retryIntervals[0], uint64(0))

	// Second attempt should be interrupted on context timeout
	if time.Duration(retryIntervals[1]) > timeout {
		t.Errorf("Client didn't awake on context cancel")
	}
	cancelFunc()
}

func Test_retryDefaultTo0(t *testing.T) {
	ts := createGetServer(t)
	defer ts.Close()
	var attempt uint = 0
	c := NewRetryableClient(NewClient(hclog.L())).
		AddRetryCondition(
			func(response *Response, err error) bool {
				attempt++
				return true
			})
	c.client.httpClient.Timeout = 500 * time.Millisecond
	_, err := c.Get(ts.URL + "/set-retrycount-test").Execute()
	assert.Error(t, err)
	assert.EqualValues(t, 0, attempt)
}
