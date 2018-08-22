package main

import (
	"fmt"
	"net/http"
	"regexp"
)

var PORT = "8080"

type ChoseRequestMux struct {
	requestType string
	routes      map[string]func(http.ResponseWriter, *http.Request)
}

func (mux *ChoseRequestMux) Dispatch(w http.ResponseWriter, r *http.Request) {
	match := false
	for pattern, f := range mux.routes {
		innerMatch, _ := regexp.MatchString(pattern, r.URL.String())
		if innerMatch {
			f(w, r)
			break
		} else {
			match = innerMatch
		}
	}
	if !match {
		http.Redirect(w, r, "", http.StatusBadRequest)
	}
}

func (mux *ChoseRequestMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case mux.requestType:
		mux.Dispatch(w, r)
	default:
		//fmt.Fprint(w, r.URL.String())
		http.Redirect(w, r, "", http.StatusNotFound)
	}
}

func main() {
	fmt.Print("Launching server")
	var AppRoutes = make(map[string]func(http.ResponseWriter, *http.Request))
	AppRoutes["/"] = func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hello world") }

	mux := &ChoseRequestMux{requestType: http.MethodPost, routes: AppRoutes}
	http.ListenAndServe(":"+PORT, mux)
}
