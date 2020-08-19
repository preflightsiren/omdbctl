package omdbctl

import (
	"fmt"
	"regexp"

	"../../../pkg/omdb"
)

const ()

type OMDBCTL struct {
	Client *omdb.Client
}

func NewOMDB(apikey string) *OMDBCTL {
	return &OMDBCTL{
		Client: omdb.NewClient(apikey),
	}
}

func (o *OMDBCTL) Search(input string) {
	if IsID(input) {
		fmt.Printf("Looking up Movie by ID: %s\n", input)
		m, err := o.Client.GetMovieByID(input)
		if err != nil {
			fmt.Printf("An error occured: %v\n", err)
			return
		}
		fmt.Println(m)
		return
	}

	fmt.Printf("Looking up Movie by title: %s\n", input)
	m, err := o.Client.GetMovieByTitle(input)
	if err == nil {
		fmt.Println(m)
		return
	}

	if err == omdb.NotFoundError {
		m, err := o.Client.GetMoviesBySearchTerm(input)
		if err != nil {
			println(m)
		}
	}

	fmt.Printf("An error occured: %v\n", err)
	return
}

func IsID(input string) bool {
	titleRegex, _ := regexp.Compile(`tt[\d+]`)

	return titleRegex.MatchString(input)
}
