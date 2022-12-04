package github_http_fetcher

import (
	"encoding/json"
	"fmt"
	"github.com/Scalingo/go-utils/logger"
	"github.com/bachrc/profile-stats/internal/domain"
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
	var githubPublicRepositories GithubPublicRepositories

	fetcher.FetchResource(GithubPublicReposUrl, &githubPublicRepositories)

	repositories := githubPublicRepositories.toDomain()

	for i := range repositories {
		fetcher.wgForLanguages.Add(1)
		i := i
		go func() {
			fetcher.fetchRepositoryLanguages(&repositories[i])
			fetcher.wgForLanguages.Done()
		}()
	}

	fetcher.wgForLanguages.Wait()

	return repositories, nil
}

func (fetcher *GithubFetcher) fetchRepositoryLanguages(repository *domain.Repository) {
	log.Infof("Fetching interfaces for %s", repository.Name)

	var githubLanguages GithubLanguagesForRepository

	_ = fetcher.FetchResource(fmt.Sprintf(GithubLanguagesForRepoTemplate, repository.Name), &githubLanguages)

	repository.Languages = githubLanguages.toDomain()
}

func (fetcher *GithubFetcher) FetchResource(url string, receiverObject interface{}) error {
	request, _ := http.NewRequest(http.MethodGet, url, nil)

	if fetcher.token != "" {
		request.Header.Set("Authorization", "Bearer "+fetcher.token)
	}

	response, err := fetcher.Client.Do(request)

	if err != nil {
		log.WithError(err).Error("Request execution failed")
		return err
	}

	if err := json.NewDecoder(response.Body).Decode(&receiverObject); err != nil {
		log.WithError(err).Error("Response decoding failed")
		return err
	}

	return nil
}
