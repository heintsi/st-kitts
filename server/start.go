package server

import (
	"net/http"
)

func Start() {
	http.HandleFunc("/", IndexHandle)
	http.ListenAndServe(":8080", nil)
}
