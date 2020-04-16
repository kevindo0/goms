package middleware

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func First() gin.HandlerFunc {
    return func( c *gin.Context) {
        fmt.Println("first middleware before next()")
        key := c.Query("key")
        if key == "1" { 
            // first middleware before next()
            // first middleware key==1
            // current inside of second middleware
            // **hello**
            // first middleware after next()
            // first middleware defer func
            fmt.Println("first middleware key==1")
            c.Next()
        } else if key == "2" {
            // first middleware before next()
            // first middleware key==2 before
            // first middleware key==2 after
            fmt.Println("first middleware key==2 before")
            c.Abort()
            fmt.Println("first middleware key==2 after")
            c.String(200, "key=2")
            return
        } else if key == "3" {
            // first middleware before next()
            // first middleware key==3
            // current inside of second middleware
            // **hello**
            fmt.Println("first middleware key==3")
            return
        } else if key == "4" {
            // first middleware before next()
            // first middleware key==4
            // current inside of second middleware
            // **hello**
            fmt.Println("first middleware key==4")
            c.Next()
            return
        } else {
            // first middleware before next()
            // first middleware key others
            // first middleware after next()
            // first middleware defer func
            // current inside of second middleware
            // **hello**
            fmt.Println("first middleware key others")
        }

        fmt.Println("first middleware after next()")

        defer func() {
            fmt.Println("first middleware defer func")
        }()
    }
}

func Second() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("current inside of second middleware")
    }
}
