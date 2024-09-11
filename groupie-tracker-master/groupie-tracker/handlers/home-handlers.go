package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func SiteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		serveCustomErrorPage(w, r, http.StatusNotFound, "static/404.html")
		return
	}

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		serveCustomErrorPage(w, r, http.StatusInternalServerError, "static/500.html")
		fmt.Print("Error parsing index.html")
		return
	}

	artistData, err := FetchCards()
	if err != nil {
		serveCustomErrorPage(w, r, http.StatusInternalServerError, "static/500.html")
		fmt.Print("Error fetching data")
		return
	}
	buf := &bytes.Buffer{}
	err = tmpl.ExecuteTemplate(buf, "index.html", artistData)
	if err != nil {
		serveCustomErrorPage(w, r, http.StatusInternalServerError, "static/500.html")
		fmt.Print("Error executing template")
		return
	}

	w.Header().Set("Content-Type", "text/html")
	buf.WriteTo(w)
}

func serveCustomErrorPage(w http.ResponseWriter, r *http.Request, statusCode int, filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, http.StatusText(statusCode), statusCode)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, filePath)
}
