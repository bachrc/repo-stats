package web

import (
	"encoding/json"
	"fmt"
	"github.com/Scalingo/go-utils/logger"
	"github.com/bachrc/profile-stats/internal/domain"
	"github.com/sirupsen/logrus"

	//	"github.com/Scalingo/go-handlers"
	"github.com/Scalingo/go-handlers"
	"net/http"
)

type ProfileStatsWebHandler struct {
	router  handlers.Router
	port    int
	log     logrus.FieldLogger
	fetcher domain.ProfileFetcher
}

func NewHandler(log logrus.FieldLogger, port int, fetcher domain.ProfileFetcher) ProfileStatsWebHandler {
	router := *handlers.NewRouter(log)

	// Initialize web server and configure the following routes:
	// GET /repos
	// GET /stats

	handler := ProfileStatsWebHandler{
		router:  router,
		port:    port,
		log:     log,
		fetcher: fetcher,
	}

	router.HandleFunc("/ping", handler.PongHandler)

	return handler
}

func (handler *ProfileStatsWebHandler) PongHandler(w http.ResponseWriter, r *http.Request, params map[string]string) error {
	log := logger.Get(r.Context())
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	pongResponse := handler.fetcher.Pong().Pong

	err := json.NewEncoder(w).Encode(map[string]string{"response": pongResponse})
	if err != nil {
		log.WithError(err).Error("Fail to encode JSON")
	}
	return nil
}

func (handler *ProfileStatsWebHandler) Serve() {
	handler.log.WithField("port", handler.port).Info("Listening...")
	err := http.ListenAndServe(fmt.Sprintf(":%d", handler.port), handler.router)
	if err != nil {
		logrus.Errorln("Erreur lors du démarrage de l'application")
	}
}
