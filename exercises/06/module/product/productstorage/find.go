package productstorage

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/product/productmodel"
)

func (s *productMySql) Find(
	ctx context.Context,
	cond map[string]interface{},
	moreInfos ...string,
) (*productmodel.Product, error) {
	db := s.db.Table(productmodel.Product{}.TableName()).Where(cond)

	for i := range moreInfos {
		db = db.Preload(moreInfos[i])
	}

	var rs productmodel.Product

	if err := db.First(&rs).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &rs, nil
}
