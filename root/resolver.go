//go:generate go run github.com/99designs/gqlgen
package root

import (
	"context"

	"github.com/codespawner-api/root/models"
	"github.com/codespawner-api/root/postgres"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	UsersRepo postgres.UserRepo
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*models.User, error) {

	user := &models.User{
		Name:  input.Name,
		Email: input.Email,
	}

	return r.UsersRepo.CreateUser(user)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context) ([]*models.User, error) {
	return r.UsersRepo.GetUsers()
}
