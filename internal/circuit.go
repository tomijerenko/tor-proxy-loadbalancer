package loadbalancer

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/proxy"
)

type TargetCircuit struct {
	address string
	client  *http.Client
}

func (ts *TargetCircuit) Initialize() {
	dialer, err := proxy.SOCKS5("tcp", ts.address, nil, proxy.Direct)
	if err != nil {
		fmt.Println(err)
	}

	ts.client = &http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			Dial:                  dialer.Dial,
			MaxIdleConns:          100,
			IdleConnTimeout:       5 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			ForceAttemptHTTP2:     false,
			DisableCompression:    true,
		},
	}
}
