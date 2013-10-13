package server

import (
	"net/http"
)

func Start() {
	http.HandleFunc("/", GameDataHandle)
	http.ListenAndServe(":8080", nil)
}
