package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// handler function to the home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Initially it catched all URL's, because every url includes /
	// Now, we want / to match only the home page, and other pages will be 404 Page Not found
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID = %d", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// We want to create only with POST method, otherwise we return error
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	_, err := w.Write([]byte("Create a new snippet..."))
	if err != nil {
		app.errorLog.Println("Error while writing response: ", err)
	}
}
