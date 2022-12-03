package github_http_fetcher

import (
	"encoding/json"
	"github.com/bachrc/profile-stats/internal/domain"
	"net/http"
)

const (
	githubApiUrl         = "https://api.github.com"
	GithubPublicReposUrl = githubApiUrl + "/repositories"
)

type PublicRepositories []struct {
	Id       uint   `json:"id"`
	FullName string `json:"full_name"`
}

func (repositories PublicRepositories) toDomain() domain.Repositories {
	var domainRepositories []domain.Repository
	for _, repository := range repositories {
		domainRepositories = append(domainRepositories, domain.Repository{
			Id:   repository.Id,
			Name: repository.FullName,
		})
	}

	return domain.Repositories{
		Repositories: domainRepositories,
	}
}

func (fetcher GithubFetcher) GetAllRepositories() (domain.Repositories, error) {
	request, _ := http.NewRequest(http.MethodGet, GithubPublicReposUrl, nil)

	var repositories PublicRepositories
	response, err := fetcher.Client.Do(request)
	if err != nil {
		return domain.Repositories{}, err
	}

	if err := json.NewDecoder(response.Body).Decode(&repositories); err != nil {
		return domain.Repositories{}, err
	}

	return repositories.toDomain(), nil
}
