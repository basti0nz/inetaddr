package inetaddr

import (
	"bytes"
	"fmt"
	"net"
	"time"

	"github.com/valyala/fasthttp"
)

func HttpCheck(r HttpServer) (net.IP, error) {
	// Set timeout and retry times
	client := &fasthttp.Client{
		NoDefaultUserAgentHeader:      true, // Don't send: User-Agent: fasthttp
		MaxConnsPerHost:               100,
		ReadBufferSize:                4096, // Make sure to set this big enough that your whole request can be read at once.
		WriteBufferSize:               4096, // Same but for your response.
		ReadTimeout:                   time.Second,
		WriteTimeout:                  time.Second,
		MaxIdleConnDuration:           time.Minute,
		DisableHeaderNamesNormalizing: true, // If you set the case on your headers correctly you can enable this.
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(r.Server)

	client.Do(req, resp)
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("Response error: status code is %d", resp.StatusCode())
	}
	bodyBytes := bytes.TrimSpace(resp.Body())
	ip := net.ParseIP(string(bodyBytes[:]))
	if ip == nil {
		return nil, fmt.Errorf("Response error: not ip address %s", string(bodyBytes[:]))
	}
	return ip, nil
}
