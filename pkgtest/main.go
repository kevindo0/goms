package main

import (
    "fmt"
    "net/http"
    "pkgte/middleware"
    "github.com/gin-gonic/gin"
)

func main(){
    r := gin.Default()
    r.Use(middleware.First(), middleware.Second())
    r.GET("/hello", func(c *gin.Context) {
        fmt.Println("**hello**")
        c.String(http.StatusOK, "hello")
    })
    r.Run()
}
