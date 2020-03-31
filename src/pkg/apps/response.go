package apps

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"goms/pkg/merr"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(code int, message interface{}) {
	if (message == nil) {
		message = merr.GetMsg(code)
	}
	g.C.JSON(http.StatusOK, gin.H{
		"code": code,
		"message": message,
	})
}