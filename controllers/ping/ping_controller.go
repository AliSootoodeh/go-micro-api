package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping shows that our service is up
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
