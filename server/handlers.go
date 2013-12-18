package server

import (
	"github.com/heintsi/st-kitts/game"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	regex := regexp.MustCompile("/([^/]*\\.[^/]*)$")
	matches := regex.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		mux.ServeHTTP(w, r)
	} else {
		html, err := os.Open("public/index.html")
		if err != nil {
			http.Error(w, "Something went wrong.",
				http.StatusInternalServerError)
			return
		}
		defer html.Close()
		io.Copy(w, html)
	}
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
	if r.Method == "POST" {
		fmt.Fprint(w, "Got POST.")
	} else {
		fmt.Fprintf(w, "Got %s instead of POST", r.Method)
	}
}
