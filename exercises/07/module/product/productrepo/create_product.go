package productrepo

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/07/common"
	"github.com/hthai2201/dw-go-23/exercises/07/module/product/productmodel"
)

type CreateProductStorage interface {
	Create(ctx context.Context, data *productmodel.ProductCreate) error
}

type createProductRepo struct {
	store CreateProductStorage
}

func NewCreateProductRepo(store CreateProductStorage) *createProductRepo {
	return &createProductRepo{store: store}
}

func (repo *createProductRepo) CreateProduct(ctx context.Context, data *productmodel.ProductCreate) error {

	err := repo.store.Create(ctx, data)

	if err != nil {
		return common.ErrCannotCreateEntity(productmodel.EntityName, err)
	}

	return nil
}
