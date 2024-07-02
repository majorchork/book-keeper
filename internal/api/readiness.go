package api

import (
	"github.com/gin-gonic/gin"
	"github.com/majorchork/book-keeper/internal/util"
)

// Health is to check if server is up
func (u *HTTPHandler) Health(c *gin.Context) {
	data := "server is up and running"

	// healthcheck
	util.Response(c, "healthy", 200, data, nil)
}
