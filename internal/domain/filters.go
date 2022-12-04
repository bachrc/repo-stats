package domain

type RepositoryFilter interface {
	Filter(repos Repositories) Repositories
}

type LanguageFilter struct {
	Language string
}

type LicenseFilter struct {
	License string
}

func (filter LicenseFilter) Filter(repos Repositories) Repositories {
	var correspondingRepositories Repositories

	for i := range repos {
		if repos[i].License == filter.License {
			correspondingRepositories = append(correspondingRepositories, repos[i])
		}
	}

	return correspondingRepositories
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
