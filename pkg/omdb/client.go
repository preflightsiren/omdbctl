package omdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// {"Title":"Hackers","Year":"1995","imdbID":"tt0113243","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BNmExMTkyYjItZTg0YS00NWYzLTkwMjItZWJiOWQ2M2ZkYjE4XkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"}
type SearchResult struct {
	ID     string `json:"imdbID,omitempty"`
	Title  string `json:"Title,omitempty"`
	Year   string `json:"Year,omitempty"`
	Poster string `json:"Poster,omitempty"`
}

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
type OMDBResponse struct {
	ID         string    `json:"imdbID,omitempty"`
	Title      string    `json:"Title,omitempty"`
	Year       string    `json:"Year,omitempty"`
	Rated      string    `json:"Rated,omitempty"`
	Released   string    `json:"Released,omitempty"`
	Runtime    string    `json:"Runtime,omitempty"`
	Genre      string    `json:"Genre,omitempty"`
	Director   string    `json:"Director,omitempty"`
	Writer     string    `json:"Writer,omitempty"`
	Actors     string    `json:"Actors,omitempty"`
	Plot       string    `json:"Plot,omitempty"`
	Language   string    `json:"Language,omitempty"`
	Country    string    `json:"Country,omitempty"`
	Awards     string    `json:"Awards,omitempty"`
	Poster     string    `json:"Poster,omitempty"`
	Ratings    []*Rating `json:"Ratings,omitempty"`
	Metascore  string    `json:"-"`
	imdbRating string    `json:"-"`
	imdbVotes  string    `json:"-"`
	Type       string    `json:"-"`
	DVD        string    `json:"-"`
	BoxOffice  string    `json:"-"`
	Production string    `json:"-"`
	Website    string    `json:"-"`
	Response   string    `json:"Response"`
	Error      string    `json:"Error,omitempty"`
}

type OMDBSearchResponse struct {
	Search       []*SearchResult `json:"Search,omitempty"`
	TotalResults string          `json:"totalResults,omitempty"`
	Response     string          `json:"Response"`
	Error        string          `json:"Error,omitempty"`
}
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	httpClient HTTPClient
	apiKey     string
}

type Error string

func (e Error) Error() string { return string(e) }

const (
	baseURL       = "http://www.omdbapi.com"
	NotFoundError = Error("Movie not found!")
)

// NewClient generates a new omdb client
func NewClient(apiKey string) *Client {
	return &Client{
		httpClient: &http.Client{},
		apiKey:     apiKey,
	}
}

// GetMovieByID returns a movie based upon its IMDB ID
func (c *Client) GetMovieByID(id string) (*Movie, error) {
	result := &OMDBResponse{}
	url := fmt.Sprintf("%s/?i=%s&apikey=%s", baseURL, id, c.apiKey)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		jsonBlob, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(jsonBlob, result)

		if err != nil {
			fmt.Println("Error decoding json")
			return nil, err
		}

		if result.Response != "True" {
			return nil, fmt.Errorf(result.Error)
		}

		return result.Movie(), nil
	}

	return nil, fmt.Errorf("Unexpected response %v", resp.StatusCode)

}

// GetMovieByTitle returns a movie based upon its title
func (c *Client) GetMovieByTitle(title string) (*Movie, error) {
	result := &OMDBResponse{}
	titleURL := fmt.Sprintf("%s/?t=%s&apikey=%s", baseURL, url.PathEscape(title), c.apiKey)
	req, err := http.NewRequest("GET", titleURL, nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		jsonBlob, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(jsonBlob, result)

		if err != nil {
			fmt.Println("Error decoding json")
			return nil, err
		}

		if result.Response != "True" {
			return nil, fmt.Errorf(result.Error)
		}

		return result.Movie(), nil
	}

	return nil, fmt.Errorf("Unexpected response %v", resp.StatusCode)

}

// GetMoviesBySearchTerm returns an array of  movies based upon a search query
func (c *Client) GetMoviesBySearchTerm(searchTerm string) ([]*Movie, error) {
	result := &OMDBSearchResponse{}
	url := fmt.Sprintf("%s/?s=%s&apikey=%s", baseURL, url.PathEscape(searchTerm), c.apiKey)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		jsonBlob, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(jsonBlob, result)

		if err != nil {
			fmt.Println("Error decoding json")
			return nil, err
		}

		if result.Response != "True" {
			return nil, fmt.Errorf(result.Error)
		}

		movies := []*Movie{}

		for _, m := range result.Search {
			movies = append(movies, m.Movie())
		}

		return movies, nil
	}

	return nil, fmt.Errorf("Unexpected response %v", resp.StatusCode)

}
