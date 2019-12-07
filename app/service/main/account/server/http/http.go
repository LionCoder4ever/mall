package http

import (
	"github.com/gin-gonic/gin"
	"mall/app/service/main/account/service"
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

	r := gin.Default()
	r.GET("/getacc", h.GetAccount)
	r.GET("/delacc", h.DelAccount)
	r.POST("/createacc", h.CreateAccount)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func (h *hs) JSON(c *gin.Context, data interface{}, errMsg string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": errMsg,
		"data":    data,
	})
}
