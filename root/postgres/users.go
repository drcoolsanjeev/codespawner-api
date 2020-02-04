package postgres

import (
	"fmt"

	"github.com/codespawner-api/root/models"
	"github.com/go-pg/pg/v9"
)

// UserRepo struct to hold User detail
type UserRepo struct {
	DB *pg.DB
}

// GetUsers All Users
func (u *UserRepo) GetUsers() ([]*models.User, error) {
	var users []*models.User
	err := u.DB.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser to create users
func (u *UserRepo) CreateUser(user *models.User) (*models.User, error) {
	_, err := u.DB.Model(user).Returning("*").Insert()
	return user, err
}

func (u *UserRepo) GetUserByField(field, value string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &user, err
}

func (u *UserRepo) GetUserByID(id string) (*models.User, error) {
	return u.GetUserByField("id", id)
}
