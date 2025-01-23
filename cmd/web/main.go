package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Создаем обработчик статических файлов, которые есть в ./ui/static
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./static")})
	// Переходя на URL /static/css/main.css мы должны убрать префикс /static чтобы
	// было не -> ./ui/static/static/css/main.css <- А БЫЛО -> ./ui/static/css/main.css <-
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Server is listening on http://127.0.0.1:4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
		}
	}
	return f, nil
}
