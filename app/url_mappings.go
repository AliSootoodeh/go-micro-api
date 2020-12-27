package app

import (
	"github.com/AliSootoodeh/go-micro-api/controllers/ping"
	"github.com/AliSootoodeh/go-micro-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/search", users.SearchUser)
	router.GET("/users/user_id", users.GetUser)

	router.POST("/users", users.CreateUser)
}
