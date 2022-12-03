package web

import (
	"fmt"
	"github.com/bachrc/profile-stats/internal/domain"
	"github.com/sirupsen/logrus"

	//	"github.com/Scalingo/go-handlers"
	"github.com/Scalingo/go-handlers"
	"net/http"
)

type ProfileStatsWebHandler struct {
	router handlers.Router
	port   int
	log    logrus.FieldLogger
	domain domain.RepoStatsDomain
}

func NewHandler(log logrus.FieldLogger, port int, fetcher domain.RepoStatsDomain) ProfileStatsWebHandler {
	router := *handlers.NewRouter(log)

	handler := ProfileStatsWebHandler{
		router: router,
		port:   port,
		log:    log,
		domain: fetcher,
	}

	router.HandleFunc("/ping", handler.PongHandler)
	router.HandleFunc("/repos", handler.RepositoriesHandler)

	return handler
}

func (handler *ProfileStatsWebHandler) Serve() {
	handler.log.WithField("port", handler.port).Info("Listening...")
	err := http.ListenAndServe(fmt.Sprintf(":%d", handler.port), handler.router)
	if err != nil {
		logrus.Errorln("Erreur lors du démarrage de l'application")
	}
}
