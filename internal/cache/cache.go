package cache

import (
	"encoding/json"
	"fmt"
    "path/filepath"

	. "github.com/afonsocraposo/alfred-ghrepos/internal/types"
)

func LoadRepos(path string) ([]Repo, bool) {
	repos := []Repo{}
	bytes, success := readTextBytes(path)
	if !success {
		return repos, false
	}
	err := json.Unmarshal(bytes, &repos)
	if err != nil {
		fmt.Println(err)
		return repos, false
	}
	return repos, true
}

func WriteRepos(path string, repos []Repo) bool {
    directory := filepath.Dir(path)
    success := MkdirIfNotExist(directory)
    if !success {
        fmt.Println("Failed to create directory", directory)
        return false
    }

	u, err := json.Marshal(repos)
	if err != nil {
		panic(err)
	}

	return writeTextFile(u, path)
}
