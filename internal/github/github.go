package github

import (
	"os"
	"github.com/afonsocraposo/alfred-ghrepos/internal/rest"
	. "github.com/afonsocraposo/alfred-ghrepos/internal/types"
)

func FetchRepos() []Repo {
	client := rest.GHRestClient{}
    token := os.Getenv("GH_TOKEN")
	client.Login(token)

	repos, _ := client.ListAuthenticatedUserRepos()

    return repos
}
