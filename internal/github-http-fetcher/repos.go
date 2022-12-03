package github_http_fetcher

import (
	"encoding/json"
	"fmt"
	"github.com/Scalingo/go-utils/logger"
	"github.com/bachrc/profile-stats/internal/domain"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	githubApiUrl                   = "https://api.github.com"
	GithubPublicReposUrl           = githubApiUrl + "/repositories"
	GithubLanguagesForRepoTemplate = githubApiUrl + "/repos/%s/languages"
)

var (
	log = logger.Default()
)

func (fetcher *GithubFetcher) GetAllRepositories() (domain.Repositories, error) {
	request, _ := http.NewRequest(http.MethodGet, GithubPublicReposUrl, nil)
	if fetcher.token != "" {
		request.Header.Set("Authorization", "Bearer "+fetcher.token)
	}

	var githubPublicRepositories GithubPublicRepositories
	response, err := fetcher.Client.Do(request)
	if err != nil {
		return domain.Repositories{}, err
	}

	if err := json.NewDecoder(response.Body).Decode(&githubPublicRepositories); err != nil {
		logrus.Error(err)
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

func (fetcher *GithubFetcher) fetchRepositoryLanguages(repository *domain.Repository) {
	defer fetcher.wgForLanguages.Done()
	log.Infof("Fetching interfaces for %s", repository.Name)
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(GithubLanguagesForRepoTemplate, repository.Name), nil)

	if fetcher.token != "" {
		request.Header.Set("Authorization", "Bearer "+fetcher.token)
	}
	var githubLanguages GithubLanguagesForRepository
	response, err := fetcher.Client.Do(request)
	if err != nil {
		log.WithError(err).Error("Request execution failed")
		return
	}

	if err := json.NewDecoder(response.Body).Decode(&githubLanguages); err != nil {
		log.WithError(err).Error("Response decoding failed")
		return
	}

	repository.Languages = githubLanguages.toDomain()
}
