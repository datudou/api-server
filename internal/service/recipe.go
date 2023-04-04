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

func (rs *recipeService) UpdateByID(ctx context.Context, recipe model.Recipe) (*response.Recipe, error) {
	res, err := rs.repo.UpdateByID(ctx, recipe)
	if err != nil {
		return nil, err
	}
	r := response.Recipe{
		ID:          res.ID,
		Title:       res.Title,
		MakingTime:  res.MakingTime,
		Serves:      res.Serves,
		Ingredients: res.Ingredients,
		Cost:        res.Cost,
	}
	return &r, nil
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

func (rs *recipeService) Create(ctx context.Context, recipe model.Recipe) (*response.Recipe, error) {
	res, err := rs.repo.Create(ctx, recipe)
	if err != nil {
		return nil, err
	}
	r := response.Recipe{
		ID:          res.ID,
		Title:       res.Title,
		MakingTime:  res.MakingTime,
		Serves:      res.Serves,
		Ingredients: res.Ingredients,
		Cost:        res.Cost,
	}
	return &r, nil
}
