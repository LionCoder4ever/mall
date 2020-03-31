package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/app/service"
	"mall/library/jwt"
	"mall/library/log/middleware"
	"net/http"
)

// http server
type httpServer struct {
	srv *service.Service
}

// Init init
func New(svc *service.Service) {
	h := &httpServer{
		srv: svc,
	}
	j, err := jwt.New(h.jwtAuthFunc)
	if err != nil {
		panic(fmt.Sprintf("jwt middleware crate failed %s", err.Error()))
	}
	r := gin.New()
	r.Use(middleware.LoggerWithZap())
	r.Use(middleware.RecoveryWithZap())
	v1 := r.Group("/v1")

	auth := v1.Group("/auth")
	auth.Use(j.MiddlewareFunc())
	{
		auth.GET("/del", h.DelAccount)
		auth.GET("/account/:id", h.GetAccount)
	}

	v1.POST("/login", j.LoginHandler)
	v1.POST("/register", h.CreateAccount)
	//v1.POST("/shop/register", h.CreateShop)
	//v1.GET("/shop/info/:id", h.GetShop)
	go r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func (h *httpServer) JSON(c *gin.Context, data interface{}, errMsg string) {
	// TODO custom error status code
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": errMsg,
		"data":    data,
	})
}

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	Id       uint
}

func (h *httpServer) jwtAuthFunc(c *gin.Context) (interface{}, error) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		return "", fmt.Errorf("error payload")
	}
	uid, err := h.srv.Login(json.User, json.Password)
	if err != nil {
		return nil, err
	}
	return uid, nil
}
