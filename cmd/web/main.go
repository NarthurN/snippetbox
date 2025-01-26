package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"os"
)

func main() {
	// Конфиги через командную строку.
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")
	staticDir := flag.String("staticDir", "./ui/static", "Папка статических файлов")
	flag.Parse()

	// Логеры
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Создаем обработчик статических файлов, которые есть в ./ui/static
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(*staticDir)})
	// Переходя на URL /static/css/main.css мы должны убрать префикс /static чтобы
	// было не -> ./ui/static/static/css/main.css <- А БЫЛО -> ./ui/static/css/main.css <-
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Свой настриваемый сервер.
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	infoLog.Printf("Server is listening on %s", *addr)
	errorLog.Fatal(srv.ListenAndServe())
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
