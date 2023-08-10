package cartrepo

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/07/common"
	"github.com/hthai2201/dw-go-23/exercises/07/module/cart/cartmodel"
)

type ListCartsStorage interface {
	List(
		ctx context.Context,
		paging *common.Paging,
	) ([]cartmodel.Cart, error)
}

type listCartsRepo struct {
	store ListCartsStorage
}

func NewListStorage(store ListCartsStorage) *listCartsRepo {
	return &listCartsRepo{store: store}
}

func (repo *listCartsRepo) List(ctx context.Context, paging *common.Paging) ([]cartmodel.Cart, error) {
	carts, err := repo.store.List(ctx, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(cartmodel.EntityName, err)
	}

	return carts, nil
}
