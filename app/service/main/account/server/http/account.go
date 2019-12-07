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
	if idFromQuery, err = strconv.Atoi(c.Query("id")); err != nil {
		h.JSON(c, nil, err.Error())
		return
	}
	result := h.srv.GetAccount(idFromQuery)
	if result != nil {
		h.JSON(c, result, "")
		return
	}
	h.JSON(c, nil, "id not found")
}

func (h *hs) CreateAccount(c *gin.Context) {
	var id uint
	acc := new(model.Account)
	if err := c.ShouldBindJSON(acc); err != nil {
		h.JSON(c, nil, "json parse failed")
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
