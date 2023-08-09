package cartmodel

type Checkout struct {
	Total float64 `json:"total" gorm:"column:total;"`
}

func (Checkout) TableName() string { return Cart{}.TableName() }
