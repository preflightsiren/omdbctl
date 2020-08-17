package main

import (
	"fmt"
	"os"

	"../../internal/pkg/omdbctl"
)

func main() {
	fmt.Printf("omdbctl - the Open Movie DB search tool.")
	// Validate the input
	omdb := omdbctl.NewOMDB(os.Getenv("APIKEY"))
	omdb.Search("tt0113243")
}
