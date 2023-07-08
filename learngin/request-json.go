/*
 * @Author: zzzzztw
 * @Date: 2023-07-08 13:22:05
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-08 15:52:59
 * @FilePath: /learngin/request-json.go
 */
package main

import "github.com/gin-gonic/gin"

func _string(c *gin.Context) {
	c.String(200, "hello world\n")
}

func _json(c *gin.Context) {
	//json响应结构体
	// type Userinfo struct {
	// 	Username string `json:"user_name"`
	// 	Age      int    `json:"age"`
	// 	Password int    `json:"-"` //打上- 响应是不渲染password，忽略转换为json
	// }
	// c.JSON(200, Userinfo{
	// 	Username: "ztw",
	// 	Age:      10,
	// 	Password: 123456,
	// })

	//json响应map
	// mp := map[string]string{
	// 	"age":      "123456",
	// 	"username": "ztw",
	// }

	// c.JSON(200, mp)

	//json直接响应

	c.JSON(200, gin.H{
		"username": "ztw",
		"age":      10,
	})
}

func _html(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func _redirect(c *gin.Context) {
	c.Redirect(301, "https://www.baidu.com")
}

func main() {
	r := gin.Default()
	//加载模板目录下所有模板目录
	r.LoadHTMLGlob("template/*")
	r.GET("/", _html)

	//重定向
	r.GET("/baidu", _redirect)
	r.Run(":10086")
}
