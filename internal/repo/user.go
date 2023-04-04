package repo

import (
	"context"
	"fmt"

	"github.com/retail-ai-test/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) Create(ctx context.Context, user model.User) (*model.User, error) {
	if err := ur.db.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepo) DeleteByID(ctx context.Context, userID string) error {
	user := model.User{}
	tx := ur.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&user)
	if tx.RowsAffected < 1 {
		return fmt.Errorf("row with id=%s cannot be deleted because it doesn't exist", userID)
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (ur *userRepo) FindByID(ctx context.Context, userID string) (*model.User, error) {
	var user model.User
	err := ur.db.WithContext(ctx).Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepo) UpdateByID(ctx context.Context, user model.User) (*model.User, error) {

	tx := ur.db.Table(user.TableName()).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("user_id = ?", user.UserID).
		Select("nick_name", "comment").
		Updates(map[string]interface{}{
			"nick_name": user.NickName,
			"comment":   user.Comment,
		})
	if tx.RowsAffected < 1 {
		return nil, fmt.Errorf("row with id=%s cannot be updated may be same params or not match the constraint", user.UserID)
	}

	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (ur *userRepo) Validate(ctx context.Context, userID string, password string) (bool, error) {
	var user model.User
	err := ur.db.WithContext(ctx).Where("user_id = ? and password = ?", userID, password).
		First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
