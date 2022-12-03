package github_http_fetcher

import (
	"encoding/json"
	"fmt"
	"github.com/bachrc/profile-stats/internal/domain"
	"net/http"
)

const (
	githubApiUrl                   = "https://api.github.com"
	GithubPublicReposUrl           = githubApiUrl + "/repositories"
	GithubLanguagesForRepoTemplate = githubApiUrl + "/repos/%s/languages"
)

func (fetcher *GithubFetcher) GetAllRepositories() (domain.Repositories, error) {
	request, _ := http.NewRequest(http.MethodGet, GithubPublicReposUrl, nil)

	var githubPublicRepositories GithubPublicRepositories
	response, err := fetcher.Client.Do(request)
	if err != nil {
		return domain.Repositories{}, err
	}

	if err := json.NewDecoder(response.Body).Decode(&githubPublicRepositories); err != nil {
		return domain.Repositories{}, err
	}

	repositories := githubPublicRepositories.toDomain()

	for i := range repositories {
		fetcher.wgForLanguages.Add(1)
		go fetcher.fetchRepositoryLanguages(&repositories[i])
	}

	fetcher.wgForLanguages.Wait()

	return repositories, nil
}

func (fetcher *GithubFetcher) fetchRepositoryLanguages(repository *domain.Repository) error {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(GithubLanguagesForRepoTemplate, repository.Name), nil)

	var githubLanguages GithubLanguagesForRepository
	response, err := fetcher.Client.Do(request)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(response.Body).Decode(&githubLanguages); err != nil {
		return err
	}

	repository.Languages = githubLanguages.toDomain()

	fetcher.wgForLanguages.Done()

	return nil
}
