package cartstorage

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/cart/cartmodel"
)

func (s *cartMySql) RemoveProduct(ctx context.Context, productId int) error {
	if err := s.db.Table(cartmodel.Cart{}.
		TableName()).
		Where("product_id = ?", productId).
		UpdateColumns(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
