package domain

type Repositories []Repository

type Repository struct {
	Id        uint
	Name      string
	Languages []string
	License   string
}

func (domain RepoStatsDomain) GetAllRepositories() (Repositories, error) {
	return domain.fetcher.GetAllRepositories()
}
