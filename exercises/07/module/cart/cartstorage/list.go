package cartstorage

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/07/common"
	"github.com/hthai2201/dw-go-23/exercises/07/module/cart/cartmodel"
)

func (store *cartMySql) List(
	ctx context.Context,
	paging *common.Paging,
) ([]cartmodel.Cart, error) {

	var carts []cartmodel.Cart
	db := store.db.Table(cartmodel.Cart{}.TableName())

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := store.db.
		Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).Preload("Product").
		Find(&carts).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return carts, nil
}
