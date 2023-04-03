package service

import (
	"context"

	"github.com/retail-ai-test/internal/model"
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

func (rs *recipeService) FindByID(ctx context.Context, ID uint) (*model.Recipe, error) {
	return rs.repo.FindByID(ctx, ID)
}

func (rs *recipeService) FindAll(ctx context.Context) ([]*model.Recipe, error) {
	recipes, err := rs.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func (rs *recipeService) Create(ctx context.Context, recipe model.Recipe) ([]*model.Recipe, error) {
	r, err := rs.repo.Create(ctx, recipe)
	if err != nil {
		return nil, err
	}
	return []*model.Recipe{r}, nil
}
