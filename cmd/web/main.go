package main

import (
	"github.com/bachrc/profile-stats/internal/domain"
	githubhttpfetcher "github.com/bachrc/profile-stats/internal/github-http-fetcher"
	"github.com/bachrc/profile-stats/internal/web"
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

	handler := web.NewHandler(log, cfg.Port, domain.NewProfileStats(githubhttpfetcher.New()))

	handler.Serve()
}
