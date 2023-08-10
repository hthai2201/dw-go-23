package productstorage

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/07/common"
	"github.com/hthai2201/dw-go-23/exercises/07/module/product/productmodel"
)

func (s *productMySql) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(productmodel.Product{}.
		TableName()).
		Where("id = ?", id).
		UpdateColumns(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
