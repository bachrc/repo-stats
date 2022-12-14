package main

import (
	"github.com/bachrc/repo-stats/internal/domain"
	githubhttpfetcher "github.com/bachrc/repo-stats/internal/github-http-fetcher"
	"github.com/bachrc/repo-stats/internal/web"
	"os"

	"github.com/Scalingo/go-utils/logger"
)

func main() {
	log := logger.Default()
	log.Info("Initializing app")
	cfg, err := NewConfig()
	if err != nil {
		log.WithError(err).Error("Fail to initialize configuration")
		os.Exit(-1)
	}

	fetcher := githubhttpfetcher.New(cfg.GithubAccessToken)
	handler := web.NewHandler(log, cfg.Port, domain.NewProfileStats(fetcher))

	handler.Serve()
}
