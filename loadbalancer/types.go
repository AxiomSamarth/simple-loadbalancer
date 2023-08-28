package loadbalancer

import (
	"net/http/httputil"

	"github.com/AxiomSamarth/loadbalancers/server"
)

type Loadbalancer struct {
	port            string
	servers         []server.Server
	roundRobinCount int
	reverseProxy    *httputil.ReverseProxy
}
