package github

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepositoryRepositoryGet(t *testing.T) {
	tests := []struct {
		name     string
		owner    string
		repoName string
		repo     Repository
	}{
		{
			name:     "hashicorp/terraform",
			owner:    "hashicorp",
			repoName: "terraform",
			repo:     newRepository("hashicorp", "terraform"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			r := NewRepositoryRepository(githubTokenForTest)
			repo, err := r.get(ctx, tt.owner, tt.name)
			require.NoError(err)
			require.Equal(tt.repo, repo)
		})
	}
}

func TestRepositoryRepositorySearch(t *testing.T) {
	tests := []struct {
		name  string
		query string
		repo  Repository
	}{
		{
			name:  "hashicorp/terraform",
			query: "terraform",
			repo:  newRepository("hashicorp", "terraform"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			r := NewRepositoryRepository(githubTokenForTest)
			repo, err := r.search(ctx, tt.query)
			require.NoError(err)
			require.Equal(tt.repo, repo)
		})
	}
}
