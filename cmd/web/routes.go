package main

import "net/http"

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.homescreen)

	mux.HandleFunc("/salvar", app.savestate)

	mux.HandleFunc("/excluir", app.delete)

	mux.HandleFunc("/editar", app.edit)

	mux.HandleFunc("/editarValidate", app.edit2)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return app.logRequest(secureHeaders(mux))
}
