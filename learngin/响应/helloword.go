/*
 * @Author: zzzzztw
 * @Date: 2023-07-02 11:35:03
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-08 13:15:46
 * @FilePath: /learngin/helloword.go
 */
package main

import (
	"github.com/gin-gonic/gin"
)

func Idex(c *gin.Context) {
	c.String(200, "hello world\n")
}

/*
func main() {
	//创建一个默认路由
	r := gin.Default()

	//绑定路由规则和路由函数，访问这个地址的路由将有对应的函数去处理
	r.GET("/helloworld/", Idex)

	//启动监听，gin会吧web服务放在0.0.0.0：10086端口上
	r.Run(":10086")
}
*/
