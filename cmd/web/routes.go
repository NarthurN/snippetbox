package main

import "net/http"

func (app *application) routes(staticDir *string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// Создаем обработчик статических файлов, которые есть в ./ui/static
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(*staticDir)})
	// Переходя на URL /static/css/main.css мы должны убрать префикс /static чтобы
	// было не -> ./ui/static/static/css/main.css <- А БЫЛО -> ./ui/static/css/main.css <-
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
