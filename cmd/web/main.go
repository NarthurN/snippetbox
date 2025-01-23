package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Создаем обработчик статических файлов, которые есть в ./ui/static
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// Переходя на URL /static/css/main.css мы должны убрать префикс /static чтобы
	// было не -> ./ui/static/static/css/main.css <- А БЫЛО -> ./ui/static/css/main.css <-
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Server is listening on http://127.0.0.1:4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
