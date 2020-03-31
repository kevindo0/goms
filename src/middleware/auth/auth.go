package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Auth func 56")
		c.Next()
	}
}