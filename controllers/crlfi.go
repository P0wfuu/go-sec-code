package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CRLFSafe1 CRLF 注入安全示例
func CRLFSafe1(c *gin.Context) {
	header := c.Query("header")
	c.Header("header", header)
	c.String(http.StatusOK, "")
}
