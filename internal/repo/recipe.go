package repo

import (
	"context"

	"github.com/retail-ai-test/internal/model"
	"gorm.io/gorm"
)

type recipeRepo struct {
	DB *gorm.DB
}

func NewRecipeRepo(db *gorm.DB) IRecipeRepo {
	return &recipeRepo{
		DB: db,
	}
}

func (*recipeRepo) DeleteByID(ctx context.Context, ID uint) (*model.Recipe, error) {
	panic("unimplemented")
}

func (*recipeRepo) UpdateByID(ctx context.Context, ID uint) (*model.Recipe, error) {
	panic("unimplemented")
}

func (pr *recipeRepo) FindByID(ctx context.Context, ID uint) (*model.Recipe, error) {
	var Recipe model.Recipe
	err := pr.DB.Table("recipes").Where("id = ?", ID).First(&Recipe).Error
	if err != nil {
		return nil, err
	}
	return &Recipe, nil
}

func (pr *recipeRepo) FindAll(ctx context.Context) ([]*model.Recipe, error) {
	var recipes []*model.Recipe
	err := pr.DB.Table("recipes").Find(&recipes).Error
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func (pr *recipeRepo) FindByTeamID(ctx context.Context, teamID uint) (*model.Recipe, error) {
	var Recipe model.Recipe
	err := pr.DB.Table("recipes").Where("team_id = ?", teamID).First(&Recipe).Error
	if err != nil {
		return nil, err
	}
	return &Recipe, nil
}

func (pr *recipeRepo) Create(ctx context.Context, recipe model.Recipe) (*model.Recipe, error) {
	if err := pr.DB.Table("recipes").Create(&recipe).Error; err != nil {
		return nil, err
	}
	return &recipe, nil
}
