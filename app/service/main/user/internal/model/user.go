package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Uid    string `json:"uid"`
	Name   string `json:"name" binding:"required"`
	Avatar string `json:"avatar"`
	Gender string `json:"gender"`
}

func (*User) TableName() string {
	return "user"
}
