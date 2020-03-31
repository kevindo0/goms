package storage

import (
	"goms/pkg/apps"
	"github.com/gin-gonic/gin"

)

func GetKey(c *gin.Context) {
	app := apps.Gin{C: c}
	res := "good"
	app.Response(200, res)
}