package service

import (
	"errors"
	"fmt"
	"github.com/13283339616/orm"
	"github.com/13283339616/util"
	"github.com/13283339616/vo"
)

type HelloService struct {
}

func (h HelloService) Index(v vo.LoginRequestVo) (token string, err error) {
	u := orm.User{}
	user := u.GetUserByUserCode(v.Username)
	if user == nil {
		return "", errors.New("用户不存在")
	}
	if !user.Enabled {
		return "", errors.New("用户被禁止登陆")
	}
	password := util.Md5([]byte(v.Password))
	fmt.Println(password)
	fmt.Println(user.Password)
	if user.Password != password {
		return "", errors.New("账号密码不正确")
	}

	return "abc", nil
}
