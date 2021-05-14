package model

import (
	"time"
)

type User struct {
	ID               uint   `json:"id" gorm:"primary_key"`
	Name             string `json:"name"`
	Mobile           int    `json:"mobile"`
	MobileVerifiedAt string `json:"mobile_verified_at"`
	Password         string `json:"password"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
