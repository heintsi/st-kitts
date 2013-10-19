package server

import (
	"net/http"
)

func Start() {
	http.HandleFunc("/", IndexHandle)
	http.HandleFunc("/game", GameDataHandle)
	http.ListenAndServe(":8080", nil)
}
