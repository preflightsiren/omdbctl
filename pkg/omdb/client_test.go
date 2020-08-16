package omdb

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	json := `{"Title":"Hackers","Year":"1995","Rated":"PG-13","Released":"15 Sep 1995","Runtime":"105 min","Genre":"Comedy, Crime, Drama, Thriller","Director":"Iain Softley","Writer":"Rafael Moreu","Actors":"Jonny Lee Miller, Angelina Jolie, Jesse Bradford, Matthew Lillard","Plot":"Hackers are blamed for making a virus that will capsize five oil tankers.","Language":"English, Italian, Japanese, Russian","Country":"USA","Awards":"N/A","Poster":"https://m.media-amazon.com/images/M/MV5BNmExMTkyYjItZTg0YS00NWYzLTkwMjItZWJiOWQ2M2ZkYjE4XkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg","Ratings":[{"Source":"Internet Movie Database","Value":"6.3/10"},{"Source":"Rotten Tomatoes","Value":"33%"},{"Source":"Metacritic","Value":"46/100"}],"Metascore":"46","imdbRating":"6.3","imdbVotes":"64,262","imdbID":"tt0113243","Type":"movie","DVD":"24 Apr 2001","BoxOffice":"N/A","Production":"MGM","Website":"N/A","Response":"True"}`
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{
		StatusCode: 200,
		Body:       r,
	}, nil
}

func TestClient(t *testing.T) {
	testClient := &Client{
		httpClient: &MockClient{},
		apiKey:     "testkey",
	}

	resp, err := testClient.GetMovieByID("12345")

	fmt.Println(resp)
	fmt.Println(err)

	expectedMove := &Movie{
		title:    "Hackers",
		year:     1995,
		rated:    "PG-13",
		released: "15 Sep 1995",
		runtime:  "105 min",
		genre:    []string{"Comedy", "Crime", "Drama", "Thriller"},
		director: "Iain Softley",
		writer:   "Rafael Moreu",
		actors:   []string{"Jonny Lee Miller", "Angelina Jolie", "Jesse Bradford", "Matthew Lillard"},
		plot:     "Hackers are blamed for making a virus that will capsize five oil tankers.",
		language: []string{"English", "Italian", "Japanese", "Russian"},
		country:  "USA",
		awards:   "N/A",
		poster:   "https://m.media-amazon.com/images/M/MV5BNmExMTkyYjItZTg0YS00NWYzLTkwMjItZWJiOWQ2M2ZkYjE4XkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg",
		ratings:  nil,
	}
}
