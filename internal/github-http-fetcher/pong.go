package github_http_fetcher

import (
	"github.com/bachrc/repo-stats/internal/domain"
)

func (fetcher *GithubFetcher) Ping() domain.PongContent {
	return domain.PongContent{
		Pong: "pongggg",
	}
}
