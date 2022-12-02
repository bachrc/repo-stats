package web

import (
	"encoding/json"
	"github.com/Scalingo/go-utils/logger"
	githubhttpfetcher "github.com/bachrc/profile-stats/internal/github-http-fetcher"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PongResponse struct {
	Response string `json:"response"`
}

func TestPong(t *testing.T) {
	handler := NewHandler(logger.Default(), 9876, githubhttpfetcher.GithubFetcher{})

	t.Run("we shall PONG", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/pong", nil)
		response := httptest.NewRecorder()

		err := handler.PongHandler(response, request, nil)
		if err != nil {
			t.Errorf(err.Error())
		}

		var got PongResponse
		_ = json.NewDecoder(response.Body).Decode(&got)
		want := "pongggg"

		if got.Response != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
