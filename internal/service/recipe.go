package service

import (
	"context"

	"github.com/retail-ai-test/internal/model"
	"github.com/retail-ai-test/internal/model/response"
	"github.com/retail-ai-test/internal/repo"
)

type recipeService struct {
	repo repo.IRecipeRepo
}

func NewRecipeService(pr repo.IRecipeRepo) IRecipeService {
	return &recipeService{
		repo: pr,
	}
}

func (rs *recipeService) DeleteByID(ctx context.Context, ID uint) error {
	return rs.repo.DeleteByID(ctx, ID)
}

func (rs *recipeService) UpdateByID(ctx context.Context, recipe model.Recipe) (*model.Recipe, error) {
	return rs.repo.UpdateByID(ctx, recipe)
}

func (rs *recipeService) FindByID(ctx context.Context, ID uint) ([]*response.Recipe, error) {
	recipe, err := rs.repo.FindByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	rr := response.Recipe{
		ID:          recipe.ID,
		Title:       recipe.Title,
		MakingTime:  recipe.MakingTime,
		Serves:      recipe.Serves,
		Ingredients: recipe.Ingredients,
		Cost:        recipe.Cost,
	}
	return []*response.Recipe{&rr}, nil
}

func (rs *recipeService) FindAll(ctx context.Context) ([]*response.Recipe, error) {
	recipes, err := rs.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var rsrs []*response.Recipe
	for _, recipe := range recipes {
		r := response.Recipe{
			ID:          recipe.ID,
			Title:       recipe.Title,
			MakingTime:  recipe.MakingTime,
			Serves:      recipe.Serves,
			Ingredients: recipe.Ingredients,
			Cost:        recipe.Cost,
		}
		rsrs = append(rsrs, &r)
	}
	return rsrs, nil
}

func (rs *recipeService) Create(ctx context.Context, recipe model.Recipe) (*model.Recipe, error) {
	return rs.repo.Create(ctx, recipe)
}
