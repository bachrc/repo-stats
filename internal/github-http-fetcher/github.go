package github_http_fetcher

import "github.com/bachrc/profile-stats/internal/domain"

type GithubPublicRepositories []GithubPublicRepository

type GithubPublicRepository struct {
	Id       uint   `json:"id"`
	FullName string `json:"full_name"`
}

type GithubLanguagesForRepository map[string]int

func (repositories GithubPublicRepositories) toDomain() domain.Repositories {
	var domainRepositories domain.Repositories
	for _, repository := range repositories {
		domainRepositories = append(domainRepositories, domain.Repository{
			Id:   repository.Id,
			Name: repository.FullName,
		})
	}

	return domainRepositories
}

func (githubLanguages GithubLanguagesForRepository) toDomain() []string {
	languages := make([]string, len(githubLanguages))
	i := 0
	for k := range githubLanguages {
		languages[i] = k
		i++
	}

	return languages
}
