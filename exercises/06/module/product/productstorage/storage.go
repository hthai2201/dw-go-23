package productstorage

import "gorm.io/gorm"

type productMySql struct {
	db *gorm.DB
}

func NewMysql(db *gorm.DB) *productMySql {
	return &productMySql{db: db}
}
