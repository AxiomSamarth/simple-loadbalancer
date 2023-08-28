package loadbalancer

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/AxiomSamarth/loadbalancers/server"
)

func NewLoadbalancer(port string) *Loadbalancer {
	return &Loadbalancer{
		port: port,
	}
}

func (lb *Loadbalancer) GetPort() string {
	return lb.port
}

func (lb *Loadbalancer) RegisterServer(server server.Server) error {
	if server != nil {
		lb.servers = append(lb.servers, server)
		return nil
	}
	return fmt.Errorf("failed to register server %s", server.GetAddress())
}

func (lb *Loadbalancer) GetNextAvailableServer() server.Server {
	nextAvailableServer := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !nextAvailableServer.IsAlive() {
		nextAvailableServer = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return nextAvailableServer
}

func (lb *Loadbalancer) ServeProxy(w http.ResponseWriter, r *http.Request) {
	nextavailableServer := lb.GetNextAvailableServer()
	fmt.Printf("\nforwarding requests to server %s", nextavailableServer.GetAddress())
	lb.reverseProxy = httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   nextavailableServer.GetAddress(),
	})
	lb.reverseProxy.ServeHTTP(w, r)
}
