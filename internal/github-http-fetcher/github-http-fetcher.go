package github_http_fetcher

import (
	"net/http"
	"sync"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type GithubFetcher struct {
	Client HTTPClient
	wg     sync.WaitGroup
	token  string
}

func New(githubToken string) *GithubFetcher {
	return &GithubFetcher{
		Client: &http.Client{},
		token:  githubToken,
	}
}
