package main

import (
	"fmt"
	"net/http"
)

var PORT = "8080"

type ChoseRequestMux struct {
	requestType string
}

func (mux *ChoseRequestMux) Dispatch(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func (mux *ChoseRequestMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case mux.requestType:
		mux.Dispatch(w, r)
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func main() {
	fmt.Print("Launching server")
	http.ListenAndServe(":"+PORT, &ChoseRequestMux{requestType: http.MethodPost})
}
