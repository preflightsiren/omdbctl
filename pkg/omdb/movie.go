package omdb

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// //`{"Title":"Hackers"
// "Year":"1995"
// "Rated":"PG-13"
// "Released":"15 Sep 1995"
// "Runtime":"105 min"
// "Genre":"Comedy
//  Crime
//  Drama
//  Thriller"
// "Director":"Iain Softley"
// "Writer":"Rafael Moreu"
// "Actors":"Jonny Lee Miller
//  Angelina Jolie
//  Jesse Bradford
//  Matthew Lillard"
// "Plot":"Hackers are blamed for making a virus that will capsize five oil tankers."
// "Language":"English
//  Italian
//  Japanese
//  Russian"
// "Country":"USA"
// "Awards":"N/A"
// "Poster":"https://m.media-amazon.com/images/M/MV5BNmExMTkyYjItZTg0YS00NWYzLTkwMjItZWJiOWQ2M2ZkYjE4XkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"
// "Ratings":[{"Source":"Internet Movie Database"
// "Value":"6.3/10"}
// {"Source":"Rotten Tomatoes"
// "Value":"33%"}
// {"Source":"Metacritic"
// "Value":"46/100"}]
// "Metascore":"46"
// "imdbRating":"6.3"
// "imdbVotes":"64
// 262"
// "imdbID":"tt0113243"
// "Type":"movie"
// "DVD":"24 Apr 2001"
// "BoxOffice":"N/A"
// "Production":"MGM"
// "Website":"N/A"
// "Response":"True"}`
//

type Movie struct {
	ID       string    `json:"imdbID,omitempty"`
	Title    string    `json:"Title,omitempty"`
	Year     int       `json:"Year,omitempty"`
	Rated    string    `json:"Rated,omitempty"`
	Released string    `json:"Released,omitempty"`
	Runtime  string    `json:"Runtime,omitempty"`
	Genre    []string  `json:"Genre,omitempty"`
	Director string    `json:"Director,omitempty"`
	Writer   string    `json:"Writer,omitempty"`
	Actors   []string  `json:"Actors,omitempty"`
	Plot     string    `json:"Plot,omitempty"`
	Language []string  `json:"Language,omitempty"`
	Country  string    `json:"Country,omitempty"`
	Awards   string    `json:"Awards,omitempty"`
	Poster   *url.URL  `json:"Poster,omitempty"`
	Ratings  []*Rating `json:"Ratings,omitempty"`
}

// Rating is a movie rating from a publication, for example Rotten tomatos
type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

// Movie converts a OMDBResponse to a Movie struct
func (r *OMDBResponse) Movie() *Movie {
	posterURL, _ := url.Parse(r.Poster)
	year, _ := strconv.ParseInt(r.Year, 10, 64)
	return &Movie{
		ID:       r.ID,
		Title:    r.Title,
		Year:     int(year),
		Rated:    r.Rated,
		Released: r.Released,
		Runtime:  r.Runtime,
		Genre:    strings.Split(r.Genre, ", "),
		Director: r.Director,
		Writer:   r.Writer,
		Actors:   strings.Split(r.Actors, ", "),
		Plot:     r.Plot,
		Language: strings.Split(r.Language, ", "),
		Country:  r.Country,
		Awards:   r.Awards,
		Poster:   posterURL,
		Ratings:  r.Ratings,
	}
}

func (m *Movie) String() string {
	var printStrings = []string{}
	printStrings = append(printStrings, fmt.Sprintf(`
%s (%d)
---
Plot: %s
Starring: %s
Released: %s
Rated: %s
Ratings:`,
		m.Title,
		m.Year,
		m.Plot,
		strings.Join(m.Actors, ", "),
		m.Released,
		m.Rated))
	for _, r := range m.Ratings {
		printStrings = append(printStrings, fmt.Sprintf("%s: %s", r.Source, r.Value))
	}

	return strings.Join(append(printStrings), "\n")
	// Internet Movie Database: 6.3/10
	// Rotten Tomatoes: 33%
	// Metacritic 46/100
}
