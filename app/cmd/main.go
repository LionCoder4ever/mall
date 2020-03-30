package main

import (
	"flag"
	"fmt"
	"mall/app/conf"
	"mall/app/server/http"
	"mall/app/service"
	"mall/library/log"
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
	log.New(conf.Conf.Log)
	log.Info("conf load success", "from dsn ", conf.Conf.MySQL.DSN)
	log.Info("logger init success")
	// init service
	svc := service.New(&conf.Conf)
	// init http server
	http.New(svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			svc.Close()
			log.Info("service exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
