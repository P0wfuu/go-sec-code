package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MainPage 主页处理函数
func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tpl", nil)
}

// MainPagePost 处理POST请求
func MainPagePost(c *gin.Context) {
	foo := c.Query("foo")
	c.String(http.StatusOK, foo)
}
