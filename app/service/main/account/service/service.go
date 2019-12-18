package service

import (
	"mall/app/service/main/account/conf"
	"mall/app/service/main/account/internal/dao"
	"mall/app/service/main/account/internal/model"
)

type Service interface {
	Close()
	GetAccount(int) (*model.Account, error)
	DelAccount(int) error
	CreateAccount(*model.Account) (uint, error)
	Login(string, string) (uint, error)
}

type service struct {
	dao *dao.Dao
}

func New(c *conf.Config) Service {
	s := &service{
		dao: dao.New(c.MySQL),
	}
	return s
}

func (s *service) Close() {
	s.dao.Close()
}
