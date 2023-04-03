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

func (rp *recipeRepo) DeleteByID(ctx context.Context, ID uint) error {
	recipe := model.Recipe{ID: ID}
	err := rp.DB.Table("recipes").Delete(&recipe).Error
	if err != nil {
		return err
	}
	return nil
}

func (rp *recipeRepo) UpdateByID(ctx context.Context, recipe model.Recipe) (*model.Recipe, error) {
	err := rp.DB.Table("recipes").
		Set("gorm:query_option", "FOR update").
		Where("id = ?", recipe.ID).
		Select("title", "making_time", "serves", "ingredients").
		Updates(map[string]interface{}{
			"title":       recipe.Title,
			"making_time": recipe.MakingTime,
			"serves":      recipe.Serves,
			"ingredients": recipe.Ingredients,
		}).
		Error
	if err != nil {
		return nil, err
	}
	return &recipe, nil
}

func (pr *recipeRepo) FindByID(ctx context.Context, ID uint) (*model.Recipe, error) {
	var recipe model.Recipe
	err := pr.DB.Table("recipes").Where("id = ?", ID).First(&recipe).Error
	if err != nil {
		return nil, err
	}
	return &recipe, nil
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
	if err := pr.DB.Create(&recipe).Error; err != nil {
		return nil, err
	}
	return &recipe, nil
}
