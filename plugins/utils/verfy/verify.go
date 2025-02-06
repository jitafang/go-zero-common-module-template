package verfy

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Verify struct {
	validate *validator.Validate
}

func NewVerify() *Verify {
	validate := validator.New()
	registerValidation(validate)
	return &Verify{
		validate: validate,
	}
}

func registerValidation(validate *validator.Validate) {
	_ = validate.RegisterValidation("space", func(fl validator.FieldLevel) bool {
		return space(fl.Field().String())
	})
	_ = validate.RegisterValidation("tel", func(fl validator.FieldLevel) bool {
		return tel(fl.Field().String())
	})
	_ = validate.RegisterValidation("password1", func(fl validator.FieldLevel) bool {
		return password1(fl.Field().String())
	})
}

func (v *Verify) ValidateAndHandleError(req interface{}) error {
	if err := v.validate.Struct(req); err != nil {
		// 转换为 validator.ValidationErrors
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				switch e.Tag() {
				case "space":
					return fmt.Errorf("%s 字段不能包含只有空格", e.Field())
				case "tel":
					return fmt.Errorf("手机号不合规")
				case "password1":
					return fmt.Errorf("密码要以字母开头，包含字母、数字、下划线，长度在6到18之间")
				default:
					return errors.New(e.Error()) // 默认错误信息
				}
			}
		}
		return err // 其他非校验错误
	}
	return nil
}
