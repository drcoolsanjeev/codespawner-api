package graphql

import (
	"context"
	"errors"
	"log"

	"github.com/codespawner-api/root/models"
)

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	user, err := r.UsersRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, err
	}
	token, err := user.GenToken()
	if err != nil {
		return nil, err
	}
	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (r *mutationResolver) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	_, err := r.UsersRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, errors.New("email already in use")
	}

	user := &models.User{
		Email: input.Email,
		Name:  input.Name,
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		log.Fatal("error while hashing password: %v", err)
		return nil, errors.New("something went wrong")
	}

	tx, err := r.UsersRepo.DB.Begin()
	if err != nil {
		log.Fatal("error creating a transaction: %v", err)
		return nil, errors.New("something went wrong")
	}

	defer tx.Rollback()
	if _, err := r.UsersRepo.CreateUser(tx, user); err != nil {
		log.Fatal("error creating a transaction: %v", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Fatal("error creating a transaction: %v", err)
		return nil, err
	}

	token, err := user.GenToken()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}
