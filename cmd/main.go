package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/internal/app/account"
	"mall/internal/pkg/conf"
	"mall/internal/pkg/dao"
	"mall/internal/pkg/http"
	"mall/internal/pkg/jwt"
	. "mall/internal/pkg/log"
	"mall/internal/pkg/log/middleware"
	"mall/internal/pkg/uuid"
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
	// init dao
	dao.NewDao(conf.Conf.MySQL)
	// init business service
	account.NewAccountService()
	// init http server
	NewServer()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		Logger.Infof("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			// close db
			dao.Close()
			Logger.Info("service exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func NewServer() {
	j, err := jwt.New(jwtAuthFunc)
	if err != nil {
		panic(fmt.Sprintf("jwt middleware crate failed %s", err.Error()))
	}

	r := gin.New()
	r.Use(middleware.LoggerWithZap())
	r.Use(middleware.RecoveryWithZap())

	v1 := r.Group(http.Apiversion)
	auth := v1.Group("/auth")
	auth.Use(j.MiddlewareFunc())
	{
		auth.GET("/del", account.Destroy)
		auth.GET("/account/:uid", account.Show)
	}

	v1.POST("/login", j.LoginHandler)
	v1.POST("/register", account.Store)
	go r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type Login struct {
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func jwtAuthFunc(c *gin.Context) (interface{}, error) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		return "", fmt.Errorf("error payload")
	}
	uid, err := account.Login(json.Phone, json.Password)
	if err != nil {
		return nil, err
	}
	return uid, nil
}
