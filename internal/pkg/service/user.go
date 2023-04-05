package service

import (
	"context"

	"github.com/retail-ai-test/internal/pkg/model"
	"github.com/retail-ai-test/internal/pkg/model/response"
	"github.com/retail-ai-test/internal/pkg/repo"
)

type userService struct {
	repo repo.IUserRepo
}

// Create implements IUserService
func (ur *userService) Create(ctx context.Context, user model.User) (*response.User, error) {
	res, err := ur.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	u := response.User{
		UserID:   res.UserID,
		NickName: res.NickName,
		Comment:  res.Comment,
	}
	return &u, nil
}

func (ur *userService) DeleteByID(ctx context.Context, userID string) error {
	return ur.repo.DeleteByID(ctx, userID)
}

func (ur *userService) FindByID(ctx context.Context, userID string) (*response.User, error) {
	user, err := ur.repo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	u := response.User{
		UserID:   user.UserID,
		NickName: user.NickName,
		Comment:  user.Comment,
	}
	return &u, nil
}

func (ur *userService) UpdateByID(ctx context.Context, user model.User) (*response.User, error) {
	res, err := ur.repo.UpdateByID(ctx, user)
	if err != nil {
		return nil, err
	}
	u := response.User{
		UserID:   res.UserID,
		NickName: res.NickName,
		Comment:  res.Comment,
	}
	return &u, nil
}

func NewUserService(ur repo.IUserRepo) IUserService {
	return &userService{
		repo: ur,
	}
}
