# Simple Loadbalancer

As I started exploring Distributed Systems, one of the important and interesting concepts I learnt is Loadbalancer. This repository is the simple implementation of roundrobin based Loadbalancer.

## Idea

```
type Loadbalancer struct{
    port string
    reverseProxy *httputil.ReverseProxy
}
```

A `Loadbalancer` type is defined with a `port` and `reverseProxy`. This `Loadbalancer` will distribute the incoming traffic to the backend servers registered to it via its `reverseProxy` in roundrobin fashion. The client will never know which server is actually serving it but it will have its response proxied by this `Loadbalancer`.

### What is Reverse Proxy?
A reverse proxy is a server or service that sits between client devices and backend servers. Its primary function is to handle client requests and distribute them to appropriate backend servers, often to provide several functions such as load balancing, caching, security, and more.

**Hence, technically Loadbalancer is a kind of Reverse Proxy.**

## Usage

1. Clone the repository

```
git clone https://github.com/AxiomSamarth/simple-loadbalancer.git
```

2. Create the docker image from the Dockerfile present
```
cd simple-loadbalancer
docker build -t loadbalancer .
```

3. Run the docker image
```
docker run -p 8080:8080 loadbalancer:latest
```

4. In another terminal, upon firing requests to `localhost:8080` (which is the loadbalancer endpoint exposed to the outside world), the loadbalancing can be seen in round robin fashion from the output something like in the previous terminal

```
forwarding requests to server localhost:8000
forwarding requests to server localhost:8001
forwarding requests to server localhost:8002
forwarding requests to server localhost:8000
forwarding requests to server localhost:8001
forwarding requests to server localhost:8002
forwarding requests to server localhost:8000
forwarding requests to server localhost:8001
forwarding requests to server localhost:8002
```