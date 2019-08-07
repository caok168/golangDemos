package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

var router *gin.Engine

func init() {
	// 初始化路由
	router = gin.Default()
	router.GET("/getString", OnGetStringRequest)
	router.POST("/loginForm", OnLoginRequestForForm)
	router.POST("/loginJson", OnLoginRequestForJson)
	//router.LoadHTMLGlob("E:/mygo/resources/pages/*") //定义模板文件路径
	router.GET("/practice", OnPracticeRequest)
}

// TestOnGetStringRequest 测试以Get方式获取一段字符串的接口
func TestOnGetStringRequest(t *testing.T) {
	// 初始化请求地址
	uri := "/getString"

	// 发起Get请求
	body := Get(uri, router)
	fmt.Printf("response:%v\n", string(body))

	// 判断响应是否与预期一致
	if string(body) != "success" {
		t.Errorf("响应字符串不符，body:%v\n", string(body))
	}
}

// TestOnLoginRequestForJson 测试以Json形式传递参数的登录接口
func TestOnLoginRequestForJson(t *testing.T) {
	// 初始化请求地址和请求参数
	uri := "/loginJson"

	param := make(map[string]interface{})
	param["username"] = "valiben"
	param["password"] = "123"
	param["age"] = 1

	// 发起post请求，以Json形式传递参数
	body := PostJson(uri, param, router)
	fmt.Printf("response:%v\n", string(body))

	// 解析响应，判断响应是否与预期一致
	response := &LoginResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}
	if response.Data.Username != "valiben" {
		t.Errorf("响应数据不符，username:%v\n", response.Data.Username)
	}
}

// Get 根据特定请求uri，发起get请求返回响应
func Get(uri string, router *gin.Engine) []byte {
	// 构造get请求
	req := httptest.NewRequest("GET", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

// PostForm 根据特定请求uri和参数param，以表单形式传递参数，发起post请求返回响应
func PostForm(uri string, param map[string]string, router *gin.Engine) []byte {
	// 构造post请求，表单数据以querystring的形式加在uri之后
	req := httptest.NewRequest("POST", uri+ParseToStr(param), nil)

	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

// PostJson 根据特定请求uri和参数param，以Json形式传递参数，发起post请求返回响应
func PostJson(uri string, param map[string]interface{}, router *gin.Engine) []byte {
	// 将参数转化为json比特流
	jsonByte, _ := json.Marshal(param)

	// 构造post请求，json数据以请求body的形式传递
	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))

	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}
