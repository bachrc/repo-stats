package web

import (
	"bytes"
	"encoding/json"
	"github.com/Scalingo/go-utils/logger"
	"github.com/bachrc/profile-stats/internal/domain"
	githubhttpfetcher "github.com/bachrc/profile-stats/internal/github-http-fetcher"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockClient struct {
	DoFunc func(r *http.Request) (*http.Response, error)
}

func (client MockClient) Do(request *http.Request) (*http.Response, error) {
	return client.DoFunc(request)
}

func TestFetchRepositoriesNames(t *testing.T) {
	fetcher := githubhttpfetcher.GithubFetcher{
		Client: mockClient(defaultRoutes(), t),
	}
	statsDomain := domain.NewProfileStats(fetcher)
	handler := NewHandler(logger.Default(), 9876, statsDomain)

	t.Run("should return most recent repos names", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/repos", nil)
		response := httptest.NewRecorder()

		err := handler.RepositoriesHandler(response, request, nil)
		if err != nil {
			t.Errorf(err.Error())
		}

		var got Repositories
		_ = json.NewDecoder(response.Body).Decode(&got)

		wantedNumberOfRepos := 100
		gotNumberOfRepos := len(got)

		if gotNumberOfRepos != wantedNumberOfRepos {
			t.Fatalf("got %d repositories, wanted %d", gotNumberOfRepos, wantedNumberOfRepos)
		}

		t.Run("first repository must match", func(t *testing.T) {
			wantedFirstName := "cjmiyake/insoshi"
			gotFirstName := got[0].Name

			if gotFirstName != wantedFirstName {
				t.Errorf("got %q, want %q", gotFirstName, wantedFirstName)
			}
		})

	})
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
