package github_http_fetcher

type PublicRepositories struct {
	Repositories []PublicRepository
}

type PublicRepository struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
