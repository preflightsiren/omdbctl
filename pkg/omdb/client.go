package omdb

import (
	"fmt"
	"net/http"
)

type Rating struct {
	source string
	value  string
}

type Movie struct {
	ID       string
	title    string
	year     int
	rated    string
	released string
	runtime  string
	genre    []string
	director string
	writer   string
	actors   []string
	plot     string
	language []string
	country  string
	awards   string
	poster   string
	ratings  []*Rating
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	httpClient HTTPClient
	apiKey     string
}

const (
	baseURL = "http://www.omdbapi.com"
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
	titleURL := fmt.Sprintf("%s/?i=%s", baseURL, id)
	req, err := http.NewRequest("GET", titleURL, nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)
	return &Movie{}, nil
}
