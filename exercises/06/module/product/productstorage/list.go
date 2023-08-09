package productstorage

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/product/productmodel"
)

func (store *productMySql) List(
	ctx context.Context,
	paging *common.Paging,
) ([]productmodel.Product, error) {

	var products []productmodel.Product
	db := store.db.Table(productmodel.Product{}.TableName())

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := store.db.
		Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).
		Find(&products).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return products, nil
}
