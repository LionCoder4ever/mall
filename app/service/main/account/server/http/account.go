package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/app/service/main/account/internal/model"
	"strconv"
)

func (h *hs) GetAccount(c *gin.Context) {
	var (
		idFromQuery int
		err         error
	)
	if idFromQuery, err = strconv.Atoi(c.Param("id")); err != nil {
		h.JSON(c, nil, err.Error())
		return
	}
	result, err := h.srv.GetAccount(idFromQuery)
	if err != nil {
		h.JSON(c, nil, err.Error())
		return
	}
	h.JSON(c, result, "")
}

func (h *hs) CreateAccount(c *gin.Context) {
	var id uint
	acc := new(model.Account)
	if err := c.ShouldBindJSON(acc); err != nil {
		h.JSON(c, nil, fmt.Sprintf("json parse failed %s", err.Error()))
		return
	}
	id, err := h.srv.CreateAccount(acc)
	if err != nil {
		h.JSON(c, nil, err.Error())
		return
	}
	h.JSON(c, id, "")
}

func (h *hs) DelAccount(c *gin.Context) {
	var (
		idFromQuery int
		err         error
	)
	if idFromQuery, err = strconv.Atoi(c.Query("id")); err != nil {
		h.JSON(c, nil, fmt.Sprintf("get id from url failed cause: %s", err.Error()))
		return
	}
	err = h.srv.DelAccount(idFromQuery)
	if err != nil {
		h.JSON(c, nil, fmt.Sprintf("del row failed cause: %s", err.Error()))
		return
	}
	h.JSON(c, nil, "")
}
