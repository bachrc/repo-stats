package github_http_fetcher

import "github.com/bachrc/profile-stats/internal/domain"

type GithubFetcher struct{}

func (fetcher GithubFetcher) Pong() domain.PongContent {
	return domain.PongContent{
		Pong: "pongggg",
	}
}
