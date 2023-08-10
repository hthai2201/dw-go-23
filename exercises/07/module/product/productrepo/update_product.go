package productrepo

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/07/common"
	"github.com/hthai2201/dw-go-23/exercises/07/module/product/productmodel"
)

type UpdateProductStorage interface {
	Find(
		ctx context.Context,
		cond map[string]interface{},
		moreInfos ...string,
	) (*productmodel.Product, error)
	Update(ctx context.Context, data *productmodel.ProductUpdate) error
}

type updateProductRepo struct {
	store UpdateProductStorage
}

func NewUpdateProductRepo(store UpdateProductStorage) *updateProductRepo {
	return &updateProductRepo{store: store}
}

func (repo *updateProductRepo) UpdateProduct(ctx context.Context, data *productmodel.ProductUpdate) error {
	product, err := repo.store.Find(ctx, map[string]interface{}{"id": data.ID, "status": 1})

	if product == nil {
		return common.ErrCannotGetEntity(productmodel.EntityName, err)
	}
	err = repo.store.Update(ctx, data)

	if err != nil {
		return common.ErrCannotUpdateEntity(productmodel.EntityName, err)
	}

	return nil
}
