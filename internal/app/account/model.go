package account

import (
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model    `json:"-"`
	UId           int64  `json:"uid" gorm:"index:uid;unique_index;not null"`
	Name          string `json:"name" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Password      string `json:"password,omitempty" binding:"required"`
	PasswordReapt string `json:"repeat,omitempty" binding:"required" gorm:"-"`
	Avatar        string `json:"avatar"`
	AccountPrivacy
}

type AccountPrivacy struct {
	RealName string `json:",omitempty"`
	IdCard   string `json:",omitempty"`
	Phone    string `json:",omitempty" binding:"required"`
	RegistIp string `json:",omitempty"`
	HandImg  string `json:",omitempty"`
}
