package routers

import (
	"github.com/gin-gonic/gin"
	"goms/middleware/auth"
	"goms/apis/login"
	"goms/apis/storage"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.POST("/login", auth.AuthCheck(), login.Login)
	r.GET("get", storage.GetKey)
	return r
}