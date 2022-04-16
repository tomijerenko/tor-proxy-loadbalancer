package loadbalancer

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
)

type LoadBalancer struct {
	port        string
	circuitPool []TargetCircuit
}

// GetNextClient for tor circuit
func (lb *LoadBalancer) GetNextCircuit() *TargetCircuit {
	rnd := rand.Intn(len(lb.circuitPool))
	return &lb.circuitPool[rnd]
}

// Serve to proxied target server
func (lb *LoadBalancer) Serve(rw http.ResponseWriter, req *http.Request) {
	circuit := lb.GetNextCircuit()
	req.RequestURI = ""
	removeHopHeaders(req.Header)

	if ip, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		setXForwardedFor(req.Header, ip)
	}

	response, err := circuit.client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	removeHopHeaders(response.Header)

	copyHeaders(rw.Header(), response.Header)
	rw.WriteHeader(response.StatusCode)
	io.Copy(rw, response.Body)
}

// Start load balancer
func Start(port string) {
	lb := &LoadBalancer{
		port: port,
		circuitPool: []TargetCircuit{
			TargetCircuit{
				address: "circuitzero:9050",
			},
			TargetCircuit{
				address: "circuitone:9050",
			},
			TargetCircuit{
				address: "circuittwo:9050",
			},
			TargetCircuit{
				address: "circuitthree:9050",
			},
			TargetCircuit{
				address: "circuitfour:9050",
			},
		},
	}

	for idx, _ := range lb.circuitPool {
		lb.circuitPool[idx].Initialize()
	}

	http.HandleFunc("/", lb.Serve)

	fmt.Println("Serving localhost:" + lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
