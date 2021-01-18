package client

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// HTTPWrapper - Wrapper struct around base http client
type HTTPWrapper struct {
	debug  bool
	client *http.Client
}

// NewClient - Create new client with no SSL verification
func NewClient(debug bool) *HTTPWrapper {
	// Creating HTTP Client with SSL support - Its Secure but we'll skip cert verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Dial:            dialTimeout,
	}
	return NewClientFromTransport(tr, debug)
}

// NewClientFromTransport - Creates new HTTP client with transport
func NewClientFromTransport(transport http.RoundTripper, debug bool) *HTTPWrapper {
	httpWrapper := &HTTPWrapper{}
	httpWrapper.debug = debug
	httpWrapper.client = &http.Client{Transport: transport}
	return httpWrapper
}

// dialTimeout - Timeouts the connecton if the URL is not resolved
func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, time.Duration(30*time.Second))
}
