package productrepo

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/07/common"
	"github.com/hthai2201/dw-go-23/exercises/07/module/product/productmodel"
)

type ListProductsStorage interface {
	List(
		ctx context.Context,
		paging *common.Paging,
	) ([]productmodel.Product, error)
}

type listProductsRepo struct {
	store ListProductsStorage
}

func NewListStorage(store ListProductsStorage) *listProductsRepo {
	return &listProductsRepo{store: store}
}

func (repo *listProductsRepo) List(ctx context.Context, paging *common.Paging) ([]productmodel.Product, error) {
	products, err := repo.store.List(ctx, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(productmodel.EntityName, err)
	}

	return products, nil
}
