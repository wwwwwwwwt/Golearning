<!--
 * @Author: zzzzztw
 * @Date: 2023-07-08 13:09:27
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-08 16:39:29
 * @FilePath: /learngin/readme.md
-->
# gin学习笔记

# 1. helloworld 

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func Idex(c *gin.Context) {
	c.String(200, "hello world\n")
}

func main() {
	//创建一个默认路由
	r := gin.Default()

	//绑定路由规则和路由函数，访问这个地址的路由将有对应的函数去处理
	r.GET("/helloworld/", Idex)

	//启动监听，gin会吧web服务放在0.0.0.0：10086端口上, run本质就是ListenandServe的进一步封装
	r.Run(":10086")
}


```

# 2. 响应

## 响应字符串

```go
func Idex(c *gin.Context) {
	c.String(200, "hello world\n")
}
```

## 响应json

```go
1. 按结构体传递
func _json(c *gin.Context) {
	//json响应结构体
	type Userinfo struct {
		Username string `json:"user_name"`
		Age      int    `json:"age"`
		Password int    `json:"-"` //打上- 响应是不渲染password，忽略转换为json
	}
	c.JSON(200, Userinfo{
		Username: "ztw",
		Age:      10,
		Password: 123456,
	})
}
2.按map传递json
func _json(c *gin.Context){
	c.JSON(200, gin.H{
		"username": "ztw",
		"age":      10,
	})
}

3.gin.H{}直接传递
func _json(c *gin.Context){
	c.JSON(200, gin.H{
		"username": "ztw",
		"age":      10,
	})
}
```

## 响应html

1. 需要加载模板 r.LoadHTML()
```go
func _html(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
func main() {
	r := gin.Default()
	//加载模板目录下所有模板目录
	r.LoadHTMLGlob("template/*")
	r.GET("/", _html)
	r.Run(":10086")
}

```

## 响应文件

```go
//在golang中没有相对文件路径的概念，只有相对于项目的路径
r.StaticFS("/static", http.Dir("static/static"))
//配置单个文件，网页请求的路由，文件的路径
r.StaticFile("/titian.png", "static/tiantian.png")
```

## 重定向

301 可缓存重定向 ， 后台改了也没用，一直去第一次缓存的地址  
302 不可缓存重定向，每次都向后台去询问，除非是Expires或cache-control没过期

```go
func _redirect(c *gin.Context) {
	c.Redirect(301, "https://www.baidu.com")
}


```

# 3.请求

## 查询参数Query
查询形式：/xxx?user=ztw&id=2

```go
func _queiry(c *gin.Context) {

	user := c.Query("user")
	fmt.Println(user)

	user2, ok := c.GetQuery("id")
	fmt.Println(user2, ok)

	fmt.Println(c.QueryArray("user")) //如果有多个user 返回所有user的值
}

func main() {
	r := gin.Default()
	r.GET("/quiry", _queiry)
	r.Run(":10086")
}


```

## 动态参数 -> 动态路由 Param

```go
func _Param(c *gin.Context) {
	fmt.Println(c.Param("userid"))
	fmt.Println(c.Param("book"))

}
//http://localhost:10086/param/userid=123/book=1
r.GET("/param/:userid", _Param)
r.GET("/param/:userid/:book", _Param)


```
## 表单参数 PostForm