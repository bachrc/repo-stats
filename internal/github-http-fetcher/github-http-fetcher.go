package github_http_fetcher

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type GithubFetcher struct {
	Client HTTPClient
}

func New() GithubFetcher {
	return GithubFetcher{
		Client: &http.Client{},
	}
}
