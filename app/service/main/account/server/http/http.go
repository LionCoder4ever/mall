package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/app/service/main/account/service"
	"mall/library/jwt"
	"net/http"
)

// http server
type hs struct {
	srv service.Service
}

// Init init
func New(svc service.Service) {
	h := &hs{
		srv: svc,
	}
	j, err := jwt.New(h.jwtAuthFunc)
	if err != nil {
		panic(fmt.Sprintf("jwt middleware crate failed %s", err.Error()))
	}
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.POST("/login", j.LoginHandler)
	v1.GET("/acc/:id", h.GetAccount)
	v1.POST("/createacc", h.CreateAccount)

	auth := v1.Group("/auth")
	auth.Use(j.MiddlewareFunc())
	{
		auth.GET("/delacc", h.DelAccount)
	}
	go r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func (h *hs) JSON(c *gin.Context, data interface{}, errMsg string) {
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

func (h *hs) jwtAuthFunc(c *gin.Context) (interface{}, error) {
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
