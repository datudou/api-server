package repo

import (
	"context"

	"github.com/retail-ai-test/internal/model"
	"gorm.io/gorm"
)

type IRecipeRepo interface {
	FindByID(ctx context.Context, ID uint) (*model.Recipe, error)
	FindAll(ctx context.Context) ([]*model.Recipe, error)
	UpdateByID(ctx context.Context, ID uint) (*model.Recipe, error)
	Create(ctx context.Context, Recipe model.Recipe) (*model.Recipe, error)
	DeleteByID(ctx context.Context, ID uint) (*model.Recipe, error)
}

type Repositories struct {
	RecipeRepo IRecipeRepo
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		RecipeRepo: NewRecipeRepo(db),
	}
}
