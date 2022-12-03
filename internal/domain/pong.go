package domain

func (domain RepoStatsDomain) Ping() PongContent {
	return domain.fetcher.Ping()
}
