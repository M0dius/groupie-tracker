package handlers

import (
	"encoding/json"
	"fmt"
	str "groupie-tracker/API"
	"log"
	"net/http"
)

const URL = "https://groupietrackers.herokuapp.com/api"

func FetchData(url string, target interface{}) error {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		return json.NewDecoder(response.Body).Decode(target)
	}
	return err
}

func FetchCards() ([]str.ArtistName, error) {
	// Create a slice to hold the artists data
	var artists []str.ArtistName

	// Fetch the artists data from the API
	err := FetchData(fmt.Sprint(URL, "/artists"), &artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}
