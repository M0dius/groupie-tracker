package structs

type ArtistName struct {
	ID   int    `json:"id"`
	Img  string `json:"image"`
	Name string `json:"name"`
}

type ArtistInfo struct {
	ID              int                 `json:"id"`
	Img             string              `json:"image"`
	Name            string              `json:"name"`
	Members         []string            `json:"members"`
	CrtDate         int                 `json:"creationDate"`
	FrstAlbum       string              `json:"firstAlbum"`
	LocationsURL    string              `json:"locations"`
	DatesURL        string              `json:"concertDates"`
	RelationsURL    string              `json:"relations"`
	Locations       []string            `json:"locations"`
	Dates           []string            `json:"dates"`
	DatesByLocation map[string][]string `json:"datesByLocation"`
}

type Relations struct {
	Index []RelationData `json:"index"`
}

// RelationData struct represents a relation data item
type RelationData struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
type Relation struct {
}
