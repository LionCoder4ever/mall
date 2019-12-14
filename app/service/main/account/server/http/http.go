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
	v1 := r.Group("/v1")
	v1.GET("/acc/:id", h.GetAccount)
	v1.GET("/delacc", h.DelAccount)
	v1.POST("/createacc", h.CreateAccount)
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
