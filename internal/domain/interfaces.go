package domain

type PongContent struct {
	Pong string
}

type RepositoriesFetcher interface {
	Ping() PongContent
	GetAllRepositories(int) (Repositories, error)
}

type RepoStatsDomain struct {
	fetcher RepositoriesFetcher
}

func NewProfileStats(fetcher RepositoriesFetcher) RepoStatsDomain {
	return RepoStatsDomain{fetcher: fetcher}
}
