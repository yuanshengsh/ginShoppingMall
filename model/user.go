package model

type User struct {
	ID               int64  `json:"id" gorm:"primary_key"`
	Name             string `json:"name"`
	Mobile           int    `json:"mobile"`
	MobileVerifiedAt string `json:"mobile_verified_at"`
	Password         string `json:"password"`
}
