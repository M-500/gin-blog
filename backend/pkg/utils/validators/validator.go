package validators

import (
	"backend/pkg/utils/strs"
	"errors"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidatePhone(f1 validator.FieldLevel) bool {
	phone := f1.Field().String()
	ok, _ := regexp.MatchString("^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\\d{8}$", phone)
	if ok {
		return true
	}
	return false
}

func IsPassword(password, rePassword string) error {
	if strs.IsBlank(password) {
		return errors.New("请输入密码")
	}
	if strs.RuneLen(password) < 6 {
		return errors.New("密码过于简单")
	}
	if password != rePassword {
		return errors.New("两次输入密码不匹配")
	}
	return nil
}
