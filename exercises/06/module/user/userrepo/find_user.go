package userrepo

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/user/usermodel"
)

type FindUserStorage interface {
	FindUserByCondition(ctx context.Context, conditions map[string]interface{}, relations ...string) (*usermodel.User, error)
}

type findUserRepo struct {
	store FindUserStorage
}

func NewFindUserStorage(store FindUserStorage) *findUserRepo {
	return &findUserRepo{store: store}
}

func (repo *findUserRepo) FindUserByCondition(ctx context.Context, conditions map[string]interface{}, relations ...string) (*usermodel.User, error) {
	user, err := repo.store.FindUserByCondition(ctx, conditions)

	if err != nil {
		return nil, common.ErrCannotListEntity(usermodel.EntityName, err)
	}

	return user, nil
}
