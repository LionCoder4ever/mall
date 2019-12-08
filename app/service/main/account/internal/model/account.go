package model

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	Name   string `json:"name" binding:"required"`
	Avatar string `json:"avatar"`
	AccountPrivacy
}

type AccountPrivacy struct {
	RealName string `json:"real_name"`
	IdCard   string
	Phone    string
	RegistIp string
	HandImg  string
}

func (*Account) TableName() string {
	return "account"
}
