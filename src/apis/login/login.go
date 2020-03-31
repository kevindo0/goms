package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goms/pkg/merr"
	"goms/service/slogin"
	"goms/pkg/apps"
)

func Login(c *gin.Context) {
	app := apps.Gin{C: c}
	fmt.Println("login")
	username := c.PostForm("username")
	pwd := c.PostForm("password")
	data := slogin.Login(username, pwd)
	fmt.Println("data:", data)
	if (!data) {
		app.Response(merr.LOGIN_ERROR, "username or password is wrong")
		return
	}
	app.Response(merr.SUCCESS, nil)
}