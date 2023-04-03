package service

import (
	"context"

	"github.com/retail-ai-test/internal/model"
	"github.com/retail-ai-test/internal/repo"
)

type IRecipeService interface {
	FindByID(ctx context.Context, ID uint) (*model.Recipe, error)
	FindAll(ctx context.Context) ([]*model.Recipe, error)
	UpdateByID(ctx context.Context, recipe model.Recipe) (*model.Recipe, error)
	Create(ctx context.Context, recipe model.Recipe) ([]*model.Recipe, error)
	DeleteByID(ctx context.Context, ID uint) error
}

type Services struct {
	RecipeService IRecipeService
}

type Deps struct {
	Repos    *repo.Repositories
	Services *Services
}

func NewServices(deps Deps) *Services {
	rs := NewRecipeService(deps.Repos.RecipeRepo)
	return &Services{
		RecipeService: rs,
	}
}
