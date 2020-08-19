package omdb

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockClient struct {
	StatusCode   int
	ResponseBody string
	DoFunc       func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(m.ResponseBody)))
	return &http.Response{
		StatusCode: m.StatusCode,
		Body:       r,
	}, nil
}

func TestGetMovieByID(t *testing.T) {
	testClient := &Client{
		httpClient: &MockClient{
			StatusCode:   200,
			ResponseBody: `{"Title":"Hackers","Year":"1995","Rated":"PG-13","Released":"15 Sep 1995","Runtime":"105 min","Genre":"Comedy, Crime, Drama, Thriller","Director":"Iain Softley","Writer":"Rafael Moreu","Actors":"Jonny Lee Miller, Angelina Jolie, Jesse Bradford, Matthew Lillard","Plot":"Hackers are blamed for making a virus that will capsize five oil tankers.","Language":"English, Italian, Japanese, Russian","Country":"USA","Awards":"N/A","Poster":"https://m.media-amazon.com/images/M/MV5BNmExMTkyYjItZTg0YS00NWYzLTkwMjItZWJiOWQ2M2ZkYjE4XkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg","Ratings":[{"Source":"Internet Movie Database","Value":"6.3/10"},{"Source":"Rotten Tomatoes","Value":"33%"},{"Source":"Metacritic","Value":"46/100"}],"Metascore":"46","imdbRating":"6.3","imdbVotes":"64,262","imdbID":"tt0113243","Type":"movie","DVD":"24 Apr 2001","BoxOffice":"N/A","Production":"MGM","Website":"N/A","Response":"True"}`,
		},
		apiKey: "testkey",
	}

	actualMovieResponse, err := testClient.GetMovieByID("12345")
	posterURL, _ := url.Parse("https://m.media-amazon.com/images/M/MV5BNmExMTkyYjItZTg0YS00NWYzLTkwMjItZWJiOWQ2M2ZkYjE4XkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg")

	expectedMovie := &Movie{
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
		Poster:   posterURL,
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

	assert.Nil(t, err)
	assert.Equal(t, expectedMovie, actualMovieResponse)
}

func TestErrorInMovieLookup(t *testing.T) {
	testClient := &Client{
		httpClient: &MockClient{
			StatusCode:   200,
			ResponseBody: `{"Response":"False","Error":"Incorrect IMDB ID."}`,
		},
		apiKey: "testkey",
	}

	_, err := testClient.GetMovieByID("notexist")

	expectedError := errors.New("Incorrect IMDB ID.")

	assert.Equal(t, expectedError, err)
}
