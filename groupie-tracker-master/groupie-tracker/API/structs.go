package structs

type ArtistName struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
}

type ArtistInfo struct {
	ID        int      `json:"id"`
	Img       string   `json:"image"`
	Name      string   `json:"name"`
	Members   []string `json:"members"`
	CrtDate   int      `json:"creationDate"`
	FrstAlbum string   `json:"firstAlbum"`
}

// type Location struct {
// }

// type Date struct {
// }

type Relation struct {
}
