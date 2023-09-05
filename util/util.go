package util

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"time"
)

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-ID")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}

func GetGmtIso8601(expireEnd int64) string {
	tokenExpire := time.Unix(expireEnd, 0).Format("2006-01-02T15:04:05Z")
	return tokenExpire
}

func StringBuilder(strings ...string) string {
	//创建字节缓冲
	var stringBuilder bytes.Buffer
	for _, v := range strings {
		stringBuilder.WriteString(v)
	}
	return stringBuilder.String()
}
