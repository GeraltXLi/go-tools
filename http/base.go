package http

import (
	"net/http"
	"time"
)

var (
	HttpClient *http.Client
)

//var DefaultTransport RoundTripper = &Transport{
//	Proxy: ProxyFromEnvironment,
//	DialContext: (&net.Dialer{
//		Timeout:   30 * time.Second,
//		KeepAlive: 30 * time.Second,
//		DualStack: true,
//	}).DialContext,
//	MaxIdleConns:          100,
//	IdleConnTimeout:       90 * time.Second,
//	TLSHandshakeTimeout:   10 * time.Second,
//	ExpectContinueTimeout: 1 * time.Second,
//}
func init() {
	HttpClient = &http.Client{
		Timeout: 3 * time.Second,
	}
}
