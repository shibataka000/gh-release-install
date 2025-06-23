package github_test

import (
	"testing"

	"github.com/shibataka000/gh-release-install/github"
	"github.com/stretchr/testify/require"
)

func TestParseRepository(t *testing.T) {
	tests := []struct {
		name string
		s    string
		repo github.Repository
	}{
		{
			name: "hashicorp/terraform",
			s:    "hashicorp/terraform",
			repo: github.Repository{Host: "github.com", Owner: "hashicorp", Name: "terraform"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			repo, err := github.ParseRepository(tt.s)
			require.NoError(err)
			require.Equal(tt.repo, repo)
		})
	}
}
