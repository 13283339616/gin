package controller

import (
	"encoding/json"
	"github.com/13283339616/service"
	"github.com/13283339616/vo"
	"github.com/gin-gonic/gin"
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
	err, m := baseRes.ResponseValid(loginRequestVo)
	s := service.HelloService{}
	token, err := s.Index(*loginRequestVo)
	if err != nil {
		c.JSON(http.StatusOK, baseRes.ResponseError(m))
	} else {
		responseVo := vo.LoginResponseVo{}
		responseVo.Token = token
		c.JSON(http.StatusOK, baseRes.ResponseSuccess(responseVo))
	}
}
