package postgres

import (
	"github.com/codespawner-api/root/models"
	"github.com/go-pg/pg/v9"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetUsers() ([]*models.User, error) {
	var users []*models.User
	err := u.DB.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepo) CreateUser(user *models.User) (*models.User, error) {
	_, err := u.DB.Model(user).Returning("*").Insert()
	return user, err
}
