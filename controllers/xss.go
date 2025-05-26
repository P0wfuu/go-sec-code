package controllers

import (
	"go-sec-code/utils"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// XSSVuln1 XSS 漏洞示例1 - 反射型
func XSSVuln1(c *gin.Context) {
	xss := c.DefaultQuery("xss", "hello")
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, xss)
}

// XSSVuln2 XSS 漏洞示例2 - 存储型
func XSSVuln2(c *gin.Context) {
	if c.Request.Method == "GET" {
		xss := c.GetString("xss")
		if xss == "" {
			xss = "hello"
		}
		c.HTML(http.StatusOK, "xss.tpl", gin.H{
			"xss": template.HTML(xss),
		})
		return
	}

	// POST 处理
	xss := c.DefaultPostForm("xss", "hello")
	c.SetCookie("xss", xss, 3600, "/", "", false, false)
	c.HTML(http.StatusOK, "xss.tpl", gin.H{
		"xss": template.HTML(xss),
	})
}

// XSSVuln3 XSS 漏洞示例3 - SVG
func XSSVuln3(c *gin.Context) {
	file, err := os.ReadFile("static/xss/poc.svg")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Data(http.StatusOK, "image/svg+xml", file)
}

// XSSVuln4 XSS 漏洞示例4 - PDF
func XSSVuln4(c *gin.Context) {
	file, err := os.ReadFile("static/xss/poc.pdf")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Security-Policy", `default-src 'self';`)
	c.Data(http.StatusOK, "application/pdf", file)
}

// XSSSafe1 XSS 安全示例1 - 过滤
func XSSSafe1(c *gin.Context) {
	xss := c.DefaultQuery("xss", "hello")
	xssFilter := utils.XSSFilter{}
	xss = xssFilter.DoFilter(xss)
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, xss)
}

// XSSSafe2 XSS 安全示例2 - CSP
func XSSSafe2(c *gin.Context) {
	file, err := os.ReadFile("static/xss/poc.svg")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Security-Policy", `default-src 'self';`)
	c.Data(http.StatusOK, "image/svg+xml", file)
}
