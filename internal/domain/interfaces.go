package domain

type PongContent struct {
	Pong string
}

type RepositoriesFetcher interface {
	Ping() PongContent
	GetAllRepositories() (Repositories, error)
}

type RepoStatsDomain struct {
	fetcher RepositoriesFetcher
}

func NewProfileStats(fetcher RepositoriesFetcher) RepoStatsDomain {
	return RepoStatsDomain{fetcher: fetcher}
}
