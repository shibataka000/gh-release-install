package github

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseRepository(t *testing.T) {
	tests := []struct {
		name string
		s    string
		repo Repository
	}{
		{
			name: "hashicorp/terraform",
			s:    "hashicorp/terraform",
			repo: newRepository("hashicorp", "terraform"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			repo, err := parseRepository(tt.s)
			require.NoError(err)
			require.Equal(tt.repo, repo)
		})
	}
}
