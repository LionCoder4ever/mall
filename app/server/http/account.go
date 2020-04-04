package http

import (
	"github.com/gin-gonic/gin"
	"mall/app/internal/model"
	"mall/library/ecode"
	"mall/library/log"
	"strconv"
)

func (h *httpServer) CreateAccount(c *gin.Context) {
	acc := new(model.Account)
	if err := c.ShouldBindJSON(acc); err != nil {
		log.Logger.Errorf("json parse failed %s", err.Error())
		c.Render(h.WrapResponse(nil, ecode.ParameterErr))
		return
	}
	uid, err := h.srv.CreateAccount(acc)
	if err != nil {
		c.Render(h.WrapResponse(nil, err))
		return
	}
	c.Render(h.WrapResponse(uid, err))
}

func (h *httpServer) ReadAccount(c *gin.Context) {
	var (
		uid int64
		err error
	)
	if uid, err = strconv.ParseInt(c.Param("uid"), 10, 64); err != nil {
		c.Render(h.WrapResponse(nil, err))
		return
	}
	result, err := h.srv.ReadAccount(uid)
	if err != nil {
		c.Render(h.WrapResponse(nil, err))
		return
	}
	c.Render(h.WrapResponse(result, err))
}

func (h *httpServer) UpdateAccount(c *gin.Context) {

}

func (h *httpServer) DeleteAccount(c *gin.Context) {
	var (
		uid int64
		err error
	)
	if uid, err = strconv.ParseInt(c.Query("id"), 10, 64); err != nil {
		log.Logger.Errorf("get id from url failed cause: %s", err.Error())
		c.Render(h.WrapResponse(nil, ecode.ParameterErr))
		return
	}
	err = h.srv.DeleteAccount(uid)
	if err != nil {
		log.Logger.Errorf("del row failed cause: %s", err.Error())
		c.Render(h.WrapResponse(nil, err))
		return
	}
	c.Render(h.WrapResponse(nil, nil))
}
