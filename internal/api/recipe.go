package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/retail-ai-test/internal/model"
	"github.com/retail-ai-test/internal/model/request"
	"github.com/retail-ai-test/internal/model/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (h *Handler) getRecipeByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		badRequestErrorRes(c, response.ErrorMessage{
			Message: "Recipe details not found",
		})
		return
	}
	result, err := h.Services.RecipeService.FindByID(ctx, uint(id))
	if err != nil {
		badRequestErrorRes(c, response.ErrorMessage{
			Message: "Recipe details not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe details by id",
		"recipe":  result,
	})
}

func (h *Handler) getRecipes(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := h.Services.RecipeService.FindAll(ctx)
	if err != nil {
		badRequestErrorRes(c, response.ErrorMessage{
			Message: "Recipe details not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"recipes": result,
	})
}

func (h *Handler) createRecipe(c *gin.Context) {
	var param request.CreateRecipeParam
	if err := c.ShouldBindJSON(&param); err != nil {
		badRequestErrorRes(c, response.ErrorMessage{
			Message:  "Recipe creation failed!",
			Required: "title, making_time, serves, ingredients, cost",
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
		internalErrorRes(c, response.ErrorMessage{
			Message: "Recipe creation failed!",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Recipe successfully created!",
		"recipe":  []*model.Recipe{result},
	})
}

func (h *Handler) deleteRecipeByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		badRequestErrorRes(c, response.ErrorMessage{
			Message: "Invalid recipe id!",
		})
		return
	}
	err = h.Services.RecipeService.DeleteByID(ctx, uint(id))
	if err != nil {
		zap.S().Errorf("Error while deleting recipe: %v", err)
		internalErrorRes(c, response.ErrorMessage{
			Message: "Recipe not found!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe successfully removed!",
	})
}

func (h *Handler) updateRecipeByID(c *gin.Context) {

	var param request.UpdateRecipeParam
	if err := c.ShouldBindJSON(&param); err != nil {
		badRequestErrorRes(c, response.ErrorMessage{
			Message: "Recipe update failed!",
		})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		badRequestErrorRes(c, response.ErrorMessage{
			Message: "Invalid recipe id!",
		})
		return
	}
	recipe := model.Recipe{
		Model:       gorm.Model{ID: uint(id)},
		Title:       param.Title,
		MakingTime:  param.MakingTime,
		Serves:      param.Serves,
		Ingredients: param.Ingredients,
		Cost:        param.Cost,
	}

	ctx := c.Request.Context()
	result, err := h.Services.RecipeService.UpdateByID(ctx, recipe)
	if err != nil {
		internalErrorRes(c, response.ErrorMessage{
			Message: "Recipe update failed!",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Recipe successfully updated!",
		"recipe":  []*model.Recipe{result},
	})
}
