package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// StartApplication starts our web server on port 8080
func StartApplication() {
	mapUrls()
	router.Run(":8080")

}
