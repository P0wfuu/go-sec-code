package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// 注意：在使用 r.StaticFile("/favicon.ico", "./favicon.ico") 后不需要这个处理函数了
// 保留这里是为了示例如何在 Gin 中处理静态文件

// Favicon 处理 favicon 请求
func Favicon(c *gin.Context) {
	icon, err := os.ReadFile("favicon.ico")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading favicon")
		return
	}
	c.Data(http.StatusOK, "image/x-icon", icon)
}
