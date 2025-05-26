package controllers

import (
	"bytes"
	"net/http"
	"os"

	"github.com/beevik/etree"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/libxml2/parser"
)

// XXEVuln1 XXE 漏洞示例
func XXEVuln1(c *gin.Context) {
	if c.Request.Method == "GET" {
		file, err := os.ReadFile("static/xml/xxe.xml")
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.HTML(http.StatusOK, "xxe.tpl", gin.H{
			"xxe": string(file),
		})
		return
	}

	// POST 处理
	file := c.PostForm("file")
	p := parser.New(parser.XMLParseNoEnt)
	doc, err := p.ParseReader(bytes.NewReader([]byte(file)))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer doc.Free()
	root, err := doc.DocumentElement()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	xxe := root.TextContent()
	c.HTML(http.StatusOK, "xxe.tpl", gin.H{
		"xxe": xxe,
	})
}

// XXESafe1 XXE 安全示例
func XXESafe1(c *gin.Context) {
	if c.Request.Method == "GET" {
		file, err := os.ReadFile("static/xml/xxe.xml")
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.HTML(http.StatusOK, "xxe.tpl", gin.H{
			"xxe": string(file),
		})
		return
	}

	// POST 处理
	file := c.PostForm("file")
	err := os.WriteFile("tmp/upload.xml", []byte(file), 0777)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	entityMap := make(map[string]string)
	entityMap["xxe"] = "default xxe value"
	doc := etree.NewDocument()
	doc.ReadSettings.Entity = entityMap
	if err := doc.ReadFromFile("tmp/upload.xml"); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	xxe := doc.SelectElement("root").Text()
	c.HTML(http.StatusOK, "xxe.tpl", gin.H{
		"xxe": xxe,
	})
}
