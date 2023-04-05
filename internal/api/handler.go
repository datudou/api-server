package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/retail-ai-test/internal/api/middleware"
	"github.com/retail-ai-test/internal/pkg/service"
)

// Handler struct holds required services for handler to function
type Handler struct {
	Services *service.Services
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R        *gin.Engine
	Services *service.Services
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	h := &Handler{
		Services: c.Services,
	}

	c.R.Use(gin.Recovery())
	c.R.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Recipe API",
		})
	})

	// recipe routes
	{
		c.R.GET("/recipes", h.getRecipes)
		c.R.GET("/recipes/:id", h.getRecipeByID)
		c.R.POST("/recipes", h.createRecipe)
		c.R.DELETE("/recipes/:id", h.deleteRecipeByID)
		c.R.PATCH("/recipes/:id", h.updateRecipeByID)
	}

	// user routes
	c.R.POST("/signup", h.createUser)
	auth := c.R.Use(middleware.Auth())
	{
		auth.GET("/users/:user_id", h.getUserByID)
		auth.PATCH("/users/:user_id", h.updateUserByID)
		auth.DELETE("/users/:user_id", h.deleteUserByID)
	}
}
