package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/movies", getMoviesHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func getMoviesHandler(w http.ResponseWriter, r *http.Request) {
	m := Movie{Title: "Memento", Director: "Christopher Nolan"}
	jsonM, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonM)
}

type Movie struct {
	Title    string `json:"title"`
	Director string `json:"director"`
}
