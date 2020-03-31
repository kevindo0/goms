package middleware

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// 全局中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				res := gin.H{
					"code": 1,
					"message": fmt.Sprintf("%v", err),
				}
				c.JSON(http.StatusOK, res)
			}
		}()
		c.Next()
	}
}