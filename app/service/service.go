package service

import (
	"mall/app/conf"
	"mall/app/internal/dao"
)

type Service struct {
	dao *dao.Dao
}

func New(c *conf.Config) *Service {
	s := &Service{
		dao.New(c.MySQL),
	}
	return s
}

func (s *Service) Close() {
	s.dao.Close()
}
