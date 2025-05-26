package controllers

import (
	"fmt"
	"go-sec-code/utils"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// CommandInjectVuln1 命令注入漏洞示例1
func CommandInjectVuln1(c *gin.Context) {
	dir := c.Query("dir")
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Data(http.StatusOK, "text/plain", out)
}

type AdConfig struct {
	AdAccount  string `json:"ad_account"`
	AdPassword string `json:"ad_password"`
}

// CommandInjectVuln2 命令注入漏洞示例2
func CommandInjectVuln2(c *gin.Context) {
	var adConfig AdConfig
	if err := c.ShouldBindJSON(&adConfig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input := fmt.Sprintf("add_user %s %s", adConfig.AdAccount, adConfig.AdPassword)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Data(http.StatusOK, "text/plain", out)
}

// CommandInjectVuln3 命令注入漏洞示例3
func CommandInjectVuln3(c *gin.Context) {
	repoUrl := c.DefaultQuery("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Data(http.StatusOK, "text/plain", out)
}

// CommandInjectSafe1 命令注入安全示例1
func CommandInjectSafe1(c *gin.Context) {
	dir := c.Query("dir")
	commandInjectFilter := utils.CommandInjectFilter{}
	evil := commandInjectFilter.DoFilter(dir)
	if evil == false {
		c.String(http.StatusBadRequest, "evil input")
		return
	}
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Data(http.StatusOK, "text/plain", out)
}
