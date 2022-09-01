package fetch

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/hashicorp/go-hclog"
	"github.com/motoko9/aptos-go/common/stringutil"
	"github.com/motoko9/aptos-go/common/tlsutil"
	"net"
	"net/http"
	"time"
)

type ClientBuilder struct {
	endpoint           string
	logger             hclog.Logger
	traceBodySizeLimit int64

	tlsConfig *tls.Config
	header    map[string]string
}

func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{}
}

func (cb *ClientBuilder) WithEndpoint(endpoint string) *ClientBuilder {
	cb.endpoint = endpoint
	return cb
}

func (cb *ClientBuilder) WithLogger(logger hclog.Logger) *ClientBuilder {
	cb.logger = logger
	return cb
}

func (cb *ClientBuilder) WithTraceBodySizeLimit(limit int64) *ClientBuilder {
	cb.traceBodySizeLimit = limit
	return cb
}

func (cb *ClientBuilder) WithCert(cert string) *ClientBuilder {
	if stringutil.IsBlank(cert) {
		return cb
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM([]byte(cert))

	cb.tlsConfig = &tls.Config{
		RootCAs:    pool,
		MinVersion: tls.VersionTLS12,
		MaxVersion: tls.VersionTLS13,
	}

	return cb
}

func (cb *ClientBuilder) WithTLSConfig(conf tlsutil.TLSConfig) *ClientBuilder {
	tcfg, err := tlsutil.SetupTLSConfig(conf)
	if err != nil {
		panic(err)
	}
	cb.tlsConfig = tcfg
	return cb
}

func (cb *ClientBuilder) WithHeaders(header map[string]string) *ClientBuilder {
	cb.header = header
	return cb
}

func (cb *ClientBuilder) Build() *Client {
	hc := &http.Client{}
	if cb.tlsConfig != nil {
		hc.Transport = &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig:       cb.tlsConfig,
		}
	}
	c := NewClientWithCustomHttpClient(hc, cb.endpoint, cb.logger)

	c.AddHeaders(cb.header)

	return c
}
