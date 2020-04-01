package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/app/internal/model"
	"strconv"
)

func (h *httpServer) CreateAccount(c *gin.Context) {
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

func (h *httpServer) ReadAccount(c *gin.Context) {
	var (
		idFromQuery int
		err         error
	)
	if idFromQuery, err = strconv.Atoi(c.Param("id")); err != nil {
		h.JSON(c, nil, err.Error())
		return
	}
	result, err := h.srv.ReadAccount(idFromQuery)
	if err != nil {
		h.JSON(c, nil, err.Error())
		return
	}
	h.JSON(c, result, "")
}

func (h *httpServer) UpdateAccount(c *gin.Context) {

}

func (h *httpServer) DeleteAccount(c *gin.Context) {
	var (
		idFromQuery int
		err         error
	)
	if idFromQuery, err = strconv.Atoi(c.Query("id")); err != nil {
		h.JSON(c, nil, fmt.Sprintf("get id from url failed cause: %s", err.Error()))
		return
	}
	err = h.srv.DeleteAccount(idFromQuery)
	if err != nil {
		h.JSON(c, nil, fmt.Sprintf("del row failed cause: %s", err.Error()))
		return
	}
	h.JSON(c, nil, "")
}