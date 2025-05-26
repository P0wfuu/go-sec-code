package controllers

import (
	"go-sec-code/utils"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// PathTraversalVuln1 路径遍历漏洞示例1 - 无过滤
func PathTraversalVuln1(c *gin.Context) {
	file := c.Query("file")
	output, err := os.ReadFile(file)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Data(http.StatusOK, "text/plain", output)
}

// PathTraversalVuln2 路径遍历漏洞示例2 - 使用Clean但仍有漏洞
func PathTraversalVuln2(c *gin.Context) {
	file := c.Query("file")
	file = filepath.Clean(file)
	output, err := os.ReadFile(file)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Data(http.StatusOK, "text/plain", output)
}

// PathTraversalSafe1 路径遍历安全示例1 - 使用过滤器
func PathTraversalSafe1(c *gin.Context) {
	file := c.Query("file")
	pathTraversalFilter := utils.PathTraversalFilter{}
	if pathTraversalFilter.DoFilter(file) {
		c.String(http.StatusBadRequest, "evil input")
	} else {
		output, err := os.ReadFile("static/" + file)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Data(http.StatusOK, "text/plain", output)
	}
}

// PathTraversalSafe2 路径遍历安全示例2 - 路径检查
func PathTraversalSafe2(c *gin.Context) {
	file := c.Query("file")
	file = filepath.Join("static/", file)
	if !strings.HasPrefix(file, "static/") {
		c.String(http.StatusBadRequest, "evil input")
	} else {
		output, err := os.ReadFile(file)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Data(http.StatusOK, "text/plain", output)
	}
}
