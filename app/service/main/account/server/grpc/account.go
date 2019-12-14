package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mall/app/service/main/account/api"
	"mall/app/service/main/account/internal/model"
	"mall/app/service/main/account/service"
	"mall/library/log"
	"net"
)

type AccountServer struct {
	srv service.Service
}

func New(srv service.Service) {
	acc := &AccountServer{srv}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9013))
	if err != nil {
		log.Infof("failed to listen: %v", err)
	}
	log.Infof("grpc server listen on %s", lis.Addr())
	g := grpc.NewServer()
	api.RegisterAccountServer(g, acc)
	go func() {
		if err := g.Serve(lis); err != nil {
			panic(fmt.Sprintf("failed to listen: %v", err))
		}
		log.Info("grpc server init success")
	}()
}

func (s *AccountServer) Read(c context.Context, id *api.Id) (*api.AccountInfo, error) {
	var (
		t   *model.Account
		err error
	)
	t, err = s.srv.GetAccount(int(id.Id))
	t.AccountInfo.Model = &api.Model{Id: t.ID}
	return &t.AccountInfo, err
}
