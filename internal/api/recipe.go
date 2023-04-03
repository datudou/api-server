package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/retail-ai-test/internal/model"
	"github.com/retail-ai-test/internal/model/apperrors"
	"go.uber.org/zap"
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
	recipe, err := h.Services.RecipeService.FindByID(ctx, uint(id))
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Recipe details by id",
		"recipe":  []*model.Recipe{recipe},
	})
}

func (h *Handler) getRecipes(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := h.Services.RecipeService.FindAll(ctx)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"recipes": result,
	})
}

func (h *Handler) createRecipe(c *gin.Context) {

	type createRecipeParam struct {
		Title       string `json:"title" binding:"required"`
		MakingTime  string `json:"making_time" binding:"required"`
		Serves      string `json:"serves" binding:"required"`
		Ingredients string `json:"ingredients" binding:"required"`
		Cost        string `json:"cost" binding:"required"`
	}
	var param createRecipeParam
	if err := c.ShouldBind(&param); err != nil {
		zap.S().Errorf("Failed to bind request body: %v", err)
		c.JSON(apperrors.Status(err), gin.H{
			"message":  "Recipe creation failed!",
			"required": "title, making_time, serves, ingredients, cost",
		})
		return
	}
	ctx := c.Request.Context()
	recipe := model.Recipe{
		Title:       param.Title,
		MakingTime:  param.MakingTime,
		Serves:      param.Serves,
		Ingredients: param.Ingredients,
		Cost:        param.Cost,
	}

	result, err := h.Services.RecipeService.Create(ctx, recipe)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Recipe successfully created!",
		"recipe":  []*model.Recipe{result},
	})
}

func (h *Handler) deleteRecipeByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	err = h.Services.RecipeService.DeleteByID(ctx, uint(id))
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"message": "No recipe found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Recipe successfully removed!",
	})
}

func (h *Handler) updateRecipeByID(c *gin.Context) {
	type updateRecipeParam struct {
		Title       string `json:"title"`
		MakingTime  string `json:"making_time"`
		Serves      string `json:"serves"`
		Ingredients string `json:"ingredients"`
		Cost        string `json:"cost"`
	}
	ctx := c.Request.Context()
	var param updateRecipeParam
	if err := c.Bind(&param); err != nil {
		zap.S().Errorf("Failed to bind request body: %v", err)
		c.JSON(apperrors.Status(err), gin.H{
			"message":  "Recipe update failed!",
			"required": "need to specify id",
		})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	recipe := model.Recipe{
		ID:          uint(id),
		Title:       param.Title,
		MakingTime:  param.MakingTime,
		Serves:      param.Serves,
		Ingredients: param.Ingredients,
		Cost:        param.Cost,
	}

	result, err := h.Services.RecipeService.UpdateByID(ctx, recipe)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"err":     err,
			"message": "No recipe found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Recipe successfully updated!",
		"recipe":  []*model.Recipe{result},
	})
}
