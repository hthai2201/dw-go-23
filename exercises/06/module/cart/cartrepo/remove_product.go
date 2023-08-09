package cartrepo

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/cart/cartmodel"
)

type RemoveCardProductStorage interface {
	FindProduct(
		ctx context.Context,
		cond map[string]interface{},
		moreInfos ...string,
	) (*cartmodel.Cart, error)
	RemoveProduct(ctx context.Context, productId int) error
}

type removeCardProductRepo struct {
	store RemoveCardProductStorage
}

func NewRemoveCardProductRepo(store RemoveCardProductStorage) *removeCardProductRepo {
	return &removeCardProductRepo{store: store}
}

func (repo *removeCardProductRepo) RemoveCartProduct(ctx context.Context, data *cartmodel.CartProductRemove) (*cartmodel.Cart, error) {
	cart, err := repo.store.FindProduct(ctx, map[string]interface{}{"product_id": data.ProductId, "status": 1})

	if cart == nil {
		return nil, common.ErrCannotGetEntity(cartmodel.EntityName, err)
	}

	if err := repo.store.RemoveProduct(ctx, data.ProductId); err != nil {
		return nil, common.ErrCannotDeleteEntity(cartmodel.EntityName, err)
	}

	return cart, nil
}
