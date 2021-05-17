package model

import (
	"time"
)

type ProductSpecification struct {
	ID        int    `json:"id" gorm:"primary_key"`
	ProductId int    `json:"product_id"`
	SpecId    int    `json:"spec_id"`
	Name      string `json:"name"`
	Type      string `json:"price"`
	Setting   string `json:"setting"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ProductSpecification) TableName() string {
	return "product_specification"
}
