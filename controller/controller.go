package controller

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var validate *validator.Validate
var trans ut.Translator
var baseRes *Res

func init() {
	uni := ut.New(zh.New())
	transNew, _ := uni.GetTranslator("zh")
	trans = transNew
	validate = validator.New()
	//验证器注册翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}
	baseRes = new(Res)
}

type Res struct {
}
type ResDetail struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Res) ResponseSuccess(v interface{}) (res ResDetail) {
	res = ResDetail{
		Code:    200,
		Message: "返回成功",
		Data:    v,
	}
	return res
}

func (r *Res) ResponseError(v string) (res ResDetail) {
	res = ResDetail{
		Code:    500,
		Message: v,
		Data:    nil,
	}
	return res
}

func (r *Res) ResponseValid(v interface{}) (err error, msg string) {
	errNew := validate.Struct(v)
	if errNew != nil {
		for _, errNew := range errNew.(validator.ValidationErrors) {
			msg = errNew.Translate(trans)
		}
	}
	err = errNew
	return err, msg

}

func init() {

}
