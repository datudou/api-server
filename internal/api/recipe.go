package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/retail-ai-test/internal/model/apperrors"
)

func (h *Handler) getRecipeByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	gameLog, err := h.Services.RecipeService.FindByID(ctx, uint(id))
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gameLog)
}

func (h *Handler) getRecipes(c *gin.Context) {
	ctx := c.Request.Context()
	games, err := h.Services.RecipeService.FindAll(ctx)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, games)
}
