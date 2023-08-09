package cartstorage

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/cart/cartmodel"
)

func (s *cartMySql) FindProduct(
	ctx context.Context,
	cond map[string]interface{},
	moreInfos ...string,
) (*cartmodel.Cart, error) {
	db := s.db.Table(cartmodel.Cart{}.TableName())
	for i := range moreInfos {
		db = db.Preload(moreInfos[i])
	}
	db.Where(cond)

	var rs cartmodel.Cart

	if err := db.First(&rs).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &rs, nil
}
