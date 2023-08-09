package cartmodel

import (
	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/product/productmodel"
)

type CartProductUpdate struct {
	common.SQLModel `json:",inline"`
	ProductId       int                  `json:"product_id" gorm:"column:product_id;"`
	Product         productmodel.Product `gorm:"foreignKey:product_id;references:id"`
	Quantity        int                  `json:"quantity" gorm:"column:quantity;"`
}

func (CartProductUpdate) TableName() string { return Cart{}.TableName() }
