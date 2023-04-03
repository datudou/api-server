package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/retail-ai-test/internal/middleware"
	"github.com/retail-ai-test/internal/model/apperrors"
	"github.com/retail-ai-test/internal/service"
)

// Handler struct holds required services for handler to function
type Handler struct {
	Services *service.Services
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R               *gin.Engine
	Services        *service.Services
	TimeoutDuration time.Duration
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	h := &Handler{
		Services: c.Services,
	}

	c.R.Use(gin.Recovery())
	c.R.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))

	g := c.R.Group("/api/v1")
	{
		g.GET("/recipes", h.getRecipes)
		g.GET("/recipes/:id", h.getRecipeByID)
		g.POST("/recipes", h.createRecipe)
		g.DELETE("/recipes/:id", h.deleteRecipeByID)
		g.PATCH("/recipes/:id", h.updateRecipeByID)
	}
}
