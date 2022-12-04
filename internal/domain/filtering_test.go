package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	repositories = Repositories{
		Repository{
			Id:        1,
			Name:      "pouetpouet/labalayette",
			Languages: []string{"Rust", "Go"},
			License:   "none",
		}, Repository{
			Id:        2,
			Name:      "johnny/lefeu",
			Languages: []string{"Rust", "Java"},
			License:   "mit",
		}, Repository{
			Id:        3,
			Name:      "la/pin",
			Languages: []string{"Kotlin"},
			License:   "mit",
		},
	}
)

func TestFiltering(t *testing.T) {
	t.Run("Language filtering", func(t *testing.T) {
		t.Run("Language filter should filter a Language", func(t *testing.T) {
			filter := LanguageFilter{
				Language: "Java",
			}

			filteredRepos := filter.Filter(repositories)

			assert.Len(t, filteredRepos, 1)
			assert.Equal(t, "johnny/lefeu", filteredRepos[0].Name)
		})

		t.Run("Language filter should return an empty slice when none match", func(t *testing.T) {
			filter := LanguageFilter{
				Language: "Pouetpouet",
			}

			filteredRepos := filter.Filter(repositories)

			assert.Empty(t, filteredRepos)
		})
	})

	t.Run("License filtering", func(t *testing.T) {
		t.Run("repositories should be filterable based on license", func(t *testing.T) {
			filter := LicenseFilter{
				License: "mit",
			}

			filteredRepos := filter.Filter(repositories)

			assert.Len(t, filteredRepos, 2)
			assert.Equal(t, "johnny/lefeu", filteredRepos[0].Name)
			assert.Equal(t, "la/pin", filteredRepos[1].Name)
		})
	})
}
