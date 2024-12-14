package github

import (
	"github.com/google/go-github/v62/github"
	"golang.org/x/net/context"
)

// Repository represents a GitHub repository.
type Repository struct {
	owner string
	name  string
}

// newRepository returns a new [Repository] object.
func newRepository(owner string, name string) Repository {
	return Repository{
		owner: owner,
		name:  name,
	}
}

// AssetRepository is a repository for [Repository].
type RepositoryRepository struct {
	client *github.Client
}

// NewAssetRepository returns a new [AssetRepository] object.
func NewRepositoryRepository(token string) *RepositoryRepository {
	return &RepositoryRepository{
		client: newGitHubClient(token),
	}
}

// get a GitHub repository.
func (r *RepositoryRepository) get(ctx context.Context, owner string, name string) (Repository, error) {
	repo, _, err := r.client.Repositories.Get(ctx, owner, name)
	if err != nil {
		return Repository{}, err
	}
	return newRepository(repo.GetOwner().GetLogin(), repo.GetName()), nil
}

// search a GitHub repository.
func (r *RepositoryRepository) search(ctx context.Context, query string) (Repository, error) {
	result, _, err := r.client.Search.Repositories(ctx, query, &github.SearchOptions{
		ListOptions: github.ListOptions{
			PerPage: 1,
		},
	})
	if err != nil {
		return Repository{}, err
	}
	if len(result.Repositories) == 0 {
		return Repository{}, ErrNotFound
	}
	repo := result.Repositories[0]
	return newRepository(repo.GetOwner().GetLogin(), repo.GetName()), nil
}
