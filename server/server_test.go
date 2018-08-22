package server

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

func (mux *ChoseRequestMux) ServeHttp(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case mux.requestType:
		mux.Dispatch(w, r)
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}

}
