package github

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRepositoryFromFullName(t *testing.T) {
	tests := []struct {
		name     string
		fullName string
		repo     Repository
	}{
		{
			name:     "hashicorp/terraform",
			fullName: "hashicorp/terraform",
			repo:     NewRepository("hashicorp", "terraform"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			repo, err := NewRepositoryFromFullName(tt.fullName)
			require.NoError(err)
			require.Equal(tt.repo, repo)
		})
	}
}
