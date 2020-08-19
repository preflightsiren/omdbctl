package omdb

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	hackers = &Movie{
		ID:       "tt0113243",
		Title:    "Hackers",
		Year:     1995,
		Rated:    "PG-13",
		Released: "15 Sep 1995",
		Runtime:  "105 min",
		Genre:    []string{"Comedy", "Crime", "Drama", "Thriller"},
		Director: "Iain Softley",
		Writer:   "Rafael Moreu",
		Actors:   []string{"Jonny Lee Miller", "Angelina Jolie", "Jesse Bradford", "Matthew Lillard"},
		Plot:     "Hackers are blamed for making a virus that will capsize five oil tankers.",
		Language: []string{"English", "Italian", "Japanese", "Russian"},
		Country:  "USA",
		Awards:   "N/A",
		Poster: &url.URL{
			Scheme:     "https",
			Opaque:     "",
			User:       nil,
			Host:       "m.media-amazon.com",
			Path:       "/images/M/MV5BNmExMTkyYjItZTg0YS00NWYzLTkwMjItZWJiOWQ2M2ZkYjE4XkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg",
			RawPath:    "",
			ForceQuery: false,
			RawQuery:   "",
			Fragment:   "",
		}, //"https://m.media-amazon.com/images/M/MV5BNmExMTkyYjItZTg0YS00NWYzLTkwMjItZWJiOWQ2M2ZkYjE4XkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg",
		Ratings: []*Rating{
			&Rating{
				Source: "Internet Movie Database",
				Value:  "6.3/10",
			},
			&Rating{
				Source: "Rotten Tomatoes",
				Value:  "33%",
			},
			&Rating{
				Source: "Metacritic",
				Value:  "46/100",
			},
		},
	}
)

func TestConvertOMDBResponseToMovie(t *testing.T) {
	response := &OMDBResponse{
		ID:       "tt0113243",
		Title:    "Hackers",
		Year:     "1995",
		Rated:    "PG-13",
		Released: "15 Sep 1995",
		Runtime:  "105 min",
		Genre:    "Comedy, Crime, Drama, Thriller", //[]string{"Comedy", "Crime", "Drama", "Thriller"},
		Director: "Iain Softley",
		Writer:   "Rafael Moreu",
		Actors:   "Jonny Lee Miller, Angelina Jolie, Jesse Bradford, Matthew Lillard", //[]string{"Jonny Lee Miller", "Angelina Jolie", "Jesse Bradford", "Matthew Lillard"},
		Plot:     "Hackers are blamed for making a virus that will capsize five oil tankers.",
		Language: "English, Italian, Japanese, Russian", //[]string{"English", "Italian", "Japanese", "Russian"},
		Country:  "USA",
		Awards:   "N/A",
		Poster:   "https://m.media-amazon.com/images/M/MV5BNmExMTkyYjItZTg0YS00NWYzLTkwMjItZWJiOWQ2M2ZkYjE4XkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg",
		Ratings: []*Rating{
			&Rating{
				Source: "Internet Movie Database",
				Value:  "6.3/10",
			},
			&Rating{
				Source: "Rotten Tomatoes",
				Value:  "33%",
			},
			&Rating{
				Source: "Metacritic",
				Value:  "46/100",
			},
		},
		// "Metascore":"46",
		// "imdbRating":"6.3",
		// "imdbVotes":"64,262",
		// "imdbID":"tt0113243",
		// "Type":"movie",
		// "DVD":"24 Apr 2001",
		// "BoxOffice":"N/A",
		// "Production":"MGM",
		// "Website":"N/A",
		// "Response":"True"
	}

	assert.Equal(t, hackers, response.Movie())
}

func TestPrintableMovie(t *testing.T) {
	expectedString := `Hackers (1995)
---
Plot: Hackers are blamed for making a virus that will capsize five oil tankers.
Starring: Jonny Lee Miller, Angelina Jolie, Jesse Bradford, Matthew Lillard
Released: 15 Sep 1995
Rated: PG-13
Ratings:
Internet Movie Database: 6.3/10
Rotten Tomatoes: 33%
Metacritic: 46/100
`
	assert.Equal(t, expectedString, hackers.String())
}

func TestPrintableSummary(t *testing.T) {
	expectedString := "Hackers (1995)\n"
	posterURL, _ := url.Parse("https://m.media-amazon.com/images/M/MV5BNmExMTkyYjItZTg0YS00NWYzLTkwMjItZWJiOWQ2M2ZkYjE4XkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg")
	hackers = &Movie{
		ID:     "tt0113243",
		Title:  "Hackers",
		Year:   1995,
		Poster: posterURL,
	}

	assert.Equal(t, expectedString, hackers.String())
}
