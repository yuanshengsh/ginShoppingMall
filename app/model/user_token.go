package model

import (
	boot "ginShoppingMall/bootstrap"
	"strconv"
	"time"
)

type UserToken struct {
	Token     string `json:"token"`
	UserId    int    `json:"user_id"`
	ExpiresAt string `json:"expires_at"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func SaveUserToken(ut *UserToken) error {
	currentTime := time.Now()
	m, _ := time.ParseDuration("+" + strconv.Itoa(boot.Config.Token.JwtTokenCreatedExpireAt) + "m")
	ut.ExpiresAt = currentTime.Add(m).Format("2006-01-02 15:04:05")

	if err := boot.DB.Create(ut).Error; err != nil {
		return err
	}
	return nil
}

func (UserToken) TableName() string {
	return "user_token"
}
