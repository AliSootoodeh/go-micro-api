package services

import (
	"github.com/AliSootoodeh/go-micro-api/domain/users"
)

// CreateUser ...
func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
