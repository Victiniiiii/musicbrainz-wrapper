package musicbrainz_wrapper

type RecordingSearchResponse struct {
	Recordings []Recording `json:"recordings"`
}

type Recording struct {
	Title        string         `json:"title"`
	Score        int            `json:"score"`
	ArtistCredit []ArtistCredit `json:"artist-credit"`
	Releases     []Release      `json:"releases"`
}

type ArtistCredit struct {
	Name   string `json:"name"`
	Artist struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"artist"`
}

type Release struct {
	Country string `json:"country"`
}

type ArtistResponse struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Tags    []Tag  `json:"tags"`
}

type Tag struct {
	Name string `json:"name"`
}
