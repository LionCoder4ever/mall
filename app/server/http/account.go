package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/app/internal/model"
	"strconv"
)

func (h *httpServer) CreateAccount(c *gin.Context) {
	acc := new(model.Account)
	if err := c.ShouldBindJSON(acc); err != nil {
		h.JSON(c, nil, fmt.Sprintf("json parse failed %s", err.Error()))
		return
	}
	uid, err := h.srv.CreateAccount(acc)
	if err != nil {
		h.JSON(c, nil, err.Error())
		return
	}
	h.JSON(c, uid, "")
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
		h.JSON(c, nil, fmt.Sprintf("get id from url failed cause: %s", err.Error()))
		return
	}
	err = h.srv.DeleteAccount(uid)
	if err != nil {
		h.JSON(c, nil, fmt.Sprintf("del row failed cause: %s", err.Error()))
		return
	}
	h.JSON(c, nil, "")
}
