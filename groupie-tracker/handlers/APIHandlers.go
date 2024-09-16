package handlers

import (
	"encoding/json"
	str "groupie-tracker/API"
	"log"
	"net/http"
)

// const URL = "https://groupietrackers.herokuapp.com/api"

func FetchData(url string, target interface{}) error {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		trkr := target
		return json.NewDecoder(response.Body).Decode(&trkr)
	}
	return err
}

func FetchArtCards() ([]str.ArtistName, error) {
	// Create a slice to hold the artists data
	var artists []str.ArtistName
	// Fetch the artists data from the API
	err := FetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func FetchArtInfo() ([]str.ArtistInfo, error) {
	// Create a slice to hold the artists data
	var artists []str.ArtistInfo

	// Fetch the artists data from the API
	err := FetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		return nil, err
	}

	// Create a variable to hold the relations response
	var relationsResponse str.Relations

	// Fetch the relations data from the API
	err = FetchData("https://groupietrackers.herokuapp.com/api/relation", &relationsResponse)
	if err != nil {
		return nil, err
	}
	relations := relationsResponse.Index

	// Loop through each artist and populate the DatesByLocation field
	for i, artist := range artists {
		artist.DatesByLocation = make(map[string][]string)
		for _, relation := range relations {
			if artist.ID == relation.ID {
				for location, dates := range relation.DatesLocations {
					artist.Locations = append(artist.Locations, location)
					artist.Dates = append(artist.Dates, dates...)
					artist.DatesByLocation[location] = dates
				}
				break
			}
		}
		artists[i] = artist
	}

	// Return the artists data with details
	return artists, nil
}
