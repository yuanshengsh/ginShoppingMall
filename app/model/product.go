package model

import (
	"time"
)

type Product struct {
	ID        int   `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Type      int    `json:"type"`
	Price     string `json:"price"`
	Status    int    `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Product) TableName() string {
	return "product"
}
