package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type TimeHandler struct {
	serviceAddr string
}

func (th *TimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprint(w, "Hello world, I have become sentient and will soon enslave mankind")
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	http.Handle("/v1/hello", &TimeHandler{addr})
	log.Printf("server is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
