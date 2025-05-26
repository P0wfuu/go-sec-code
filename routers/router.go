package routers

import (
	"go-sec-code/controllers"

	"github.com/gin-gonic/gin"
)

// InitRoutes 初始化所有路由
func InitRoutes(r *gin.Engine) {
	// 静态文件
	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./favicon.ico")

	// 设置模板路径
	r.LoadHTMLGlob("views/*")

	// 主页
	r.GET("/", controllers.MainPage)

	// 命令注入
	r.GET("/commandInject/vuln", controllers.CommandInjectVuln1)
	r.GET("/commandInject/vuln/host", controllers.CommandInjectVuln2)
	r.GET("/commandInject/vuln/git", controllers.CommandInjectVuln3)
	r.GET("/commandInject/safe", controllers.CommandInjectSafe1)

	// CORS
	r.GET("/cors/vuln/reflect", controllers.CorsVuln1)
	r.GET("/cors/vuln/any-origin-with-credential", controllers.CorsVuln2)
	r.GET("/cors/safe", controllers.CorsSafe1)

	// CRLF注入
	r.GET("/crlfInjection/safe", controllers.CRLFSafe1)

	// 文件上传
	r.GET("/fileUpload/vuln", controllers.FileUploadVuln1)
	r.GET("/fileUpload/safe", controllers.FileUploadSafe1)

	// JSONP
	r.GET("/jsonp/vuln/noCheck", controllers.JsonpVuln1)
	r.GET("/jsonp/vuln/emptyReferer", controllers.JsonpVuln1)
	r.GET("/jsonp/safe", controllers.JsonpSafe1)

	// 路径遍历
	r.GET("/pathTraversal/vuln", controllers.PathTraversalVuln1)
	r.GET("/pathTraversal/vuln/clean", controllers.PathTraversalVuln2)
	r.GET("/pathTraversal/safe/filter", controllers.PathTraversalSafe1)
	r.GET("/pathTraversal/safe/check", controllers.PathTraversalSafe2)

	// SQL注入
	r.GET("/sqlInjection/native/vuln/integer", controllers.SqlInjectionVuln1)
	r.GET("/sqlInjection/native/vuln/string", controllers.SqlInjectionVuln2)
	r.GET("/sqlInjection/orm/vuln/xorm", controllers.SqlInjectionVuln3)
	r.GET("/sqlInjection/generator/vuln/squirrel", controllers.SqlInjectionVuln4)
	r.GET("/sqlInjection/native/safe/integer", controllers.SqlInjectionSafe1)
	r.GET("/sqlInjection/native/safe/string", controllers.SqlInjectionSafe2)
	r.GET("/sqlInjection/orm/safe/beego", controllers.SqlInjectionSafe3)

	// SSRF
	r.GET("/ssrf/vuln", controllers.SSRFVuln1)
	r.GET("/ssrf/vuln/obfuscation", controllers.SSRFVuln2)
	r.GET("/ssrf/vuln/302", controllers.SSRFVuln3)
	r.GET("/ssrf/safe/whitelists", controllers.SSRFSafe1)

	// SSTI
	r.GET("/ssti/vuln", controllers.SSTIVuln1)
	r.GET("/ssti/safe", controllers.SSTISafe1)

	// XSS
	r.GET("/xss/vuln", controllers.XSSVuln1)
	r.GET("/xss/vuln/store", controllers.XSSVuln2)
	r.GET("/xss/vuln/svg", controllers.XSSVuln3)
	r.GET("/xss/vuln/pdf", controllers.XSSVuln4)
	r.GET("/xss/safe", controllers.XSSSafe1)
	r.GET("/xss/safe/svg", controllers.XSSSafe2)

	// XXE
	r.GET("/xxe/vuln", controllers.XXEVuln1)
	r.GET("/xxe/safe", controllers.XXESafe1)

	// ZipSlip
	r.GET("/zipslip/vuln", controllers.ZipSlipVuln1)
}
