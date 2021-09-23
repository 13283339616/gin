package controller

import (
	"encoding/json"
	"fmt"
	"github.com/13283339616/vo"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
)

type Hello struct {
}

func (hello *Hello) Load(r *gin.Engine, f func() gin.HandlerFunc) {
	userGroup := r.Group("/hello")
	{
		userGroup.POST("/index", Index)
	}
}

func Index(c *gin.Context) {
	b, _ := c.GetRawData()
	loginRequestVo := &vo.LoginRequestVo{}
	_ = json.Unmarshal(b, loginRequestVo)
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	//验证器注册翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}
	err = validate.Struct(loginRequestVo)
	// 方法二：使用结构体
	var msg struct {
		Code    int
		Message string
		Data    interface{}
	}
	msg.Code = 200
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			msg.Message = err.Translate(trans)
			msg.Code = 301
			c.JSON(http.StatusOK, msg)
			return
		}
	}
	fmt.Println(loginRequestVo.Username)
	msg.Message = "获取成功"

	c.JSON(http.StatusOK, msg)
}
