package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayRoute(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayRoute(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "helloMyroute")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
