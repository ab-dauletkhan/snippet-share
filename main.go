package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// handler function to the home page
func home(w http.ResponseWriter, r *http.Request) {
	// Initially it catched all URL's, because every url includes /
	// Now, we want / to match only the home page, and other pages will be 404 Page Not found
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippet Share"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID = %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// We want to create only with POST method, otherwise we return error
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Creates multiplexer to match the URL to handlers
	mux := http.NewServeMux()
	// Register the handler function "home" to the URL "/"
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting the server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
