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
	"testing"
)

func TestFetchRepositories(t *testing.T) {
	fetcher := githubhttpfetcher.GithubFetcher{
		Client: mockClient(defaultRoutes(), t),
	}
	statsDomain := domain.NewProfileStats(fetcher)
	handler := NewHandler(logger.Default(), 9876, statsDomain)

	t.Run("should return most recent repos names", func(t *testing.T) {
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

	})
}

// Test utils

type MockClient struct {
	DoFunc func(r *http.Request) (*http.Response, error)
}

func (client MockClient) Do(request *http.Request) (*http.Response, error) {
	return client.DoFunc(request)
}

func defaultRoutes() map[string]string {
	return map[string]string{
		githubhttpfetcher.GithubPublicReposUrl: "../../test/data/github-data/all-repos.json",
	}
}

func mockClient(routes map[string]string, t *testing.T) MockClient {
	return MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			requestedUrl := req.URL.String()

			textStubPath, exists := routes[requestedUrl]

			if !exists {
				t.Fatalf("Requested path %s not handled", requestedUrl)
			}

			file, err := ioutil.ReadFile(textStubPath)
			if err != nil {
				t.Fatalf("Can't open %s text file", textStubPath)
			}

			responseBody := ioutil.NopCloser(bytes.NewReader(file))
			return &http.Response{
				StatusCode: 200,
				Body:       responseBody,
			}, nil
		},
	}
}
