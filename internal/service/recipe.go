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

func (rs *recipeService) Create(ctx context.Context, recipe model.Recipe) (*model.Recipe, error) {
	return rs.repo.Create(ctx, recipe)
}

func (*recipeService) DeleteByID(ctx context.Context, ID uint) (*model.Recipe, error) {
	panic("unimplemented")
}

func (*recipeService) UpdateByID(ctx context.Context, ID uint) (*model.Recipe, error) {
	panic("unimplemented")
}

func (rs *recipeService) FindByID(ctx context.Context, ID uint) (*model.Recipe, error) {
	return rs.repo.FindByID(ctx, ID)
}

func (rs *recipeService) FindAll(ctx context.Context) ([]*model.Recipe, error) {
	return rs.repo.FindAll(ctx)
}
