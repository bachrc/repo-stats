package web

import (
	"encoding/json"
	"github.com/Scalingo/go-utils/logger"
	githubhttpfetcher "github.com/bachrc/profile-stats/internal/github-http-fetcher"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchRepositoriesNames(t *testing.T) {
	handler := NewHandler(logger.Default(), 9876, githubhttpfetcher.GithubFetcher{})

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
		gotNumberOfRepos := len(got.Repositories)

		if gotNumberOfRepos != wantedNumberOfRepos {
			t.Fatalf("got %d repositories, wanted %d", gotNumberOfRepos, wantedNumberOfRepos)
		}

		t.Run("first repository must match", func(t *testing.T) {
			wantedFirstName := "cjmiyake/insoshi"
			gotFirstName := got.Repositories[0].Name

			if gotFirstName != wantedFirstName {
				t.Errorf("got %q, want %q", got, wantedFirstName)
			}
		})

	})
}
