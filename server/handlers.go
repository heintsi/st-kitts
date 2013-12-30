package server

import (
	"fmt"
	"github.com/heintsi/st-kitts/game"
	"io"
	"net/http"
	"os"
	"regexp"
)

func StaticFileHandle(w http.ResponseWriter, r *http.Request) {
	if isValidStaticFileRequest(r) {
		mux.ServeHTTP(w, r)
	} else {
		serveIndexHtml(w)
	}
}

func serveIndexHtml(w http.ResponseWriter) {
	html, err := os.Open("public/index.html")
	if err != nil {
		http.Error(w, "Something went wrong.",
			http.StatusInternalServerError)
		return
	}
	defer html.Close()
	io.Copy(w, html)
}

func isValidStaticFileRequest(r *http.Request) bool {
	regex := regexp.MustCompile("(?:/[^/]*)+\\.(?:html|css|js)$")
	matches := regex.FindStringSubmatch(r.URL.Path)
	return len(matches) > 0
}

func PlayerHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nothing in %s yet", r.URL.Path[1:])
}

func GameDataHandle(w http.ResponseWriter, r *http.Request) {
	gameHash := r.URL.Path[len("/game/"):]
	if len(gameHash) > 0 {
		state := game.ExampleState(gameHash)
		io.Copy(w, state)
	} else {
		http.Error(w, "No game id provided.", http.StatusNotFound)
	}
}

func MapHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nothing in %s yet", r.URL.Path[1:])
}

func SubmitHandle(w http.ResponseWriter, r *http.Request) {
	var turn game.Turn
	if r.Method == "POST" {
		_, err = io.Copy(turn, r.Body)
		if err != nil {
			fmt.Fprint(w, "Error: %v", err)
		}
		go turn.Submit()
	} else {
		fmt.Fprintf(w, "Got %s instead of POST", r.Method)
	}
}
