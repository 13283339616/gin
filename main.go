package main

import (
	"github.com/13283339616/controller"
	"github.com/13283339616/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()
	mid(r)
	route(r)
	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
	r.Run()
}

//mid 注册全局中间件
func mid(r *gin.Engine) {
	r.Use(middleware.StatCost())
}

//route 控制路由
func route(r *gin.Engine) {
	hello := controller.Hello{}
	hello.Load(r, nil)

}
