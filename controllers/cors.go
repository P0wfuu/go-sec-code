package controllers

import (
	"go-sec-code/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CorsVuln1 CORS漏洞示例1 - 反射Origin
func CorsVuln1(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	jsonp := make(map[string]interface{})
	jsonp["username"] = "admin"
	jsonp["password"] = "admin@123"
	c.JSON(http.StatusOK, jsonp)
}

// CorsVuln2 CORS漏洞示例2 - 允许任意来源并设置凭证
func CorsVuln2(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	jsonp := make(map[string]interface{})
	jsonp["username"] = "admin"
	jsonp["password"] = "admin@123"
	c.JSON(http.StatusOK, jsonp)
}

// CorsSafe1 CORS安全示例 - 域名白名单
func CorsSafe1(c *gin.Context) {
	origin := c.Request.Header.Get("origin")
	whitelists := []string{"localhost:233", "example.com"}
	corsFilter := utils.CorsFilter{}
	if origin != "" && corsFilter.DoFilter(origin, whitelists) {
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
	}
	jsonp := make(map[string]interface{})
	jsonp["username"] = "admin"
	jsonp["password"] = "admin@123"
	c.JSON(http.StatusOK, jsonp)
}
