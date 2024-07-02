package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/majorchork/book-keeper/internal/api"
	"github.com/majorchork/book-keeper/internal/ports"
)

// SetupRouter is where router endpoints are called
func SetupRouter(handler *api.HTTPHandler, repository ports.Repository) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r := router.Group("/computers")
	{
		r.GET("/health", handler.Health)
		r.POST("/create", handler.CreateComputer)
		r.GET("/viewAll", handler.ViewAllComputers)
		r.GET("/employee/:abbr", handler.ViewEmployeeComputers)
		r.GET("/viewInfo/:id", handler.ViewComputerInfo)
		r.DELETE("/delete/:id", handler.DeleteComputer)
		r.PUT("/assign", handler.AssignComputer)
	}

	return router
}
