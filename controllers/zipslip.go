package controllers

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// ZipSlipVuln1 Zip Slip 漏洞示例
func ZipSlipVuln1(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "fileUpload.tpl", nil)
		return
	}

	// POST 处理
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	timestamp := fmt.Sprint(time.Now().Unix())
	savePath := filepath.Join("static/upload/", timestamp+file.Filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	unzipPath := filepath.Join("static/unzip/", timestamp+file.Filename)
	r, err := zip.OpenReader(savePath)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(unzipPath, f.Name)
		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.HTML(http.StatusOK, "fileUpload.tpl", gin.H{
		"savePath": unzipPath,
	})
}
