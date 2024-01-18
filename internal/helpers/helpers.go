package helpers

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/afonsocraposo/alfred-ghrepos/internal/cache"
	"github.com/afonsocraposo/alfred-ghrepos/internal/github"
	. "github.com/afonsocraposo/alfred-ghrepos/internal/types"

	"github.com/sahilm/fuzzy"
)

const CACHE_PATH = "/Library/Caches/com.runningwithcrayons.Alfred/Workflow Data/com.afonsocraposo.ghrepos/repos.json"

func getCachePath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	filepath := path.Join(dirname, CACHE_PATH)
	return filepath
}

func GetRepos() []Repo {
	cachePath := getCachePath()

	repos, success := cache.LoadRepos(cachePath)
	if success {
		return repos
	}

	return RefreshRepos()
}

func RefreshRepos() []Repo {
	cachePath := getCachePath()

	repos := github.FetchRepos()
	success := cache.WriteRepos(cachePath, repos)
	if !success {
		fmt.Println("Failed to write repos cache")
	}

	return repos
}

type frepos struct {
    repos []Repo
}

func (r frepos) String(i int) string {
	return r.repos[i].FullName
}
func (r frepos) Len() int {
	return len(r.repos)
}

func FilterRepos(repos []Repo, query string) []Repo {
	if query == "" {
		return repos
	}

    r := frepos{repos: repos}
	matches := fuzzy.FindFrom(query, r)

	// Extract the matched strings
	var filteredData []Repo
	for _, match := range matches {
		filteredData = append(filteredData, repos[match.Index])
	}
	return filteredData
}
