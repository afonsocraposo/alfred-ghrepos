package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	. "github.com/afonsocraposo/alfred-ghrepos/internal/types"
)

type GHRestClient struct {
	client *http.Client
	token  string
}

func (restClient *GHRestClient) Login(token string) {
	restClient.token = token
}

func (restClient *GHRestClient) request(method Method, url string) (*http.Response, error) {
	if restClient.client == nil {
		restClient.client = &http.Client{}
	}
	req, _ := http.NewRequest(string(method), url, nil)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", restClient.token))
	req.Header.Set("User-Agent", "GolangAlfredGhRepos/1.0")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	res, error := restClient.client.Do(req)
	return res, error
}

func (restClient *GHRestClient) ListAuthenticatedUserRepos() ([]Repo, error) {
	repos := []Repo{}
	url := "https://api.github.com/user/repos?per_page=100"
	repos, err := restClient.fetchAllRepos(url)
	if err != nil {
		return repos, err
	}
	return repos, nil
}

func (restClient *GHRestClient) fetchAllRepos(url string) ([]Repo, error) {
	repos := []Repo{}
	for true {
		resp, err := restClient.fetchRepos(url)
		if err != nil {
			log.Fatal(err)
		}
		r := parseRepos(resp.Body)
		repos = append(repos, r...)

		linkHeader := resp.Header.Get("link")
		if linkHeader == "" || !strings.Contains(linkHeader, "rel=\"next\"") {
			break
		}

		re := regexp.MustCompile("<([^>]+)>; rel=\"next\"")
		url = re.FindStringSubmatch(linkHeader)[1]
	}
	return repos, nil
}

func (restClient *GHRestClient) fetchRepos(url string) (*http.Response, error) {
	resp, err := restClient.request(GET, url)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode != 200 {
		return resp, fmt.Errorf("Got %d on %s request", resp.StatusCode, resp.Request.URL)
	}

	return resp, nil
}

func parseRepos(rc io.ReadCloser) []Repo {
	body, err := io.ReadAll(rc)
	if err != nil {
		log.Fatal(err)
	}

	r := []Repo{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Fatal(err)
	}
	return r
}
