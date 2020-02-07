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

// CreateUser to create users
func (u *UserRepo) CreateUser(tx *pg.Tx, user *models.User) (*models.User, error) {
	_, err := tx.Model(user).Returning("*").Insert()
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

func (u *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	return u.GetUserByField("email", email)
}
