package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

// OnGetStringRequest 返回success字符串的接口
func OnGetStringRequest(c *gin.Context) {
	c.String(http.StatusOK, "success")
}

// OnPracticeRequest 返回practice.html页面的接口
func OnPracticeRequest(c *gin.Context) {
	c.HTML(http.StatusOK, "practice.html", gin.H{})
}

// OnLoginRequestForForm 以表单形式传递参数的登录接口
func OnLoginRequestForForm(c *gin.Context) {
	req := &User{}
	if err := c.ShouldBindWith(req, binding.Form); err != nil {
		fmt.Printf("err:%v", err)
		c.JSON(http.StatusOK, gin.H{
			"errno":  "1",
			"errmsg": "参数不匹配",
			"data":   "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"errno":  "0",
		"errmsg": "",
		"data":   req,
	})
}

// OnLoginRequestForJson 以Json形式传递参数的登录接口
func OnLoginRequestForJson(c *gin.Context) {
	req := &User{}
	if err := c.ShouldBindWith(req, binding.JSON); err != nil {
		fmt.Printf("err:%v", err)
		c.JSON(http.StatusOK, gin.H{
			"errno":  "1",
			"errmsg": "参数不匹配",
			"data":   "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"errno":  "0",
		"errmsg": "",
		"data":   req,
	})
}

//承接前端传过来的json数据或form表单数据
type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Age      int    `form:"age" json:"age" binding:"required"`
}

// LoginResponse 登录接口的响应参数
type LoginResponse struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
	Data   User   `json:"data"`
}

// ParseToStr 将map中的键值对输出成querystring形式
func ParseToStr(mp map[string]string) string {
	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + val
	}
	temp := values[1:]
	values = "?" + temp
	return values
}
