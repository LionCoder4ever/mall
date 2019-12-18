package jwt

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

type AuthFunc func(c *gin.Context) (interface{}, error)

func New(f AuthFunc) (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "Realmname",
		Key:           []byte("Secretkey"),
		Timeout:       time.Hour * 12,
		MaxRefresh:    time.Hour * 24,
		Authenticator: f,
		Unauthorized:  jwtUnAuthFunc,
		//LoginResponse: func(context *gin.Context, i int, s string, i2 time.Time) {
		//	context.JSON(i, gin.H{
		//		"code": i,
		//		"message": s,
		//	})
		//},
	})
	return authMiddleware, err
}

func jwtUnAuthFunc(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
