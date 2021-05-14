package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
)

//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"io/ioutil"
//	"net/http"
//	"strings"
//	"time"
//)
//
//func Logger(body string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
//		c.Request.Body.Close()
//		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
//		hjson, _ := json.Marshal(c.Request.Header)
//		start := time.Now()
//		c.Next()
//		end := time.Now()
//		log.Info("| %3d | %13v | %15s | %s  %s | %s| %s",
//			c.Writer.Status(),
//			end.Sub(start),
//			c.ClientIP(),
//			c.Request.Method,
//			c.Request.RequestURI,
//			hjson,
//			string(bodyBytes),
//		)
//		c.Set("body", body)
//		c.Next()
//	}
//}

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("haha")
				//打印错误堆栈信息
				log.Printf("panic: %v\n", r)
				debug.PrintStack()
				//封装通用json返回
				ctx.JSON(200, gin.H{
					"code": "4444",
					"msg":  "服务器内部错误",
				})
			}
		}()
		ctx.Next()
	}
}
