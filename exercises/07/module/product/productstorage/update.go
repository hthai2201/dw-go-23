package productstorage

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/07/common"
	"github.com/hthai2201/dw-go-23/exercises/07/module/product/productmodel"
)

func (s *productMySql) Update(ctx context.Context, data *productmodel.ProductUpdate) error {
	if err := s.db.Table(productmodel.ProductCreate{}.
		TableName()).Select("name", "price", "quantity").
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
