package web

import (
	"encoding/json"
	"github.com/Scalingo/go-utils/logger"
	"net/http"
)

type Repositories struct {
	Repositories []Repository
}

type Repository struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (handler *ProfileStatsWebHandler) RepositoriesHandler(w http.ResponseWriter, r *http.Request, params map[string]string) error {
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
