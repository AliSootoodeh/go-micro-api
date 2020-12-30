package users

import (
	"github.com/AliSootoodeh/go-micro-api/datasources/mysql/users_db"
	"github.com/AliSootoodeh/go-micro-api/utils/date_utils"
	"github.com/AliSootoodeh/go-micro-api/utils/errors"
)

// Get ...
func (user *User) Get() *errors.RestErr {
	if err := users_db.UsersDB.Ping(); err != nil {
		panic(err)
	}
	return nil
}

//Save ...
func (user *User) Save() *errors.RestErr {

	user.DateCreated = date_utils.GetNowString()
	return nil
}
