package alfred

import (
	"encoding/json"

	. "github.com/afonsocraposo/alfred-ghrepos/internal/types"
)

func reposToAlfredItems(repos []Repo) map[string][]AlfredItem {
    items := make([]AlfredItem, len(repos))

    for i, repo := range repos {
        item := AlfredItem{
            Title:  repo.Name,
            Subtitle: repo.FullName,
            Arg: repo.Url,
            Autocomplete: repo.FullName,
            Icon: AlfredIcon{
                IconType: "fileicon",
                Path: repo.IconPath,
            },
        }
        items[i] = item
    }

    return map[string][]AlfredItem{"items": items}
}

func ReposToJsonItems(repos []Repo) string {
    items := reposToAlfredItems(repos)
    u, err := json.Marshal(items)
    if err != nil {
        panic(err)
    }
    return string(u)
}


