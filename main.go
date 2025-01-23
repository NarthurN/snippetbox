package main

import (
	"log"
	"net/http"
)

// home - обработчик главной страницы "/"
func home(w http.ResponseWriter, r *http.Request) {
	// Путь "/" - многоуровневый. Поэтому home обрабатывает любой путь.
	// Чтобы такого не было добавим проверку if.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Привет из snippetbox!\n"))
}

// showSnippet - показывает заметку по адресу "/snippet"
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Заметко о ...\n"))
}

// createSnippet - создание заметки "/snippet/create"
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Все методы кроме POST надо запретить и вызвать ошибку 405 (метод запрещён)
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Разрешен только POST-метод\n"))
		return
	}

	w.Write([]byte("Создание заметки ... \n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Server is listening on http://127.0.0.1:4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
