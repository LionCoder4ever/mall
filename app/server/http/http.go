package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/pkg/errors"
	"mall/app/service"
	"mall/library/ecode"
	"mall/library/jwt"
	"mall/library/log/middleware"
	"net/http"
)

const apiversion = "/v1"

var jsonContentType = []string{"application/json; charset=utf-8"}

// http server
type httpServer struct {
	srv *service.Service
}

// http
type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func writeJSON(w http.ResponseWriter, data interface{}) (err error) {
	var jsonBytes []byte
	writeContentType(w, jsonContentType)
	if jsonBytes, err = json.Marshal(data); err != nil {
		err = errors.WithStack(err)
		return
	}
	if _, err = w.Write(jsonBytes); err != nil {
		err = errors.WithStack(err)
	}
	return
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

func (r Response) Render(w http.ResponseWriter) (err error) {
	return writeJSON(w, r)
}

// WriteContentType write json ContentType.
func (r Response) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
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

	v1 := r.Group(apiversion)
	auth := v1.Group("/auth")
	auth.Use(j.MiddlewareFunc())
	{
		auth.GET("/del", h.DeleteAccount)
		auth.GET("/account/:uid", h.ReadAccount)
	}

	v1.POST("/login", j.LoginHandler)
	v1.POST("/register", h.CreateAccount)
	//v1.POST("/shop/register", h.CreateShop)
	//v1.GET("/shop/info/:id", h.GetShop)
	go r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

/**
transfer err code to msg
*/
func (h *httpServer) WrapResponse(data interface{}, err error) (int, render.Render) {
	code := http.StatusOK
	customCode := ecode.Cause(err)
	return code, Response{
		Code:    customCode.Code(),
		Message: customCode.Message(),
		Data:    data,
	}
}

func (h *httpServer) JSON(c *gin.Context, data interface{}, errMsg string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": errMsg,
		"data":    data,
	})
}

type Login struct {
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (h *httpServer) jwtAuthFunc(c *gin.Context) (interface{}, error) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		return "", fmt.Errorf("error payload")
	}
	uid, err := h.srv.Login(json.Phone, json.Password)
	if err != nil {
		return nil, err
	}
	return uid, nil
}
