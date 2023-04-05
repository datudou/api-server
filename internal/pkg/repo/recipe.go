package repo

import (
	"context"
	"fmt"

	"github.com/retail-ai-test/internal/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type recipeRepo struct {
	db *gorm.DB
}

func NewRecipeRepo(db *gorm.DB) IRecipeRepo {
	return &recipeRepo{
		db: db,
	}
}

func (rp *recipeRepo) FindByID(ctx context.Context, ID uint) (*model.Recipe, error) {
	var recipe model.Recipe
	err := rp.db.WithContext(ctx).Where("id = ?", ID).First(&recipe).Error
	if err != nil {
		return nil, err
	}
	return &recipe, nil
}

func (rp *recipeRepo) FindAll(ctx context.Context) ([]*model.Recipe, error) {
	var recipes []*model.Recipe
	tx := rp.db.WithContext(ctx).Find(&recipes)
	if tx.RowsAffected < 1 {
		return nil, fmt.Errorf("no rows found")
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	return recipes, nil
}

func (rp *recipeRepo) Create(ctx context.Context, recipe model.Recipe) (*model.Recipe, error) {
	if err := rp.db.WithContext(ctx).Create(&recipe).Error; err != nil {
		return nil, err
	}
	return &recipe, nil
}

func (rp *recipeRepo) DeleteByID(ctx context.Context, id uint) error {

	recipe := model.Recipe{Model: gorm.Model{ID: id}}
	tx := rp.db.WithContext(ctx).Delete(&recipe)
	if tx.RowsAffected < 1 {
		return fmt.Errorf("row with id=%d cannot be deleted because it doesn't exist", id)
	}
	return nil
}

func (rp *recipeRepo) UpdateByID(ctx context.Context, recipe model.Recipe) (*model.Recipe, error) {

	tx := rp.db.Table(recipe.TableName()).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", recipe.ID).
		Select("title", "making_time", "serves", "ingredients", "cost").
		Updates(map[string]interface{}{
			"title":       recipe.Title,
			"making_time": recipe.MakingTime,
			"serves":      recipe.Serves,
			"ingredients": recipe.Ingredients,
			"cost":        recipe.Cost,
		})

	if tx.RowsAffected < 1 {
		return nil, fmt.Errorf("row with id=%d cannot be updated ,cannot be updated may be same params or not match the constraint", recipe.ID)
	}
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &recipe, nil
}
