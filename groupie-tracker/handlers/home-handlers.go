package handlers

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var tmpl, _ = template.ParseFiles("static/index.html")
var tmpl2, _ = template.ParseFiles("static/modal.html")
var fileServer = http.FileServer(http.Dir("./static"))

func SiteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		if _, err := os.Stat("./static/index.html"); os.IsNotExist(err) {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := template.ParseFiles("./static/.html")
			if err != nil {
				http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)

			} else {
				http.ServeFile(w, r, "./static/500.html")
			}
			return
		}

		artists, err := FetchArtCards()
		if err != nil {
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			fmt.Print("Error fetching data from API")
			return
		}

		tempErr := tmpl.Execute(w, artists)
		if tempErr != nil {
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			fmt.Print("Error executing template")
			return
		}
		return
	case "/modal":
		// Get the artist ID from the URL query parameters
		id := r.URL.Query().Get("id")
		if id == "" {
			// If no artist ID is provided, render an error page
			http.Error(w, "Error 400: Bad Request", http.StatusBadRequest)
			fmt.Print("Invalid Artist ID")
			return
		}

		// Fetch the artists data with details
		artists, err := FetchArtInfo()
		if err != nil {
			// If there's an error fetching the data, render an error page
			http.Error(w, "Artist ID is required", http.StatusInternalServerError)
			fmt.Print("Error fetching data from API")
			return
		}

		// Render the "modal.html" template with the artist data
		tempErr := tmpl2.Execute(w, artists)
		if tempErr != nil {
			http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
			fmt.Print("Error executing template")
			return
		}
		return
	}

	_, err := os.Stat("./static" + r.URL.Path)
	if os.IsNotExist(err) {
		FileNotFoundHandler(w, r)
		return
	}

	fileServer.ServeHTTP(w, r)
}

func FileNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat("./static/404.html"); os.IsNotExist(err) {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	// w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, "./static/404.html")
}
