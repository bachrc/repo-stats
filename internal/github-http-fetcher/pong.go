package github_http_fetcher

import (
	"github.com/bachrc/profile-stats/internal/domain"
)

func (fetcher *GithubFetcher) Ping() domain.PongContent {
	return domain.PongContent{
		Pong: "pongggg",
	}
}
