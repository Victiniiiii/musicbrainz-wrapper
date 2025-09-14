package musicbrainzWrapper

import (
	"encoding/json"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

var artistCache = map[string]struct {
	Genre    string
	Language string
}{}

func getArtistData(artistID string) (genre string, language string) {
	if data, ok := artistCache[artistID]; ok {
		return data.Genre, data.Language
	}

	apiURL := "https://musicbrainz.org/ws/2/artist/" + artistID + "?fmt=json&inc=tags"
	resp, err := http.Get(apiURL)
	if err != nil {
		return "unknown", "unknown"
	}
	defer resp.Body.Close()

	var artist ArtistResponse
	if err := json.NewDecoder(resp.Body).Decode(&artist); err != nil {
		return "unknown", "unknown"
	}

	lang := "unknown"
	if l, ok := countryToLang[artist.Country]; ok {
		lang = l
	}

	g := "unknown"
	if len(artist.Tags) > 0 {
		sort.Slice(artist.Tags, func(i, j int) bool {
			return artist.Tags[i].Name < artist.Tags[j].Name
		})
		g = strings.ToLower(artist.Tags[0].Name)
	}

	artistCache[artistID] = struct {
		Genre    string
		Language string
	}{g, lang}

	time.Sleep(1 * time.Second)
	return g, lang
}

func DetectMetadata(title string) (artist string, genre string, language string) {
	apiURL := "https://musicbrainz.org/ws/2/recording/"
	params := url.Values{}
	params.Set("query", "recording:"+title)
	params.Set("fmt", "json")

	resp, err := http.Get(apiURL + "?" + params.Encode())
	if err != nil {
		return "unknown", "unknown", "unknown"
	}
	defer resp.Body.Close()

	var data RecordingSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "unknown", "unknown", "unknown"
	}

	if len(data.Recordings) == 0 {
		return "unknown", "unknown", "unknown"
	}

	type candidate struct {
		artist, genre, language string
		score                   int
		completeness            int
	}

	var candidates []candidate

	for _, rec := range data.Recordings {
		if len(rec.ArtistCredit) == 0 {
			continue
		}
		aName := rec.ArtistCredit[0].Name
		aID := rec.ArtistCredit[0].Artist.ID
		g, l := getArtistData(aID)

		if l == "unknown" {
			for _, r := range rec.Releases {
				if rl, ok := countryToLang[r.Country]; ok {
					l = rl
					break
				}
			}
		}

		completeness := 0
		for _, val := range []string{aName, g, l} {
			if val != "unknown" {
				completeness++
			}
		}

		candidates = append(candidates, candidate{
			artist:       aName,
			genre:        g,
			language:     l,
			score:        rec.Score,
			completeness: completeness,
		})
	}

	if len(candidates) == 0 {
		return "unknown", "unknown", "unknown"
	}

	sort.SliceStable(candidates, func(i, j int) bool {
		if candidates[i].completeness == candidates[j].completeness {
			return candidates[i].score > candidates[j].score
		}
		return candidates[i].completeness > candidates[j].completeness
	})

	best := candidates[0]
	return best.artist, best.genre, best.language
}
