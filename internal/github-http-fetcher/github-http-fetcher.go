package github_http_fetcher

import (
	"net/http"
	"sync"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type GithubFetcher struct {
	Client         HTTPClient
	wgForLanguages sync.WaitGroup
}

func New() *GithubFetcher {
	return &GithubFetcher{
		Client: &http.Client{},
	}
}
