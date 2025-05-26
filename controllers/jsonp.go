package controllers

import (
	"encoding/json"
	"go-sec-code/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JsonpVuln1 JSONP 漏洞示例1 - 无检查
func JsonpVuln1(c *gin.Context) {
	callback := c.Query("callback")
	c.Header("Content-Type", "application/javascript")
	jsonp := make(map[string]interface{})
	jsonp["username"] = "admin"
	jsonp["password"] = "admin@123"
	data, err := json.Marshal(jsonp)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	output := callback + "(" + string(data) + ")"
	c.String(http.StatusOK, output)
}

// JsonpVuln2 JSONP 漏洞示例2 - 空 Referer
func JsonpVuln2(c *gin.Context) {
	callback := c.Query("callback")
	referer := c.Request.Header.Get("referer")
	jsonpFilter := utils.JsonpFilter{}
	whitelists := []string{"localhost:233", "example.com"}
	if referer == "" || jsonpFilter.DoFilter(referer, whitelists) {
		c.Header("Content-Type", "application/javascript")
		jsonp := make(map[string]interface{})
		jsonp["username"] = "admin"
		jsonp["password"] = "admin@123"
		data, err := json.Marshal(jsonp)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		output := callback + "(" + string(data) + ")"
		c.String(http.StatusOK, output)
	} else {
		c.String(http.StatusBadRequest, "evil input")
	}
}

// JsonpSafe1 JSONP 安全示例
func JsonpSafe1(c *gin.Context) {
	callback := c.Query("callback")
	referer := c.Request.Header.Get("referer")
	jsonpFilter := utils.JsonpFilter{}
	whitelists := []string{"localhost:233", "example.com"}
	if referer != "" && jsonpFilter.DoFilter(referer, whitelists) {
		c.Header("Content-Type", "application/javascript")
		jsonp := make(map[string]interface{})
		jsonp["username"] = "admin"
		jsonp["password"] = "admin@123"
		data, err := json.Marshal(jsonp)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		output := callback + "(" + string(data) + ")"
		c.String(http.StatusOK, output)
	} else {
		c.String(http.StatusBadRequest, "evil input")
	}
}
