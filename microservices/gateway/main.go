package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"sync"
)

type User struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

func GetCurrentUser(r *http.Request) *User {
	//does some magic with our sessions package
	return &User{
		FirstName: "James",
		LastName:  "Bond",
	}
}

func NewServiceProxy(addrs []string) *httputil.ReverseProxy {
	port := 0
	mx := sync.Mutex{}
	return &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			// modify the request to indicate remote host
			user := GetCurrentUser(r)
			userJSON, err := json.Marshal(user)
			if err != nil {
				log.Printf("error marshalling user: %v", err)
			}
			r.Header.Add("X-User", string(userJSON))

			mx.Lock()
			r.URL.Host = addrs[port%len(addrs)]
			port++
			mx.Unlock()
			r.URL.Scheme = "http"
		},
	}
}

//RootHandler handles requests for the root resource
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello from the gateway! Try requesting /v1/time")
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	//TODO: get network addresses for our
	//timesvc instances
	timesvcAddrs := os.Getenv("TIMESVC_ADDRS")
	splitTimesvcAddrs := strings.Split(timesvcAddrs, ",")

	helloAddrs := os.Getenv("HELLO_ADDRS")
	splitHelloAddrs := strings.Split(helloAddrs, ",")

	nodeSvcAddrs := os.Getenv("NODESVC_ADDRS")
	splitNodeAddrs := strings.Split(nodeSvcAddrs, ",")

	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)
	//TODO: add reverse proxy handler for `/v1/time`
	mux.Handle("/v1/time", NewServiceProxy(splitTimesvcAddrs))
	mux.Handle("/v1/hello", NewServiceProxy(splitHelloAddrs))
	mux.Handle("/v1/users/me/hello", NewServiceProxy(splitNodeAddrs))
	mux.Handle("/v1/channels", NewServiceProxy(splitNodeAddrs))

	log.Printf("server is listening at https://%s...", addr)
	log.Fatal(http.ListenAndServeTLS(addr, "tls/fullchain.pem", "tls/privkey.pem", mux))
}
