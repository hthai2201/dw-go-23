package cartstorage

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/07/common"
	"github.com/hthai2201/dw-go-23/exercises/07/module/cart/cartmodel"
	"github.com/hthai2201/dw-go-23/exercises/07/module/product/productmodel"
)

func (s *cartMySql) UpdateProduct(ctx context.Context, data *cartmodel.CartProductUpdate) error {
	if err := s.db.Table(productmodel.ProductCreate{}.
		TableName()).Select("quantity").
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
