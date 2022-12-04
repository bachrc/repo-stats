package web

import (
	"encoding/json"
	"github.com/Scalingo/go-utils/logger"
	"github.com/bachrc/profile-stats/internal/domain"
	"net/http"
	"net/url"
)

type Repositories []Repository

type Repository struct {
	Id        uint     `json:"id"`
	Name      string   `json:"name"`
	Languages []string `json:"languages"`
	License   string   `json:"license"`
}

var (
	log = logger.Default()
)

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
		License:   repository.License,
	}
}

func (handler *ProfileStatsWebHandler) RepositoriesHandler(w http.ResponseWriter, r *http.Request, params map[string]string) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	filters := parseRequestedFilters(r.URL.Query())

	domainRepositories, err := handler.domain.GetAllRepositories(filters)

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

func parseRequestedFilters(params url.Values) (filters []domain.RepositoryFilter) {
	if params.Has("language") {
		language := params.Get("language")
		filters = append(filters, domain.LanguageFilter{Language: language})
	}

	if params.Has("license") {
		license := params.Get("license")
		filters = append(filters, domain.LicenseFilter{License: license})
	}

	return filters
}
