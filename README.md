# OMDB API

This is a commandline tool for searching and querying the Open Movie DB.

OMDBCTL provides the ability to search for titles, display movie information as well as Rotten Tomatos information.

## Usage

omdbctl <id> - will return movie results based upon id

omdbctl <title> - will return movie results based upon its title

omdbctl <search query> - will return results based upon the search results

## Building

`go build ./cmd/omdbctl`

## Structure

This project follows the https://github.com/golang-standards/project-layout structure

* Things that interact with the commandline in `/cmd`
* Things that are generic to OMDB are in `/pkg/omdb`
* Things that are specific to OMDBCTL interfacing with OMDB are in `/internal/pkg/omdbctl`