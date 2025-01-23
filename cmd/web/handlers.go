package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// home - обработчик главной страницы "/"
func home(w http.ResponseWriter, r *http.Request) {
	// Путь "/" - многоуровневый. Поэтому home обрабатывает любой путь.
	// Чтобы такого не было добавим проверку if.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err = ts.Execute(w, nil); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// showSnippet - показывает заметку по адресу "/snippet"
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Отображение заметки с ID %d ...\n", id)
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
