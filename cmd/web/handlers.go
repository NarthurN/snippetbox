package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/NarthurN/snippetbox/pkg/models"
)

// home - обработчик главной страницы "/"
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Путь "/" - многоуровневый. Поэтому home обрабатывает любой путь.
	// Чтобы такого не было добавим проверку if.
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if err = ts.Execute(w, nil); err != nil {
		app.serverError(w, err)
	}
}

// showSnippet - показывает заметку по адресу "/snippet"
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%v", s)
}

// createSnippet - создание заметки "/snippet/create"
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// Все методы кроме POST надо запретить и вызвать ошибку 405 (метод запрещён)
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "aaa"
	content := "abab"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
