package grpc

import (
	"context"
	"fmt"
	"github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
	"mall/app/service/main/account/api"
	"mall/app/service/main/account/internal/model"
	"mall/app/service/main/account/service"
	"mall/library/log"
	"net"
	"time"
)

type AccountServer struct {
	srv    service.Service
	tracer *zipkin.Tracer
}

func NewTracer() *zipkin.Tracer {
	// create a reporter to be used by the tracer
	reporter := httpreporter.NewReporter("http://localhost:9411/api/v2/spans")
	// set-up the local endpoint for our service
	//TODO get endpoint from env ?
	endpoint, err := zipkin.NewEndpoint("accountService", "127.0.0.1:9013")
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}
	// initialize the tracer
	tracer, err := zipkin.NewTracer(
		reporter,
		zipkin.WithLocalEndpoint(endpoint),
	)
	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
	}

	return tracer
}

func New(srv service.Service) {
	acc := &AccountServer{srv, NewTracer()}
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9013))
	if err != nil {
		log.Infof("failed to listen: %v", err)
	}
	log.Infof("grpc server listen on %s", lis.Addr())
	// inject zipkin handler
	g := grpc.NewServer(grpc.StatsHandler(zipkingrpc.NewServerHandler(acc.tracer)))
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
	// start span
	span := s.tracer.StartSpan("getuserid")
	span.Annotate(time.Now(), fmt.Sprintf("ready back id: {%d}", t.AccountInfo.Model.Id))
	span.Finish()
	return &t.AccountInfo, err
}
