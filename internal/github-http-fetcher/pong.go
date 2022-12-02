package github_http_fetcher

import (
	"github.com/bachrc/profile-stats/internal/domain"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type GithubFetcher struct {
	Client HTTPClient
}

func (fetcher GithubFetcher) Pong() domain.PongContent {
	return domain.PongContent{
		Pong: "pongggg",
	}
}
