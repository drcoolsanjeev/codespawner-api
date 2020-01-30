package graphql

import (
	"context"

	"github.com/codespawner-api/root/models"
)

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*models.User, error) {

	user := &models.User{
		Name:  input.Name,
		Email: input.Email,
	}

	return r.UsersRepo.CreateUser(user)
}
