package main

import (
	"fmt"
	"os"
	"strings"

	"../../internal/pkg/omdbctl"
)

func main() {
	fmt.Printf("omdbctl - the Open Movie DB search tool.\n")
	// Validate the input
	if os.Getenv("APIKEY") == "" {
		fmt.Println("Environmant variable APIKEY not set. Please visit http://www.omdbapi.com/apikey.aspx to obtain new API keys.")
		return
	}
	if len(os.Args) == 1 {
		printHelp()
		return
	}
	omdb := omdbctl.NewOMDB(os.Getenv("APIKEY"))
	omdb.Search(strings.Join(os.Args[1:], " "))
}

func printHelp() {
	fmt.Printf(`Usage: %s [ID|title|search term]\n`, os.Args[0])
}
