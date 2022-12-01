package main

import (
	"github.com/Scalingo/sclng-backend-test-v1/internal/web"
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

	handler := web.NewHandler(log, cfg.Port)

	handler.Serve()
}
