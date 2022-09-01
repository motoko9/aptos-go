// Copyright (c) 2015-2021 Jeevanandam M (jeeva@myjeeva.com), All rights reserved.
// resty source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.
// Most of the codes are copy from https://github.com/go-resty/resty/blob/master/retry.go

package fetch

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	defaultMaxRetries  = 0
	defaultWaitTime    = 100 * time.Millisecond
	defaultMaxWaitTime = 3 * time.Second
)

type (
	Option func(*Options)

	RetryConditionFunc func(*Response, error) bool

	// RetryAfterFunc returns time to wait before retrying.
	// https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html
	RetryAfterFunc func(ClientInterface, *Response) (time.Duration, error)
)

type Options struct {
	maxRetries      uint
	waitTime        time.Duration
	maxWaitTime     time.Duration
	retryConditions []RetryConditionFunc
	retryAfter      RetryAfterFunc
}

// RetryableClient is a wrapper of Client that supports retry.
type RetryableClient struct {
	// Options of retry logic
	retryCount       uint
	retryWaitTime    time.Duration
	retryMaxWaitTime time.Duration
	retryConditions  []RetryConditionFunc
	retryAfter       RetryAfterFunc

	client *Client
}

func NewRetryableClient(client *Client) *RetryableClient {
	return &RetryableClient{
		client:           client,
		retryCount:       defaultMaxRetries, // default is 0
		retryWaitTime:    defaultWaitTime,
		retryMaxWaitTime: defaultMaxWaitTime,
	}
}

func (c *RetryableClient) execute(r *Request) (*Response, error) {
	var (
		resp *Response
		err  error
	)

	var attempt uint = 0

	if c.retryCount == 0 {
		// do not retry
		attempt = 1
		resp, err = c.client.execute(r)
		return resp, unwrapNoRetryErr(err)
	}

	err = exponentialBackoff(
		func() (*Response, error) {
			attempt++
			resp, err = c.client.execute(r)
			if err != nil {
				r.client.log().Warn(fmt.Sprintf("%v, Attempt %v", err, attempt))
			}
			return resp, err
		},
		Retries(c.retryCount),
		WaitTime(c.retryWaitTime),
		MaxWaitTime(c.retryMaxWaitTime),
		RetryConditions(c.retryConditions),
		RetryAfter(c.retryAfter),
	)

	return resp, err
}

func (c *RetryableClient) log() hclog.Logger {
	return c.client.logger
}

func (c *RetryableClient) Get(url string) *Request {
	return NewRequest("GET", c).SetURL(url)
}

func (c *RetryableClient) Post(url string) *Request {
	return NewRequest("POST", c).SetURL(url)
}

func (c *RetryableClient) Do(method, url string) *Request {
	return NewRequest(method, c).SetURL(url)
}

func (c *RetryableClient) OnBeforeRequest(m RequestMiddleware) *RetryableClient {
	c.client.OnBeforeRequest(m)
	return c
}

func (c *RetryableClient) OnBeforeResponse(m ResponseMiddleware) *RetryableClient {
	c.client.OnBeforeResponse(m)
	return c
}

// WithRetryCount sets the retry count for the request, it defaults to 0.
// Exponential backoff algorithm is used for retry.
func (c *RetryableClient) WithRetryCount(retryCount uint) *RetryableClient {
	c.retryCount = retryCount
	return c
}

// WithRetryWaitTime method sets default wait time to sleep before retrying
// request.
//
// Default is 100 milliseconds.
func (c *RetryableClient) WithRetryWaitTime(retryWaitTime time.Duration) *RetryableClient {
	c.retryWaitTime = retryWaitTime
	return c
}

// WithRetryMaxWaitTime method sets max wait time to sleep before retrying
// request.
//
// Default is 3 seconds. And it will set to infinite(maxInt) if it is negative.
func (c *RetryableClient) WithRetryMaxWaitTime(retryMaxWaitTime time.Duration) *RetryableClient {
	c.retryMaxWaitTime = retryMaxWaitTime
	return c
}

func (c *RetryableClient) WithRetryAfter(retryAfter RetryAfterFunc) *RetryableClient {
	c.retryAfter = retryAfter
	return c
}

// AddRetryCondition method adds a retry condition function to array of functions
// that are checked to determine if the request is retried. The request will
// retry if any of the functions return true and error is nil.
func (c *RetryableClient) AddRetryCondition(rc RetryConditionFunc) *RetryableClient {
	c.retryConditions = append(c.retryConditions, rc)
	return c
}

// Fallback methods

// Retries sets the max number of retries
func Retries(value uint) Option {
	return func(o *Options) {
		o.maxRetries = value
	}
}

// WaitTime sets the default wait time to sleep between requests
func WaitTime(value time.Duration) Option {
	return func(o *Options) {
		o.waitTime = value
	}
}

// MaxWaitTime sets the max wait time to sleep between requests
func MaxWaitTime(value time.Duration) Option {
	return func(o *Options) {
		o.maxWaitTime = value
	}
}

// RetryConditions sets the conditions that will be checked for retry
func RetryConditions(conditions []RetryConditionFunc) Option {
	return func(o *Options) {
		o.retryConditions = conditions
	}
}

// RetryAfter sets the retry-after callback function
func RetryAfter(fn RetryAfterFunc) Option {
	return func(o *Options) {
		o.retryAfter = fn
	}
}

func exponentialBackoff(operation func() (*Response, error), options ...Option) error {
	// default options
	opts := Options{
		maxRetries:      defaultMaxRetries,
		waitTime:        defaultWaitTime,
		maxWaitTime:     defaultMaxWaitTime,
		retryConditions: []RetryConditionFunc{},
	}

	for _, option := range options {
		option(&opts)
	}

	var (
		resp *Response
		err  error
	)

	// total request = retry count
	for attempt := uint(0); attempt < opts.maxRetries; attempt++ {
		resp, err = operation()
		var ctx context.Context
		if resp != nil && resp.Request.ctx != nil {
			ctx = resp.Request.ctx
		} else {
			ctx = context.Background()
		}
		if ctx.Err() != nil {
			return err
		}
		err1 := unwrapNoRetryErr(err)           // raw error, it used for return users callback.
		needsRetry := err != nil && err == err1 // retry on a few operation errors by default
		for _, condition := range opts.retryConditions {
			needsRetry = condition(resp, err1)
			if needsRetry {
				break
			}
		}

		if !needsRetry {
			return err
		}

		waitTime, err2 := sleepDuration(resp, opts.waitTime, opts.maxWaitTime, opts.retryAfter, attempt)
		if err2 != nil {
			if err == nil {
				err = err2
			}
			return err
		}

		select {
		case <-time.After(waitTime):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return err
}

func sleepDuration(resp *Response, min, max time.Duration, retryAfter RetryAfterFunc, attempt uint) (time.Duration, error) {
	const maxInt = 1<<31 - 1 // max int for arch 386
	if max < 0 {
		max = maxInt
	}
	if resp == nil {
		return jitterBackoff(min, max, attempt), nil
	}

	// Check for custom callback
	if retryAfter == nil {
		return jitterBackoff(min, max, attempt), nil
	}

	result, err := retryAfter(resp.Request.client, resp)
	if err != nil {
		return 0, err // i.e. 'API quota exceeded'
	}
	if result == 0 {
		return jitterBackoff(min, max, attempt), nil
	}
	if result < 0 || max < result {
		result = max
	}
	if result < min {
		result = min
	}
	return result, nil
}

// Return capped exponential backoff with jitter
// http://www.awsarchitectureblog.com/2015/03/backoff.html
func jitterBackoff(min, max time.Duration, attempt uint) time.Duration {
	base := float64(min)
	capLevel := float64(max)

	temp := math.Min(capLevel, base*math.Exp2(float64(attempt)))
	ri := time.Duration(temp / 2)
	result := randDuration(ri)

	if result < min {
		result = min
	}

	return result
}

var rnd = newRnd()
var rndMu sync.Mutex

func randDuration(center time.Duration) time.Duration {
	rndMu.Lock()
	defer rndMu.Unlock()

	var ri = int64(center)
	var jitter = rnd.Int63n(ri)
	return time.Duration(math.Abs(float64(ri + jitter)))
}

func newRnd() *rand.Rand {
	var seed = time.Now().UnixNano()
	var src = rand.NewSource(seed)
	return rand.New(src)
}

type noRetryErr struct {
	err error
}

func (e *noRetryErr) Error() string {
	return e.err.Error()
}

func wrapNoRetryErr(err error) error {
	if err != nil {
		err = &noRetryErr{err: err}
	}
	return err
}

func unwrapNoRetryErr(err error) error {
	if e, ok := err.(*noRetryErr); ok {
		err = e.err
	}
	return err
}
