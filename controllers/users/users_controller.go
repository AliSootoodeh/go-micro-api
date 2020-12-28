package users

import (
	"fmt"
	"net/http"

	"github.com/AliSootoodeh/go-micro-api/domain/users"
	"github.com/AliSootoodeh/go-micro-api/services"
	"github.com/gin-gonic/gin"
)

// CreateUser TODO
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		fmt.Println(saveErr)
		//TODO
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetUser TODO
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "TODO!!")
}

// SearchUser TODO
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "TODO!!")
}
