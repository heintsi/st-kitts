package server

import (
	"github.com/heintsi/st-kitts/game"
	"fmt"
	"io"
	"net/http"
	"os"
)

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	html, err := os.Open("public/index.html")
	if err != nil {
		http.Error(w, "Something went wrong.",
			http.StatusInternalServerError)
		return
	}
	defer html.Close()
	io.Copy(w, html)
}

func PlayerHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nothing in %s yet", r.URL.Path[1:])
}

func GameDataHandle(w http.ResponseWriter, r *http.Request) {
	gameHash := r.URL.Path[len("/game/"):]
	if len(gameHash) > 0 {
		state := game.ExampleState(gameHash)
		io.Copy(w,state)
	} else {
		http.Error(w, "No game id provided.", http.StatusNotFound)
	}
}

func MapHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nothing in %s yet", r.URL.Path[1:])
}

func SubmitHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nothing in %s yet", r.URL.Path[1:])
}
