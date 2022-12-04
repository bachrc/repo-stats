package domain

type Repositories []Repository

type Repository struct {
	Id        uint
	Name      string
	Languages []string
	License   string
}

func (r Repository) containsLanguage(language string) bool {
	for i := range r.Languages {
		if r.Languages[i] == language {
			return true
		}
	}

	return false
}

func (domain RepoStatsDomain) GetAllRepositories(filters []RepositoryFilter, startingId int) (Repositories, error) {
	repositories, err := domain.fetcher.GetAllRepositories(startingId)

	if err != nil {
		return repositories, err
	}

	for _, filter := range filters {
		repositories = filter.Filter(repositories)
	}

	return repositories, nil
}
