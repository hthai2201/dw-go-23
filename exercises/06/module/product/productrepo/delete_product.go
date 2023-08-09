package productrepo

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/product/productmodel"
)

type DeleteProductStorage interface {
	Find(
		ctx context.Context,
		cond map[string]interface{},
		moreInfos ...string,
	) (*productmodel.Product, error)
	Delete(ctx context.Context, id int) error
}

type deleteProductRepo struct {
	store DeleteProductStorage
}

func NewDeleteProductRepo(store DeleteProductStorage) *deleteProductRepo {
	return &deleteProductRepo{store: store}
}

func (repo *deleteProductRepo) DeleteProduct(ctx context.Context, id int) (*productmodel.Product, error) {
	product, err := repo.store.Find(ctx, map[string]interface{}{"id": id, "status": 1})

	if product == nil {
		return nil, common.ErrCannotGetEntity(productmodel.EntityName, err)
	}

	if err := repo.store.Delete(ctx, id); err != nil {
		return nil, common.ErrCannotDeleteEntity(productmodel.EntityName, err)
	}

	return product, nil
}
