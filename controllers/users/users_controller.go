package users

import (
	"net/http"
	"strconv"

	"github.com/AliSootoodeh/go-micro-api/domain/users"
	"github.com/AliSootoodeh/go-micro-api/services"
	"github.com/AliSootoodeh/go-micro-api/utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateUser TODO
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetUser ...
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("userId should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// SearchUser TODO
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "TODO!!")
}
