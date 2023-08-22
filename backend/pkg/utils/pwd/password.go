package pwd

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/8/22 10:59
//

import (
	"golang.org/x/crypto/bcrypt"
)

func EncodePassword(rawPassword string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	return string(hash)
}

func ValidatePassword(encodePassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodePassword), []byte(inputPassword))
	return err == nil
}
