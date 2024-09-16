package main

import (
	"fmt"
	hndl "groupie-tracker/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hndl.SiteHandler)
	http.HandleFunc("/modal", hndl.SiteHandler)
	fmt.Println("Starting Serer at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
