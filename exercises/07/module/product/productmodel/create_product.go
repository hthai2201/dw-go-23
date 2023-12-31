package productmodel

import "github.com/hthai2201/dw-go-23/exercises/07/common"

type ProductCreate struct {
	common.SQLModel `json:",inline"`
	Name            string  `json:"name" gorm:"column:name;"`
	Price           float64 `json:"price" gorm:"column:price;"`
	Quantity        int     `json:"quantity" gorm:"column:quantity;"`
}

func (ProductCreate) TableName() string { return Product{}.TableName() }
