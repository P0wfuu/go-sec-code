package controllers

import (
	"fmt"
	"go-sec-code/utils"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// FileUploadVuln1 文件上传漏洞示例
func FileUploadVuln1(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "fileUpload.tpl", nil)
		return
	}

	// POST 处理
	userid := c.PostForm("userid")
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	savePath := filepath.Join("static/upload/", userid+fmt.Sprint(time.Now().Unix())+file.Filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "fileUpload.tpl", gin.H{
		"savePath": savePath,
	})
}

// FileUploadSafe1 文件上传安全示例
func FileUploadSafe1(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "fileUpload.tpl", nil)
		return
	}

	// POST 处理
	userid := c.PostForm("userid")
	fileUploadFilter := utils.FileUploadFilter{}
	evil := fileUploadFilter.DoFilter(userid)
	if evil == true {
		c.String(http.StatusBadRequest, "evil input")
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	savePath := filepath.Join("static/upload/", userid+fmt.Sprint(time.Now().Unix())+file.Filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "fileUpload.tpl", gin.H{
		"savePath": savePath,
	})
}
