package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Hello struct {
}

func (hello *Hello) Load(r *gin.Engine, f func() gin.HandlerFunc) {
	userGroup := r.Group("/hello")
	{
		userGroup.GET("/index", Index)
	}
}

func Index(c *gin.Context) {
	// 方法二：使用结构体
	var msg struct {
		Name    string `json:"user"`
		Message string
		Age     int
	}
	msg.Name = "小王子"
	msg.Message = "Hello world!"
	msg.Age = 18
	c.JSON(http.StatusOK, msg)
}
