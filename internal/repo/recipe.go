package repo

import (
	"context"

	"github.com/retail-ai-test/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type recipeRepo struct {
	DB *gorm.DB
}

func NewRecipeRepo(db *gorm.DB) IRecipeRepo {
	return &recipeRepo{
		DB: db,
	}
}

func (rp *recipeRepo) FindByID(ctx context.Context, ID uint) (*model.Recipe, error) {
	var recipe model.Recipe
	err := rp.DB.Table("recipes").Where("id = ?", ID).First(&recipe).Error
	if err != nil {
		return nil, err
	}
	return &recipe, nil
}

func (rp *recipeRepo) FindAll(ctx context.Context) ([]*model.Recipe, error) {
	var recipes []*model.Recipe
	err := rp.DB.Table("recipes").Find(&recipes).Error
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func (rp *recipeRepo) Create(ctx context.Context, recipe model.Recipe) (*model.Recipe, error) {
	tx := rp.DB.Begin()
	if err := tx.Create(&recipe).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &recipe, nil
}

func (rp *recipeRepo) DeleteByID(ctx context.Context, ID uint) error {
	tx := rp.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	recipe := model.Recipe{ID: ID}
	err := tx.Table("recipes").Delete(&recipe).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (rp *recipeRepo) UpdateByID(ctx context.Context, recipe model.Recipe) (*model.Recipe, error) {
	tx := rp.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	err := tx.Table("recipes").
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", recipe.ID).
		Select("title", "making_time", "serves", "ingredients", "cost").
		Updates(map[string]interface{}{
			"title":       recipe.Title,
			"making_time": recipe.MakingTime,
			"serves":      recipe.Serves,
			"ingredients": recipe.Ingredients,
			"cost":        recipe.Cost,
		}).
		Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &recipe, nil
}
