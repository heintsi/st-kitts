package server

import (
	"fmt"
	"net/http"
)

// FIXME Dummy handlers for initial needs
func PlayerHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nothing in %s yet", r.URL.Path[1:])
}

func GameDataHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nothing in %s yet", r.URL.Path[1:])
}

func MapHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nothing in %s yet", r.URL.Path[1:])
}

func SubmitHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nothing in %s yet", r.URL.Path[1:])
}
