package cartstorage

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/cart/cartmodel"
)

func (s *cartMySql) AddProduct(ctx context.Context, data *cartmodel.CartProductAdd) error {
	if err := s.db.Table(cartmodel.CartProductAdd{}.
		TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
