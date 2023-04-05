package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/retail-ai-test/internal/pkg/model"
	"github.com/retail-ai-test/internal/pkg/repo/mock_repo"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type recipeHelper struct {
	repo *mock_repo.MockIRecipeRepo
	// service *mock_service.MockIRecipeService
}

func newRecipeHelper(t *testing.T) *recipeHelper {
	ctrl := gomock.NewController(t)
	h := &recipeHelper{}
	h.repo = mock_repo.NewMockIRecipeRepo(ctrl)
	// h.service = service.NewRecipeService(h.repo).(*mock_service.MockIRecipeService)
	return h
}

func Test_Recipe_FindByID(t *testing.T) {
	rh := newRecipeHelper(t)

	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{

			name:    "find recipe success by id",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rh.repo.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&model.Recipe{
				Model: gorm.Model{
					ID: 1,
				},
				Title:       "test",
				MakingTime:  "test",
				Serves:      "test",
				Ingredients: "test",
				Cost:        10,
			}, nil)
			recipe, err := rh.repo.FindByID(tt.args.ctx, tt.args.id)
			fmt.Println(recipe)
			if tt.wantErr {
				require.NotNil(t, err)
			}
			require.Equal(t, recipe.Title, "test")
			require.Equal(t, recipe.ID, uint(1))
		})
	}
}
