package services

import (
	"github.com/AliSootoodeh/go-micro-api/domain/users"
	"github.com/AliSootoodeh/go-micro-api/utils/errors"
)

// CreateUser ...
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
