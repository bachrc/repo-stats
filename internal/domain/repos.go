package domain

type Repositories []Repository

type Repository struct {
	Id   uint
	Name string
}

func (domain RepoStatsDomain) GetAllRepositories() (Repositories, error) {
	return domain.fetcher.GetAllRepositories()
}
