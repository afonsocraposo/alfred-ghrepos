package main

import (
	"fmt"
	"os"

	"github.com/afonsocraposo/alfred-ghrepos/internal/alfred"
	"github.com/afonsocraposo/alfred-ghrepos/internal/helpers"
)

func main() {
	args := os.Args[1:]
	var query string
	if len(args) < 1 {
		query = ""
	} else {
		query = args[0]
	}

	if query == ":refresh" {
		helpers.RefreshRepos()
		return
	}

	repos := helpers.GetRepos()
	filteredRepos := helpers.FilterRepos(repos, query)
	items := alfred.ReposToJsonItems(filteredRepos)
	fmt.Println(items)
}
