package account

import (
	"github.com/gin-gonic/gin"
	"mall/internal/pkg/ecode"
	"mall/internal/pkg/http"
	"mall/internal/pkg/log"
	"strconv"
)

var accService Service

func NewAccountService() {
	accService = &service{}
}

func Store(c *gin.Context) {
	acc := new(Account)
	if err := c.ShouldBindJSON(acc); err != nil {
		log.Logger.Errorf("json parse failed %s", err.Error())
		c.Render(http.WrapResponse(nil, ecode.ParameterErr))
		return
	}
	uid, err := accService.CreateAccount(acc)
	if err != nil {
		c.Render(http.WrapResponse(nil, err))
		return
	}
	c.Render(http.WrapResponse(uid, err))
}

func Show(c *gin.Context) {
	var (
		uid int64
		err error
	)
	if uid, err = strconv.ParseInt(c.Param("uid"), 10, 64); err != nil {
		c.Render(http.WrapResponse(nil, err))
		return
	}
	result, err := accService.ReadAccount(uid)
	if err != nil {
		c.Render(http.WrapResponse(nil, err))
		return
	}
	result.Password = ""
	result.AccountPrivacy = AccountPrivacy{}
	c.Render(http.WrapResponse(result, err))
}

func Index(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Destroy(c *gin.Context) {
	var (
		uid int64
		err error
	)
	if uid, err = strconv.ParseInt(c.Query("id"), 10, 64); err != nil {
		log.Logger.Errorf("get id from url failed cause: %s", err.Error())
		c.Render(http.WrapResponse(nil, ecode.ParameterErr))
		return
	}
	err = accService.DeleteAccount(uid)
	if err != nil {
		log.Logger.Errorf("del row failed cause: %s", err.Error())
		c.Render(http.WrapResponse(nil, err))
		return
	}
	c.Render(http.WrapResponse(nil, nil))
}

func Login(phone string, password string) (uid uint, err error) {
	return accService.Login(phone, password)
}
