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

type PublicRepositories struct {
	Repositories []PublicRepository
}

func (repositories PublicRepositories) toDomain() domain.Repositories {
	var domainRepositories []domain.Repository
	for _, repository := range repositories.Repositories {
		domainRepositories = append(domainRepositories, repository.toDomain())
	}

	return domain.Repositories{
		Repositories: domainRepositories,
	}
}

type PublicRepository struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (repository PublicRepository) toDomain() domain.Repository {
	return domain.Repository{
		Id:   repository.Id,
		Name: repository.Name,
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
