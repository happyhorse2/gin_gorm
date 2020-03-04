package main

import (
	"fmt"
	gormtest "gin_gorm/gormtest"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	testgorm()
}

func testgin() {
	//fmt.Println("hello world")
	r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello World")
	// })
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World",
	// 	})
	// })
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		fmt.Println("2222", name)
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 此 handler 将匹配 /user/john/ 和 /user/john/send
	// 如果没有其他路由匹配 /user/john，它将重定向到 /user/john/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		fmt.Println("111111", name, action, "222")
		message := name + " is " + action
		fmt.Println(message)
		c.String(http.StatusOK, message)
	})
	r.Run()
}

func testgorm() {
	gormtest.Gormtest()
}
