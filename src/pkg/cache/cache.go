package cache

import (
	"encoding/json"
	"fmt"

	log "goms/pkg/logging"

	"github.com/gin-gonic/gin"
)

type bodyCacheWriter struct {
	gin.ResponseWriter
	key string
}

func (w bodyCacheWriter) Write(b []byte) (int, error) {
	status := w.Status()
	if 200 == status {
		SetSting(w.key, string(b))
	}
	return w.ResponseWriter.Write(b)
}

func checkKey(key string, c *gin.Context) {
	res := GetString(key)
	// 数据从redis流出是否出错，若出错查询数据库
	var flag bool
	if len(res) > 0 {
		var jres map[string]interface{}
		err := json.Unmarshal(res, &jres)
		if err != nil {
			log.Error("cache check", log.ZError(err))
		} else {
			flag = true
			c.AbortWithStatusJSON(200, jres)
		}
	}
	if !flag {
		bcw := &bodyCacheWriter{ResponseWriter: c.Writer, key: key}
		c.Writer = bcw
		c.Next()
	}
}

func removeKey(keys []string) {
	Del(keys)
}

func Check(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("key:", key)
	}
}

func Remove(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("key1:", key)
		statusCode := c.Writer.Status()
		fmt.Println("status:", statusCode)
	}
}

func IDCheck(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println("key:", key, id)
		newKey := fmt.Sprintf("%s:%s", key, id)
		checkKey(newKey, c)
	}
}

func IDRemove(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		newKey := fmt.Sprintf("%s:%s", key, id)
		keys := []string{newKey}
		fmt.Println("newkey:", keys)
		removeKey(keys)
	}
}
