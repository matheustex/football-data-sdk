package football

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var defaultTestTimeout = time.Second * 1

// testServer returns an http Client, ServeMux, and Server
func testServer() (*http.Client, *http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	transport := &RewriteTransport{&http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}}
	client := &http.Client{Transport: transport}
	return client, mux, server
}

// RewriteTransport rewrites https requests to http to avoid TLS cert issues
// during testing.
type RewriteTransport struct {
	Transport http.RoundTripper
}

// RoundTrip rewrites the request scheme to http and calls through to the
// composed RoundTripper or if it is nil, to the http.DefaultTransport.
func (t *RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	if t.Transport == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return t.Transport.RoundTrip(req)
}

func assertMethod(t *testing.T, expectedMethod string, req *http.Request) {
	assert.Equal(t, expectedMethod, req.Method)
}

func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}

// assertDone asserts that the empty struct channel is closed before the given
// timeout elapses.
func assertDone(t *testing.T, ch <-chan struct{}, timeout time.Duration) {
	select {
	case <-ch:
		_, more := <-ch
		assert.False(t, more)
	case <-time.After(timeout):
		t.Errorf("expected channel to be closed within timeout %v", timeout)
	}
}

// assertClosed asserts that the channel is closed before the given timeout
// elapses.
func assertClosed(t *testing.T, ch <-chan interface{}, timeout time.Duration) {
	select {
	case <-ch:
		_, more := <-ch
		assert.False(t, more)
	case <-time.After(timeout):
		t.Errorf("expected channel to be closed within timeout %v", timeout)
	}
}
