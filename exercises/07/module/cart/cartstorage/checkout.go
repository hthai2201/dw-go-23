package cartstorage

import (
	"context"
	"fmt"

	"github.com/hthai2201/dw-go-23/exercises/07/common"
	"github.com/hthai2201/dw-go-23/exercises/07/module/cart/cartmodel"
	"github.com/hthai2201/dw-go-23/exercises/07/module/product/productmodel"
)

func (s *cartMySql) Checkout(ctx context.Context, c *cartmodel.Checkout) error {
	cartTableName := cartmodel.Cart{}.
		TableName()
	productTableName := productmodel.Product{}.TableName()
	if result := s.db.Table(cartTableName).Select("SUM(products.price * cart.quantity) as total").
		Where("cart.status = ?", 1).
		Joins(fmt.Sprintf("JOIN %s ON %s.product_id = %s.id", productTableName, cartTableName, productTableName)).
		Scan(&c); result.Error != nil {
		return common.ErrDB(result.Error)
	}

	if err := s.db.Table(cartTableName).
		Where("status = ?", 1).
		Update("status", 2).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
