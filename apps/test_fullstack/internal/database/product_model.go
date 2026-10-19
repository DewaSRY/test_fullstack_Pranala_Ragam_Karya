package database

import (
	_ "gorm.io/gorm"
)

type ProductModel struct {
	ID    int32   `gorm:"type:int;primary_key`
	Name  string  `gorm:"type:varchar(255)`
	Price float32 `gorm:"type:double`
	Stock int32   `gorm:"type:int`
}

func (ProductModel) TableName() string {
	return "product"
}
