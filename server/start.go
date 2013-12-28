package server

import (
	"net/http"
)

var mux = http.NewServeMux()

func Start() {
	mux.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/", StaticFileHandle)
	http.HandleFunc("/game/", GameDataHandle)
	http.HandleFunc("/submit/", SubmitHandle)
	http.ListenAndServe(":8080", nil)
}
