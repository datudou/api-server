package service

import (
	"context"

	"github.com/retail-ai-test/internal/model"
	"github.com/retail-ai-test/internal/model/response"
	"github.com/retail-ai-test/internal/repo"
)

type IRecipeService interface {
	FindByID(ctx context.Context, ID uint) ([]*response.Recipe, error)
	FindAll(ctx context.Context) ([]*response.Recipe, error)
	UpdateByID(ctx context.Context, recipe model.Recipe) (*model.Recipe, error)
	Create(ctx context.Context, recipe model.Recipe) (*model.Recipe, error)
	DeleteByID(ctx context.Context, ID uint) error
}

type IUserService interface {
	FindByID(ctx context.Context, userID string) (*response.User, error)
	UpdateByID(ctx context.Context, user model.User) (*response.User, error)
	Create(ctx context.Context, user model.User) (*response.User, error)
	DeleteByID(ctx context.Context, userID string) error
}

type Services struct {
	RecipeService IRecipeService
	UserService   IUserService
}

type Deps struct {
	Repos    *repo.Repositories
	Services *Services
}

func NewServices(deps Deps) *Services {
	rs := NewRecipeService(deps.Repos.RecipeRepo)
	us := NewUserService(deps.Repos.UserRepo)
	return &Services{
		RecipeService: rs,
		UserService:   us,
	}
}
