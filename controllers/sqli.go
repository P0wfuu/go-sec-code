package controllers

import (
	"database/sql"
	"fmt"
	"go-sec-code/models"
	"net/http"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	// 移除 beego orm 初始化
}

// SqlInjectionVuln1 SQL注入漏洞示例1 - 整数型注入
func SqlInjectionVuln1(c *gin.Context) {
	id := c.Query("id")
	db, err := sql.Open("mysql", source)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()
	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// SqlInjectionVuln2 SQL注入漏洞示例2 - 字符型注入
func SqlInjectionVuln2(c *gin.Context) {
	username := c.Query("username")
	db, err := sql.Open("mysql", source)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()
	sqlStr := fmt.Sprintf("select * from user where username=\"%s\"", username)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// SqlInjectionVuln3 SQL注入漏洞示例3 - ORM注入
func SqlInjectionVuln3(c *gin.Context) {
	username := c.Query("username")
	field := c.Query("field")
	engine, err := xorm.NewEngine("mysql", source)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	engine.ShowSQL(true)
	user := models.User{}
	session := engine.Prepare().And(fmt.Sprintf("%s like ?", field), username)
	ok, err := session.Get(&user)
	if !ok && err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// SqlInjectionVuln4 SQL注入漏洞示例4 - SQL生成器注入
func SqlInjectionVuln4(c *gin.Context) {
	username := c.Query("username")
	order := c.Query("order")
	db, err := sql.Open("mysql", source)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()
	expression := sq.Select("*").From("user").Where(sq.Eq{"username": username}).OrderBy(order)
	sqlStr, args, err := expression.ToSql()
	fmt.Println(sqlStr)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	user := models.User{}
	err = db.QueryRow(sqlStr, args...).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// SqlInjectionSafe1 SQL注入安全示例1 - 整数参数化查询
func SqlInjectionSafe1(c *gin.Context) {
	idStr := c.DefaultQuery("id", "1")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid id parameter")
		return
	}
	db, err := sql.Open("mysql", source)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()
	sqlStr := "select * from user where id=?"
	user := models.User{}
	err = db.QueryRow(sqlStr, id).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// SqlInjectionSafe2 SQL注入安全示例2 - 字符串参数化查询
func SqlInjectionSafe2(c *gin.Context) {
	username := c.Query("username")
	db, err := sql.Open("mysql", source)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()
	sqlStr := "select * from user where username=?"
	user := models.User{}
	err = db.QueryRow(sqlStr, username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// SqlInjectionSafe3 SQL注入安全示例3 - ORM安全查询
func SqlInjectionSafe3(c *gin.Context) {
	username := c.Query("username")
	field := c.Query("field")
	engine, err := xorm.NewEngine("mysql", source)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	engine.ShowSQL(true)
	user := models.User{}

	// 构建安全的查询条件
	session := engine.Where(field+" LIKE ?", "%"+username+"%")
	found, err := session.Get(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if !found {
		c.String(http.StatusNotFound, "User not found")
		return
	}
	c.JSON(http.StatusOK, user)
}
