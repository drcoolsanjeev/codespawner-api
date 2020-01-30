//go:generate go run github.com/99designs/gqlgen
package graphql

import (
	"github.com/codespawner-api/root/postgres"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	UsersRepo postgres.UserRepo
}
