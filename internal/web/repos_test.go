package web

import (
	"bytes"
	"encoding/json"
	"github.com/Scalingo/go-utils/logger"
	"github.com/bachrc/profile-stats/internal/domain"
	githubhttpfetcher "github.com/bachrc/profile-stats/internal/github-http-fetcher"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestFetchRepositories(t *testing.T) {
	fetcher := &githubhttpfetcher.GithubFetcher{
		Client: mockClient(t),
	}
	statsDomain := domain.NewProfileStats(fetcher)
	handler := NewHandler(logger.Default(), 9876, statsDomain)

	t.Run("should return repositories without filters", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/repos", nil)
		response := httptest.NewRecorder()

		err := handler.RepositoriesHandler(response, request, nil)
		assert.NoError(t, err)

		var receivedRepositories Repositories
		assert.NoError(t, json.NewDecoder(response.Body).Decode(&receivedRepositories))

		t.Run("all repositories should have been retrieved", func(t *testing.T) {
			assert.Equal(t, 20, len(receivedRepositories), "eh oui")
		})

		t.Run("first repository must match", func(t *testing.T) {
			firstRepository := receivedRepositories[0]

			assert.Equal(t, "cjmiyake/insoshi", firstRepository.Name)
			assert.Equal(t, uint(10005), firstRepository.Id)
		})

		t.Run("last repository must match", func(t *testing.T) {
			lastRepository := receivedRepositories[19]

			assert.Equal(t, "CodeMonkeySteve/libfinagle", lastRepository.Name)
			assert.Equal(t, uint(10085), lastRepository.Id)
		})

		t.Run("should synthetize used languages", func(t *testing.T) {
			firstRepository := receivedRepositories[0]

			assert.Contains(t, firstRepository.Languages, "Ruby")
			assert.Contains(t, firstRepository.Languages, "JavaScript")
		})

		t.Run("empty repositories return empty languages", func(t *testing.T) {
			emptyRepository := receivedRepositories[11]
			assert.Equal(t, []string{}, emptyRepository.Languages)
		})

		t.Run("should return used license", func(t *testing.T) {
			repositoryWithoutLicense := receivedRepositories[16]
			assert.Equal(t, "none", repositoryWithoutLicense.License)
		})
	})

	t.Run("should run filters on repositories", func(t *testing.T) {
		t.Run("should filter repositories to include language", func(t *testing.T) {
			var receivedRepositories Repositories

			fetchResource(t, &handler, "/repos?language="+url.QueryEscape("C++"), &receivedRepositories)

			assert.Len(t, receivedRepositories, 1)
			receivedRepository := receivedRepositories[0]

			assert.Equal(t, "CodeMonkeySteve/libfinagle", receivedRepository.Name)
			assert.Contains(t, receivedRepository.Languages, "C++")
		})

		t.Run("should filter repositories matching license", func(t *testing.T) {
			var receivedRepositories Repositories

			fetchResource(t, &handler, "/repos?license="+url.QueryEscape("mit"), &receivedRepositories)

			assert.Len(t, receivedRepositories, 5)
			receivedRepository := receivedRepositories[0]

			assert.Equal(t, "rayvinly/badger", receivedRepository.Name)
			assert.Equal(t, "mit", receivedRepository.License)
		})

		t.Run("should apply multiple filters at the same time", func(t *testing.T) {
			var receivedRepositories Repositories

			fetchResource(t, &handler, "/repos?license="+url.QueryEscape("mit")+"&language="+url.QueryEscape("JavaScript"), &receivedRepositories)

			assert.Len(t, receivedRepositories, 1)
			receivedRepository := receivedRepositories[0]

			assert.Equal(t, "bytheway/merb-core", receivedRepository.Name)
			assert.Equal(t, "mit", receivedRepository.License)
		})
	})
}

// Test utils

func fetchResource(t *testing.T, handler *ProfileStatsWebHandler, url string, receiverObject interface{}) {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	response := httptest.NewRecorder()

	err := handler.RepositoriesHandler(response, request, nil)
	assert.NoError(t, err)

	if err := json.NewDecoder(response.Body).Decode(&receiverObject); err != nil {
		t.Error(err)
	}
}

type MockClient struct {
	DoFunc func(r *http.Request) (*http.Response, error)
}

func (client MockClient) Do(request *http.Request) (*http.Response, error) {
	return client.DoFunc(request)
}

func mockClient(t *testing.T) MockClient {
	return MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			requestedGithubPath := req.URL.Path

			stubPath := "../../test/data/github-data" + requestedGithubPath + ".json"
			file, err := ioutil.ReadFile(stubPath)
			if err != nil {
				t.Fatalf("Can't open %s text file", stubPath)
			}

			responseBody := ioutil.NopCloser(bytes.NewReader(file))
			return &http.Response{
				StatusCode: 200,
				Body:       responseBody,
			}, nil
		},
	}
}
