/*
 * @Author: zzzzztw
 * @Date: 2023-07-08 16:03:11
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-08 16:37:22
 * @FilePath: /learngin/post.go
 */
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func _queiry(c *gin.Context) {

	user := c.Query("user")
	fmt.Println(user)

	user2, ok := c.GetQuery("id")
	fmt.Println(user2, ok)

	fmt.Println(c.QueryArray("user")) //如果有多个user 返回所有user的值
}

func _Param(c *gin.Context) {
	fmt.Println(c.Param("userid"))
	fmt.Println(c.Param("book"))

}

func main() {
	r := gin.Default()
	r.GET("/quiry", _queiry)
	r.GET("/param/:userid", _Param)
	r.GET("/param/:userid/:book", _Param)

	r.Run(":10086")
}
