package loadbalancer

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type LoadBalancer struct {
	port       string
	serverPool []Server
}

type Server interface {
	GetAddress() *url.URL
}

type TargetServer struct {
	address string
}

// GetAddress of target server
func (ts *TargetServer) GetAddress() *url.URL {
	url, err := url.Parse(ts.address)
	if err != nil {
		fmt.Println(err)
	}
	return url
}

// ScheduleNext target server
func (lb *LoadBalancer) ScheduleNext() Server {
	rnd := rand.Intn(len(lb.serverPool))
	return lb.serverPool[rnd]
}

// Serve to proxied target server
func (lb *LoadBalancer) Serve(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RemoteAddr)
	fmt.Println(req.Method)
	fmt.Println(req.URL)
	fmt.Println(req.Header)

	// http client.do, then copy headers and body

	proxy := httputil.NewSingleHostReverseProxy(lb.ScheduleNext().GetAddress())
	proxy.ServeHTTP(rw, req)
}

// Start load balancer
func Start(port string) {
	lb := &LoadBalancer{
		port: port,
		serverPool: []Server{
			&TargetServer{
				address: "http://localhost:8081",
			},
			&TargetServer{
				address: "http://localhost:8082",
			},
		},
	}

	http.HandleFunc("/", lb.Serve)

	fmt.Println("Serving localhost:" + lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
