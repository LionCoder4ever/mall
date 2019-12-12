package grpc

import (
	"context"
	"mall/app/service/main/account/api"
	"mall/app/service/main/account/internal/model"
	"mall/app/service/main/account/service"
)

type AccountServer struct {
	srv service.Service
}

func New(srv service.Service) *AccountServer {
	acc := &AccountServer{srv}
	return acc
}

func (s *AccountServer) Read(c context.Context, id *api.Id) (*api.AccountInfo, error) {
	var (
		t   *model.Account
		err error
	)
	t, err = s.srv.GetAccount(int(id.Id))
	return &t.AccountInfo, err
}
