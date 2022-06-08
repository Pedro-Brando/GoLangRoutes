package main

import (
	"html/template"
	"main/pkg/models"
	"net/http"
	"strconv"
)

//rota homescreen
func (app *application) homescreen(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(rw)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(rw, err)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl.html",
		"./ui/html/base.layout.tmpl.html",
		"./ui/html/footer.partial.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(rw, err)
		return
	}
	err = ts.Execute(rw, snippets)
	if err != nil {
		app.serverError(rw, err)
		return
	}

}

//rota savestate
func (app *application) savestate(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		rw.Header().Set("Allow", "POST")
		app.clientError(rw, http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	_, err := app.snippets.Insert(title, content)
	if err != nil {
		app.serverError(rw, err)
		return
	}

	http.Redirect(rw, r, "/", http.StatusSeeOther)
}

//rota delete
func (app *application) delete(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		rw.Header().Set("Allow", "POST")
		app.clientError(rw, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.FormValue("excluir"))
	if err != nil || id < 1 {
		app.notFound(rw)
		return
	}

	_, err = app.snippets.Delete(id)
	if err != nil {
		app.serverError(rw, err)
		return
	}

	http.Redirect(rw, r, "/", http.StatusSeeOther)
}

//rota edit2
func (app *application) edit2(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		rw.Header().Set("Allow", "POST")
		app.clientError(rw, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(rw)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	_, err = app.snippets.Edit(id, title, content)
	if err != nil {
		app.serverError(rw, err)
		return
	}

	http.Redirect(rw, r, "/", http.StatusSeeOther)
}

//rota edit
func (app *application) edit(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(rw)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(rw)
		return
	} else if err != nil {
		app.serverError(rw, err)
		return
	}

	files := []string{
		"./ui/html/editar.page.tmpl.html",
		"./ui/html/base.layout.tmpl.html",
		"./ui/html/footer.partial.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(rw, err)
		return
	}
	err = ts.Execute(rw, s)
	if err != nil {
		app.serverError(rw, err)
		return
	}
}
