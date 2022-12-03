package web

import (
	"encoding/json"
	"github.com/Scalingo/go-utils/logger"
	"github.com/bachrc/profile-stats/internal/domain"
	"net/http"
)

type Repositories []Repository

type Repository struct {
	Id        uint     `json:"id"`
	Name      string   `json:"name"`
	Languages []string `json:"languages"`
}

func fromDomainRepositories(domainRepositories domain.Repositories) Repositories {
	var repositories Repositories
	for _, domainRepository := range domainRepositories {
		repositories = append(repositories, fromDomainRepository(domainRepository))
	}

	return repositories
}

func fromDomainRepository(repository domain.Repository) Repository {
	return Repository{
		Id:        repository.Id,
		Name:      repository.Name,
		Languages: repository.Languages,
	}
}

func (handler *ProfileStatsWebHandler) RepositoriesHandler(w http.ResponseWriter, r *http.Request, params map[string]string) error {
	log := logger.Get(r.Context())
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	domainRepositories, err := handler.domain.GetAllRepositories()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Error("Error while fetching repositories")
		return err
	}

	repositories := fromDomainRepositories(domainRepositories)

	err = json.NewEncoder(w).Encode(repositories)
	if err != nil {
		log.WithError(err).Error("Fail to encode JSON")
		return err
	}

	return nil
}
