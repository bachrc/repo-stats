package domain

type RepositoryFilter interface {
	Filter(repos Repositories) Repositories
}

type LanguageFilter struct {
	Language string
}

func (filter LanguageFilter) Filter(repos Repositories) Repositories {
	var correspondingRepositories Repositories

	for i := range repos {
		if repos[i].containsLanguage(filter.Language) {
			correspondingRepositories = append(correspondingRepositories, repos[i])
		}
	}

	return correspondingRepositories
}
