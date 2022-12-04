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
	GithubPublicReposUrlTemplate   = githubApiUrl + "/repositories?since=%d"
	GithubLanguagesForRepoTemplate = githubApiUrl + "/repos/%s/languages"
	GithubLicenseForRepoTemplate   = githubApiUrl + "/repos/%s/license"
)

var (
	log = logger.Default()
)

func (fetcher *GithubFetcher) GetAllRepositories(startingId int) (domain.Repositories, error) {
	var githubPublicRepositories GithubPublicRepositories

	if err := fetcher.fetchResource(fmt.Sprintf(GithubPublicReposUrlTemplate, startingId), &githubPublicRepositories); err != nil {
		return domain.Repositories{}, err
	}

	repositories := githubPublicRepositories.toDomain()

	for i := range repositories {
		fetcher.wg.Add(2)
		i := i
		go func() {
			fetcher.fetchRepositoryLanguages(&repositories[i])
			fetcher.wg.Done()
		}()
		go func() {
			fetcher.fetchRepositoryLicense(&repositories[i])
			fetcher.wg.Done()
		}()
	}

	fetcher.wg.Wait()

	return repositories, nil
}

func (fetcher *GithubFetcher) fetchRepositoryLanguages(repository *domain.Repository) {
	log.Infof("Fetching interfaces for %s", repository.Name)

	var githubLanguages GithubLanguagesForRepository

	_ = fetcher.fetchResource(fmt.Sprintf(GithubLanguagesForRepoTemplate, repository.Name), &githubLanguages)

	repository.Languages = githubLanguages.toDomain()
}

func (fetcher *GithubFetcher) fetchRepositoryLicense(repository *domain.Repository) {
	log.Infof("Fetching interfaces for %s", repository.Name)

	var githubLicense GithubLicenseForRepository

	_ = fetcher.fetchResource(fmt.Sprintf(GithubLicenseForRepoTemplate, repository.Name), &githubLicense)

	repository.License = githubLicense.toDomain()
}

func (fetcher *GithubFetcher) fetchResource(url string, receiverObject interface{}) error {
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
