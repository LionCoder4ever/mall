package main

import (
	"flag"
	"fmt"
	"mall/app/conf"
	"mall/app/server/http"
	"mall/app/service"
	. "mall/library/log"
	"mall/library/uuid"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	flag.Parse()
	if err := conf.Load(); err != nil {
		panic(fmt.Sprintf("conf load failed %s", err.Error()))
	}
	// init log
	NewLogger(conf.Conf.Log)
	Logger.Info("conf load success ", "from dsn ", conf.Conf.MySQL.DSN)
	Logger.Info("logger init success")
	// init uuid
	if err := uuid.NewUUID(); err != nil {
		Logger.Errorf("service exit, uuid generator failed because %s", err.Error())
		return
	}
	// init service
	svc := service.New(&conf.Conf)
	// init http server
	http.New(svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		Logger.Infof("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			svc.Close()
			Logger.Info("service exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
