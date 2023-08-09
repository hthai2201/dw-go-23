package productmodel

type ProductUpdate struct {
	ID       int     `json:"id"`
	Name     string  `json:"name" gorm:"column:name;"`
	Price    float64 `json:"price" gorm:"column:price;"`
	Quantity int     `json:"quantity" gorm:"column:quantity;"`
}

func (ProductUpdate) TableName() string { return Product{}.TableName() }
