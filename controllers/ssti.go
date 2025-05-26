package controllers

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/Masterminds/sprig"
	"github.com/gin-gonic/gin"
)

// SSTIVuln1 SSTI 漏洞示例
func SSTIVuln1(c *gin.Context) {
	os.Setenv("go-sec-code-secret-key", "b81024f158eefcf60792ae9df9524f82")
	usertemplate := c.DefaultQuery("template", "please send your template")
	t := template.New("ssti").Funcs(sprig.FuncMap())
	t, _ = t.Parse(usertemplate)
	buff := bytes.Buffer{}
	err := t.Execute(&buff, struct{}{})
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	data, err := io.ReadAll(&buff)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "ssti.tpl", gin.H{
		"usertemplate": string(data),
	})
}

// SSTISafe1 SSTI 安全示例
func SSTISafe1(c *gin.Context) {
	usertemplate := c.DefaultQuery("template", "please send your template")
	c.HTML(http.StatusOK, "ssti.tpl", gin.H{
		"usertemplate": usertemplate,
	})
}
