package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)
	return &simpleServer{
		addr: addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type LoadBalancer struct {
	port string
	roundRobinCount int
	servers []Server
}

func NewLoadBalancer(port string, servers []Server) LoadBalancer{
	return &LoadBalancer{
		port: port,
		roundRobinCount: 0,
		servers: servers,
	}
}

func handleErr(err error){
	if err:=nil{
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func Address()

func (lb *LoadBalancer) getNextAvailableServer() Server{}
func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, r *http.Request){}

func main(){
	servers:=[]Server{
		newSimpleServer("https://google.com"),
		newSimpleServer("https://youtube.com"),
		newSimpleServer("https://en.wikipedia.org")
	}
	lb:=NewLoadBalancer("8000", servers)
	handleRedirect:=func(rw http.ResponseWriter, r *http.Request){
		lb.serveProxy(rw, r)
	}
	http.HandlerFunc("/", handleRedirect)
	fmt.Printf("serving requests at 'localhost:%s'\n", lb.port )
	http.ListenAndServe(":" + lb.port, nil) 
}