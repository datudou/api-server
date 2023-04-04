package repo

import (
	"context"

	"github.com/retail-ai-test/internal/model"
	"gorm.io/gorm"
)

type IRecipeRepo interface {
	FindByID(ctx context.Context, ID uint) (*model.Recipe, error)
	FindAll(ctx context.Context) ([]*model.Recipe, error)
	UpdateByID(ctx context.Context, recipe model.Recipe) (*model.Recipe, error)
	Create(ctx context.Context, Recipe model.Recipe) (*model.Recipe, error)
	DeleteByID(ctx context.Context, ID uint) error
}

type IUserRepo interface {
	FindByID(ctx context.Context, userID string) (*model.User, error)
	UpdateByID(ctx context.Context, user model.User) (*model.User, error)
	Create(ctx context.Context, user model.User) (*model.User, error)
	DeleteByID(ctx context.Context, userID string) error
	Validate(ctx context.Context, userID string, password string) (bool, error)
}

type Repositories struct {
	RecipeRepo IRecipeRepo
	UserRepo   IUserRepo
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		RecipeRepo: NewRecipeRepo(db),
		UserRepo:   NewUserRepo(db),
	}
}
