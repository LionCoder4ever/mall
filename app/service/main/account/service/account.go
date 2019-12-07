package service

import (
	"mall/app/service/main/account/internal/model"
)

func (s *service) GetAccount(id int) *model.Account {
	return s.dao.GetAccount(uint(id))
}

func (s *service) CreateAccount(acc *model.Account) (uint, error) {
	return s.dao.CreateAccount(acc)
}

func (s *service) DelAccount(id int) error {
	return s.dao.DelAccount(uint(id))
}
