package controllers

import (
	"go-sec-code/utils"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SSRFVuln1 SSRF 漏洞示例1 - 直接请求
func SSRFVuln1(c *gin.Context) {
	url := c.DefaultQuery("url", "http://www.example.com")
	res, err := http.Get(url)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", body)
}

// SSRFVuln2 SSRF 漏洞示例2 - 绕过黑名单
// bypass can be :
// http://LOCALHOST:233
// http://localhost.:233
// http://0:233
// and others
func SSRFVuln2(c *gin.Context) {
	url := c.DefaultQuery("url", "http://www.example.com")
	ssrfFilter := utils.SSRFFilter{}
	blacklists := []string{"localhost", "127.0.0.1"}
	evil := ssrfFilter.DoBlackFilter(url, blacklists)
	if evil == true {
		c.String(http.StatusBadRequest, "evil input")
	} else {
		res, err := http.Get(url)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", body)
	}
}

// SSRFVuln3 SSRF 漏洞示例3 - 302跳转
func SSRFVuln3(c *gin.Context) {
	url := c.DefaultQuery("url", "http://www.example.com")
	ssrfFilter := utils.SSRFFilter{}
	evil := ssrfFilter.DoGogsFilter(url)
	if evil == true {
		c.String(http.StatusBadRequest, "evil input")
	} else {
		res, err := http.Get(url)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", body)
	}
}

// SSRFSafe1 SSRF 安全示例 - 白名单
func SSRFSafe1(c *gin.Context) {
	url := c.DefaultQuery("url", "http://www.example.com")
	ssrfFilter := utils.SSRFFilter{}
	whitelists := []string{"example.com"}
	evil := ssrfFilter.DoWhiteFilter(url, whitelists)
	if evil == true {
		c.String(http.StatusBadRequest, "evil input")
	} else {
		res, err := http.Get(url)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", body)
	}
}
