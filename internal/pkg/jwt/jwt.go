package jwt

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AuthFunc func(c *gin.Context) (interface{}, error)

func New(f AuthFunc) (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "Realmname",
		Key:           []byte("Secretkey"),
		Timeout:       time.Minute * 30,
		MaxRefresh:    time.Hour * 24,
		Authenticator: f,
		Unauthorized:  jwtUnAuthFunc,
	})
	return authMiddleware, err
}

func jwtUnAuthFunc(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    nil,
	})
}
