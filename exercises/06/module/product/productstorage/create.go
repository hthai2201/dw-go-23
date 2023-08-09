package productstorage

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/product/productmodel"
)

func (s *productMySql) Create(ctx context.Context, data *productmodel.ProductCreate) error {
	if err := s.db.Table(productmodel.ProductCreate{}.
		TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
