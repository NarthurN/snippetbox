package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет из snippetbox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Server is listening on http://127.0.0.1:4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
