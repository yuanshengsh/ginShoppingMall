package middleware

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
