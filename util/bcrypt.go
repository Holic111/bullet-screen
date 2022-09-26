package util

import "golang.org/x/crypto/bcrypt"

//加密

// 密码加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// 对比密码
// @Param1 password 输入的密码
// @Param2 hash Hash密码
func BcryptCheck(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}