package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//查询参数
func _query(c *gin.Context) {
	fmt.Println(c.Query("user"))
	fmt.Println(c.GetQuery("user"))
	fmt.Println(c.QueryArray("user")) // 拿到多个相同的查询参数
	fmt.Println(c.DefaultQuery("addr", "四川省"))
}

//动态参数
func _param(c *gin.Context) {
	fmt.Println(c.Param("user_id"))
	fmt.Println(c.Param("book_id"))
}

//表单
func _form(c *gin.Context) {
	username := c.PostForm("username")
	password := c.DefaultPostForm("password", "default_password")
	id := c.DefaultPostForm("id", "--") //传的是空字符串也会返回空字符串而没有使用默认值，可以探究一下
	hobbies := c.PostFormArray("hobbies")
	hobbiesStr := strings.Join(hobbies, ", ")
	// 获取上传的文件和其他表单字段
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusInternalServerError, "MultipartForm Error: %s", err.Error())
		return
	}

	// 获取上传的文件
	// 获取上传的文件
	files := form.File["file"]
	for _, file := range files {
		// 生成保存文件的绝对路径
		targetDir := "C:\\Users\\19778\\Desktop\\"
		targetPath := filepath.Join(targetDir, file.Filename)

		// 确保目录存在
		err := os.MkdirAll(targetDir, 0755)
		if err != nil {
			c.String(http.StatusInternalServerError, "MkdirAll Error: %s", err.Error())
			return
		}

		// 保存上传的文件
		err = c.SaveUploadedFile(file, targetPath)
		if err != nil {
			c.String(http.StatusInternalServerError, "SaveUploadedFile Error: %s", err.Error())
			return
		}
	}

	c.String(http.StatusOK, "Username: %s\nPassword: %s\nid: %s\nHobbies: %s\nFile(s) uploaded successfully.",
		username, password, id, hobbiesStr)
}

//请求头
func _handler(c *gin.Context) {
	c.Header("Content-Type", "text/html") //text/plain
	ContentType := c.GetHeader("Content-Type")
	// 获取自定义的请求头字段
	//customHeader := c.GetHeader("X-Custom-Header")
	//if c.GetHeader("Authorization") != "" {
	//	// 处理身份验证逻辑
	//}
	fmt.Println(ContentType)
}

//响应头
func _Header(c *gin.Context) {
	// 设置响应头的 Content-Type 字段为 JSON 格式
	c.Header("Content-Type", "text/html") //application/json

	// 返回一个 JSON 数据作为响应
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
func main() {
	router := gin.Default()       // 创建一个默认的路由
	router.LoadHTMLGlob("html/*") // 绑定路由规则和路由函数，访问/index的路由，将由对应的函数去处理
	router.GET("/Login", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.GET("/json", _Header)
	router.GET("/handler", _handler)
	router.GET("/query", _query)
	router.GET("/param/:user_id/", _param)
	router.GET("/param/:user_id/:book_id", _param)
	router.POST("/index", _form)
	router.Run(":8080")
}
