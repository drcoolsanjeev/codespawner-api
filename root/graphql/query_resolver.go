package graphql

import (
	"context"

	"github.com/codespawner-api/root/models"
)

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context) ([]*models.User, error) {
	return r.UsersRepo.GetUsers()
}
